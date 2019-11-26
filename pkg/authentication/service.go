package authentication

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/configuration"
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/http/rest"
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/service"
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/status"
	"net/http"
	"regexp"
	"strings"
)

type ServiceAuthenticateSession struct {
	username string
	password string
}

// See service/service.go for details
func (s *ServiceAuthenticateSession) InitJiraAPI(params configuration.JiraAPIResourceParameters) (rest.JiraAPI, error) {
	s.username = *params.Username
	s.password = *params.Password

	return service.PreInitJiraAPI(s, params, http.MethodPost)
}

// See service/service.go for details
func (s *ServiceAuthenticateSession) GetResults() map[string]string {
	return nil
}

// See service/service.go for details
func (s *ServiceAuthenticateSession) SetResultsFromPrevious(result map[string]string) {
	// noop
}

// See service/service.go for details
func (s *ServiceAuthenticateSession) GetEndpoint(url string) string {
	url = strings.TrimSuffix(url, "")
	rgx := regexp.MustCompile(`(^.*)/api/.*`)
	res := rgx.ReplaceAllString(url, "${1}")

	return fmt.Sprintf("%s/auth/1/session", res)
}

// See service/service.go for details
func (s *ServiceAuthenticateSession) CreateRequestBody() []byte {
	uc := UserCredentials{}

	uc.Username = s.username
	uc.Password = s.password

	b, err := json.Marshal(uc)
	if err != nil {
		b, _ := json.Marshal(UserCredentials{})
		return b
	}

	return b
}

// See service/service.go for details
func (s *ServiceAuthenticateSession) JSONResponseObject() interface{} {
	return &Session{}
}

// See service/service.go for details
func (s *ServiceAuthenticateSession) PostAPICall(result interface{}) error {
	if session, ok := result.(*Session); !ok {
		return errors.New("failed to convert result of type interface{} to object of type authentication.SessionInfo")
	} else {
		status.SessionName = session.Info.Name
		status.SessionValue = session.Info.Value

		status.Secured = (status.SessionValue != "") && (status.SessionName != "")
	}

	return nil
}

func (s *ServiceAuthenticateSession) Name() string {
	return "ServiceAuthenticateSession"
}
