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
	issueId   string
	fieldKey  string
	fieldName string
}

func (s *ServiceReadIssue) InitJiraAPI(params configuration.JiraAPIResourceParameters) (rest.JiraAPI, error) {
	s.issueId = params.ActiveIssue
	s.fieldName = *params.CustomFieldName

	return service.PreInitJiraAPI(s, params, http.MethodGet)
}

func (s *ServiceReadIssue) GetResults() map[string]string {
	var m = make(map[string]string)
	m[helpers.ReadingFieldKey] = s.fieldKey
	return m
}

func (s *ServiceReadIssue) SetResultsFromPrevious(result map[string]string) {
}

func (s *ServiceReadIssue) GetEndpoint(url string) string {
	return fmt.Sprintf("%s/issue/%s?expand=names", url, s.issueId)
}

func (s *ServiceReadIssue) CreateRequestBody() []byte {
	return nil
}

func (s *ServiceReadIssue) JSONResponseObject() interface{} {
	return &Issue{}
}

func (s *ServiceReadIssue) PostAPICall(result interface{}) error {
	if issue, ok := result.(*Issue); !ok {
		return errors.New("failed to convert result of type interface{} to issue of type reading.Issue")
	} else {

		if s.fieldName != "" {
			s.fieldKey = helpers.FindCustomName(issue.Names.CustomFields, s.fieldName)
		}
	}

	return nil
}

func (s *ServiceReadIssue) Name() string {
	return "ServiceReadIssue"
}