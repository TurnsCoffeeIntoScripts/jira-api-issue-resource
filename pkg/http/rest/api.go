package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/domain"
	"net/http"
	"strings"
)

func IssueExists(md Metadata, id string) bool {
	issue, err := apiCall(md, "issue/" + string(id))
	if issue == nil || err != nil {
		return false
	}

	return true
}

func GetIssue(md Metadata, id string) (*domain.Issue, error) {
	return apiCall(md, "issue/" + string(id))
}

func apiCall(md Metadata, op string) (*domain.Issue, error) {
	client := http.DefaultClient

	baseUrl := md.AuthenticatedUrl()
	if !strings.HasSuffix(baseUrl, "/") {
		baseUrl = baseUrl + "/"
	}

	if !strings.HasSuffix(op, "/") {
		op = op + "/"
	}

	url := baseUrl + op
	req, newReqErr := http.NewRequest(md.HttpMethod, url, nil)

	if newReqErr != nil {
		errMsg := fmt.Sprintf("http request failed with error %s", newReqErr)
		return nil, errors.New(errMsg)
	}

	resp, respErr := client.Do(req)

	if respErr != nil {
		errMsg := fmt.Sprintf("http request failed with error %s", respErr)
		return nil, errors.New(errMsg)
	}

	defer resp.Body.Close()

	var issue domain.Issue

	if decodeErr := json.NewDecoder(resp.Body).Decode(&issue); decodeErr != nil {
		errMsg := fmt.Sprintf("http request failed with error %s", decodeErr)
		return nil, errors.New(errMsg)
	}

	if issue.Id == "" {
		errMsg := fmt.Sprintf("Issue doesn't exists")
		return nil, errors.New(errMsg)
	}

	return &issue, nil
}
