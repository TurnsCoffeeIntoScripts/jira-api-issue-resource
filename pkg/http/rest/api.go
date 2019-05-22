package rest

// TODO: https://docs.atlassian.com/software/jira/docs/api/REST/8.2.0/

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/domain"
	"net/http"
	"strings"
)

func Get(apiOperation string, id string, md Metadata) (bool, *domain.Issue) {
	apiOperation = checkUrl(apiOperation, id)

	md.HttpMethod = http.MethodGet
	issue, err := apiCall(md, apiOperation, nil)
	if issue == nil || err != nil {
		return false, nil
	}

	return true, issue
}

func Post(apiOperation string, id string, md Metadata, data map[string]string) bool {
	apiOperation = checkUrl(apiOperation, id)

	md.HttpMethod = http.MethodPost
	body := buildJsonBodyFromMap(data)
	issue, err := apiCall(md, apiOperation, body)
	if issue == nil || err != nil {
		return false
	}

	return true
}

func buildJsonBodyFromMap(data map[string]string) []byte {
	var buffer bytes.Buffer
	buffer.WriteString("{")
	for key, val := range data {
		buffer.WriteString("\"" + key + "\"")
		buffer.WriteString(":")
		buffer.WriteString("\"" + val + "\",")
	}

	stringBuffer := strings.TrimRight(buffer.String(), ",")
	buffer = *bytes.NewBufferString(stringBuffer)

	buffer.WriteString("}")

	return buffer.Bytes()
}

func checkUrl(url string, value string) string {
	if strings.Contains(url, "??") {
		return strings.ReplaceAll(url, "??", value)
	}

	return url
}

func apiCall(md Metadata, op string, body []byte) (*domain.Issue, error) {
	client := http.DefaultClient

	baseUrl := md.AuthenticatedUrl()
	if !strings.HasSuffix(baseUrl, "/") {
		baseUrl = baseUrl + "/"
	}

	if !strings.HasSuffix(op, "/") {
		op = op + "/"
	}

	url := baseUrl + op
	var req *http.Request
	var newReqErr error
	if body == nil {
		req, newReqErr = http.NewRequest(md.HttpMethod, url, nil)
	} else {
		req, newReqErr = http.NewRequest(md.HttpMethod, url, bytes.NewBuffer(body))
	}

	if newReqErr != nil {
		errMsg := fmt.Sprintf("http request failed with error %s", newReqErr)
		return nil, errors.New(errMsg)
	}

	if req != nil {
		req.Header.Set("Content-Type", "application/json")
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
