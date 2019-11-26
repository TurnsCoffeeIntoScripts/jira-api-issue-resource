package chaining

import (
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/authentication"
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/configuration"
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/editing"
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/logout"
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/noop"
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/reading"
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/service"
)

const (
	ServiceAuthenticateName    = "srv_auth_in"
	ServiceReadIssueName       = "srv_read"
	ServiceEditCustomFieldName = "srv_edit_field"
	ServiceLogoutName          = "srv_auth_out"
	ServiceUnknownName         = "srv_unknown"
)

var serviceRegistry = make(map[string]service.Service)

func InitServiceRegistry() {
	serviceRegistry[ServiceAuthenticateName] = &authentication.ServiceAuthenticateSession{}
	serviceRegistry[ServiceReadIssueName] = &reading.ServiceReadIssue{}
	serviceRegistry[ServiceEditCustomFieldName] = &editing.ServiceEditCustomField{}
	serviceRegistry[ServiceLogoutName] = &logout.ServiceLogoutSession{}
	serviceRegistry[ServiceUnknownName] = &noop.ServiceUnknown{}
}

func GetServicesChain(c configuration.Context, secured bool) []service.Service {
	chain := make([]service.Service, 0)

	switch c {
	case configuration.ReadIssue:
		chain = append(chain, serviceRegistry[ServiceReadIssueName])
	case configuration.EditCustomField:
		chain = append(chain, serviceRegistry[ServiceReadIssueName])
		chain = append(chain, serviceRegistry[ServiceEditCustomFieldName])
	case configuration.Unknown:
		fallthrough
	default:
		chain = make([]service.Service, 0)
	}

	return checkIfAuthenticationNeeded(chain, secured)
}

func checkIfAuthenticationNeeded(chain []service.Service, secured bool) []service.Service {
	if secured {
		newChain := make([]service.Service, 0)
		newChain = append(newChain, serviceRegistry[ServiceAuthenticateName])
		newChain = append(newChain, chain...)
		newChain = append(newChain, serviceRegistry[ServiceLogoutName])

		return newChain
	}

	return chain
}
