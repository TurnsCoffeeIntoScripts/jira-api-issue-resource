package rest

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/configuration"
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/domain"
	"net/http"
	"strings"
)

func ApiCall(ctx configuration.Context) (bool, error) {
	if ctx.Metadata.ResourceFlags.ZeroIssue {
		return noIssueApiCall(ctx)
	} else if ctx.Metadata.ResourceFlags.SingleIssue {
		return singleApiCall(ctx, ctx.IssueIds[0])
	} else {
		return multipleApiCall(ctx)
	}

}

func noIssueApiCall(ctx configuration.Context) (bool, error) {
	ctx.Metadata.HttpMethod = ctx.HttpMethod
	_, err := executeApiCall(ctx.Metadata, ctx.ApiEndPoint, ctx.Body, false)

	return err == nil, err
}

func singleApiCall(ctx configuration.Context, issueId string) (bool, error) {
	ctx.Metadata.HttpMethod = http.MethodGet
	issue, getErr := executeApiCall(ctx.Metadata, "/issue/"+issueId, nil, true)
	if getErr != nil {
		return false, getErr
	}

	currentIssueId := issueId
	if ctx.ForceOnParent && issue.HasParent() {
		currentIssueId = issue.GetParent().Key
	}

	apiOperation := checkUrl(ctx.ApiEndPoint, configuration.IssuePlaceholder, currentIssueId)

	ctx.Metadata.HttpMethod = ctx.HttpMethod
	_, err := executeApiCall(ctx.Metadata, apiOperation, ctx.Body, true)

	return err == nil, err
}

func multipleApiCall(ctx configuration.Context) (bool, error) {
	allOk := true
	for idx, i := range ctx.IssueIds {
		ok, err := singleApiCall(ctx, i)

		if err != nil {
			allOk = false

			fmt.Printf("API call #%d failed with reason %v", idx, err)
			if !*ctx.Metadata.ResourceFlags.ForceFinish {
				return ok, err
			}
		}
	}

	if !allOk {
		return false, errors.New("one or more API call failed")
	}

	return true, nil
}

func checkUrl(url string, findValue, replaceValue string) string {
	if strings.Contains(url, findValue) {
		return strings.ReplaceAll(url, findValue, replaceValue)
	}

	return url
}

func apiCallDummy(md configuration.Metadata, op string, body []byte) (*domain.Issue, error) {
	return nil, nil
}

func executeApiCall(md configuration.Metadata, op string, body []byte, needDomainObject bool) (*domain.Issue, error) {
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

	fmt.Println(op)
	fmt.Println(string(body))

	resp, respErr := client.Do(req)

	defer resp.Body.Close()

	status := 0
	buffer := new (bytes.Buffer)
	if resp != nil {
		_, readBodyErr := buffer.ReadFrom(resp.Body)

		if readBodyErr == nil {
			fmt.Println("BODY:")
			fmt.Println(buffer.String())
		}

		status = resp.StatusCode
	}

	if status < http.StatusOK || status > 299 {
		errMsg := fmt.Sprintf("http request failed (HTTP %d)", status)
		return nil, errors.New(errMsg)
	}

	if respErr != nil {
		errMsg := fmt.Sprintf("http request failed (HTTP %d) with error %s", status, respErr)
		return nil, errors.New(errMsg)
	}

	if needDomainObject && md.HttpMethod == http.MethodGet {
		var issue domain.Issue

		if decodeErr := json.Unmarshal(buffer.Bytes(), &issue); decodeErr != nil {
			errMsg := fmt.Sprintf("decoding of json object failed with error %s", decodeErr)
			return nil, errors.New(errMsg)
		}

		if issue.Id == "" {
			errMsg := fmt.Sprintf("Issue doesn't exists")
			return nil, errors.New(errMsg)
		}
		return &issue, nil
	}

	return nil, nil
}
