package reading

import (
	"errors"
	"fmt"
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/configuration"
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/helpers"
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/http/rest"
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/service"
	"net/http"
)

type ServiceReadIssue struct {
	issueId         string
	customFieldKey  string
	customFieldName string
}

func (s *ServiceReadIssue) InitJiraAPI(params configuration.JiraAPIResourceParameters) (rest.JiraAPI, error) {
	s.issueId = params.ActiveIssue
	s.customFieldName = *params.CustomFieldName

	return service.PreInitJiraAPI(s, params, http.MethodGet)
}

func (s *ServiceReadIssue) GetReadResults() map[string]string {
	var m = make(map[string]string)
	m[helpers.CustomFieldKey] = s.customFieldKey
	return m
}

func (s *ServiceReadIssue) SetReadResults(result map[string]string) {

}

func (s *ServiceReadIssue) GetEndpoint() string {
	return fmt.Sprintf("/issue/%s?expand=names", s.issueId)
}

func (s *ServiceReadIssue) CreateBody() []byte {
	return nil
}

func (s *ServiceReadIssue) JSONObject() interface{} {
	return &Issue{}
}

func (s *ServiceReadIssue) PostAPICall(result interface{}) error {
	if issue, ok := result.(*Issue); !ok {
		return errors.New("failed to convert result of type interface{} to issue of type reading.Issue")
	} else {

		if s.customFieldName != "" {
			s.customFieldKey = helpers.FindCustomName(issue.Names.CustomFields, s.customFieldName)
		}
	}

	return nil
}
