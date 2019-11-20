package service

import (
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/configuration"
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/http/rest"
)

type Service interface {
	InitJiraAPI(params configuration.JiraAPIResourceParameters) (rest.JiraAPI, error)
	GetReadResults() map[string]string
	SetReadResults(results map[string]string)
	GetEndpoint() string
	CreateBody() []byte
	JsonObject() interface{}
	PostAPICall(result interface{}) error
}

func PreInitJiraAPI(s Service, params configuration.JiraAPIResourceParameters, httpMethod string) (rest.JiraAPI, error) {
	api, err := rest.CreateAPIFromParams(params, s.CreateBody, s.GetEndpoint, s.JsonObject, httpMethod)
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
