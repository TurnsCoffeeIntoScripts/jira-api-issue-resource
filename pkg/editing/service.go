package editing

import (
	"encoding/json"
	"fmt"
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/configuration"
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/helpers"
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/http/rest"
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/service"
	"net/http"
	"strconv"
)

type ServiceEditCustomField struct {
	issueId          string
	customfieldKey   string
	customfieldValue string
}

func (s *ServiceEditCustomField) InitJiraAPI(params configuration.JiraAPIResourceParameters) (rest.JiraAPI, error) {
	s.issueId = params.ActiveIssue
	s.customfieldValue = *params.CustomFieldValue

	return service.PreInitJiraAPI(s, params, http.MethodPut)
}

func (s *ServiceEditCustomField) GetReadResults() map[string]string {
	return nil
}

func (s *ServiceEditCustomField) SetReadResults(result map[string]string) {
	s.customfieldKey = result[helpers.CustomFieldKey]
}

func (s *ServiceEditCustomField) GetEndpoint() string {
	return fmt.Sprintf("/issue/%s", s.issueId)
}

func (s *ServiceEditCustomField) CreateBody() []byte {
	i := Issue{}

	if numVal, err := strconv.Atoi(s.customfieldValue); err == nil {
		i.AddField(s.customfieldKey, numVal)
	} else {
		i.AddField(s.customfieldKey, s.customfieldValue)
	}
	b, err := json.Marshal(i)
	if err != nil {
		b, _ := json.Marshal(Issue{})
		return b
	}
	return b
}

func (s *ServiceEditCustomField) JsonObject() interface{} {
	return nil
}

func (s *ServiceEditCustomField) PostAPICall(result interface{}) error {
	return nil
}
