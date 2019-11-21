// Package editing provides a Jira API interface service and implementation of Jira's domain object as
// Go structures in the context of editing an issue.
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

// The ServiceEditCustomField struct implements the service.Service interface. It defines the workflow of editing
// an existing Jira issue.
type ServiceEditCustomField struct {
	issueId          string
	customfieldKey   string
	customfieldValue string
}

// See service/service.go for details
func (s *ServiceEditCustomField) InitJiraAPI(params configuration.JiraAPIResourceParameters) (rest.JiraAPI, error) {
	s.issueId = params.ActiveIssue
	s.customfieldValue = *params.CustomFieldValue

	return service.PreInitJiraAPI(s, params, http.MethodPut)
}

// See service/service.go for details
func (s *ServiceEditCustomField) GetReadResults() map[string]string {
	return nil
}

// See service/service.go for details
func (s *ServiceEditCustomField) SetReadResults(result map[string]string) {
	s.customfieldKey = result[helpers.CustomFieldKey]
}

// See service/service.go for details
func (s *ServiceEditCustomField) GetEndpoint() string {
	return fmt.Sprintf("/issue/%s", s.issueId)
}

// See service/service.go for details
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

// See service/service.go for details
func (s *ServiceEditCustomField) JSONObject() interface{} {
	return nil
}

// See service/service.go for details
func (s *ServiceEditCustomField) PostAPICall(result interface{}) error {
	return nil
}
