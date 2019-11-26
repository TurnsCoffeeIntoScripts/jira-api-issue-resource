package logout

import (
	"fmt"
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/configuration"
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/http/rest"
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/service"
	"net/http"
	"regexp"
	"strings"
)

type ServiceLogoutSession struct {
}

// See service/service.go for details
func (s *ServiceLogoutSession) InitJiraAPI(params configuration.JiraAPIResourceParameters) (rest.JiraAPI, error) {
	return service.PreInitJiraAPI(s, params, http.MethodDelete)
}

// See service/service.go for details
func (s *ServiceLogoutSession) GetResults() map[string]string {
	return nil
}

// See service/service.go for details
func (s *ServiceLogoutSession) SetResultsFromPrevious(result map[string]string) {
	// noop
}

// See service/service.go for details
func (s *ServiceLogoutSession) GetEndpoint(url string) string {
	url = strings.TrimSuffix(url, "")
	rgx := regexp.MustCompile(`(^.*)/api/.*`)
	res := rgx.ReplaceAllString(url, "${1}")

	return fmt.Sprintf("%s/auth/1/session", res)
}

// See service/service.go for details
func (s *ServiceLogoutSession) CreateRequestBody() []byte {
	return nil
}

// See service/service.go for details
func (s *ServiceLogoutSession) JSONResponseObject() interface{} {
	return nil
}

// See service/service.go for details
func (s *ServiceLogoutSession) PostAPICall(result interface{}) error {
	return nil
}

func (s *ServiceLogoutSession) Name() string {
	return "ServiceLogoutSession"
}
