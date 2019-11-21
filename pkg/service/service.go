package service

import (
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/configuration"
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/http/rest"
)

type Service interface {
	// Method that will init and return a rest.JiraAPI struct. The initialization is done according
	// to a specific context
	InitJiraAPI(params configuration.JiraAPIResourceParameters) (rest.JiraAPI, error)

	// This method will return a map containing the result(s) of execution of the service. This method is used only
	// when the service is used as a "pre-service" (called before the main service of the specified context)
	GetReadResults() map[string]string

	// This method takes a map as input that represents the result(s) of a previous service. Each service will fetch
	// and use the value of this map as it sees fit.
	SetReadResults(results map[string]string)

	// Returns the parametrized endpoint (what comes after the /rest/api/).
	GetEndpoint() string

	// Returns a slice of byte representing the JSON body of the request (if any).
	CreateBody() []byte

	// Returns a new object (struct) that is the implementation of the Jira object.
	JSONObject() interface{}

	// Any custom operation to be performed after the API call. Such as setting value in the map
	// returned by GetReadResults()
	PostAPICall(result interface{}) error
}

func PreInitJiraAPI(s Service, params configuration.JiraAPIResourceParameters, httpMethod string) (rest.JiraAPI, error) {
	api, err := rest.CreateAPIFromParams(params, s.CreateBody, s.GetEndpoint, s.JSONObject, httpMethod)
	if err != nil {
		return api, err
	}

	return api, nil
}

func ExecuteService(s Service, params configuration.JiraAPIResourceParameters) error {
	if params.Meta.MultipleIssue {
		for i := range params.IssueList {
			result, err := exec(s, params, i)
			if err != nil {
				return err
			}

			err = s.PostAPICall(result)
			if err != nil {
				return err
			}
		}

	} else {
		result, err := exec(s, params, 0)

		if err != nil {
			return err
		}

		return s.PostAPICall(result)
	}

	return nil
}

func exec(s Service, params configuration.JiraAPIResourceParameters, activeIssueIndex int) (interface{}, error) {
	params.ActiveIssue = params.IssueList[activeIssueIndex]
	api, err := s.InitJiraAPI(params)

	if err != nil {
		return nil, err
	} else {
		return api.Call()
	}
}
