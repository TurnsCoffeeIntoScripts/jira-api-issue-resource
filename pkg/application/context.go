// See application.go for this package's comment
package application

import (
	"errors"
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/configuration"
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/editing"
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/reading"
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/service"
)

func (app *JiraAPIResourceApp) executeFromContext() error {
	var srv service.Service
	var preSrv service.Service

	switch app.params.Context {
	case configuration.ReadIssue:
		srv = &reading.ServiceReadIssue{}
	case configuration.EditCustomField:
		srv = &editing.ServiceEditCustomField{}
		preSrv = &reading.ServiceReadIssue{}
	case configuration.Unknown:
		fallthrough
	default:
		srv = nil
	}

	if srv == nil {
		return errors.New("unable to determine inner service to execute from context")
	}

	if preSrv != nil {
		if err := service.ExecuteService(preSrv, app.params); err != nil {
			return err
		}

		srv.SetResultsFromPrevious(preSrv.GetResults())
	}

	return service.ExecuteService(srv, app.params)
}