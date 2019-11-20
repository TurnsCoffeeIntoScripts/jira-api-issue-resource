package rest

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/configuration"
	"net/http"
	"strings"
)

type CreateBodyFN func() []byte
type GetEndpointFN func() string
type JsonObjectFN func() interface{}

type JiraAPI struct {
	HttpMethod string
	Body       []byte
	JsonObject interface{}

	// 'url' is not exported because at this level it contains the credentials
	url string
}

func CreateAPIFromParams(params configuration.JiraAPIResourceParameters, fnBody CreateBodyFN, fnEndpoint GetEndpointFN, fnJsonObj JsonObjectFN, httpMethod string) (JiraAPI, error) {
	var err error
	api := JiraAPI{}
	api.createUrlWithCredentials(*params.JiraAPIUrl, *params.Username, *params.Password)
	if fnBody != nil {
		api.Body = fnBody()
	}

	if fnJsonObj != nil {
		api.JsonObject = fnJsonObj()
	}

	api.HttpMethod = httpMethod

	if fnEndpoint == nil {
		return api, errors.New("not allowed to have null endpoint creation function")
	} else {
		api.url = api.url + fnEndpoint()
	}

	return api, err
}

func (api *JiraAPI) Call() (interface{}, error) {
	client := http.DefaultClient
	req, newRequestErr := http.NewRequest(api.HttpMethod, api.url, bytes.NewBuffer(api.Body))

	if newRequestErr != nil {
		return nil, newRequestErr
	}

	if req != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, _ := client.Do(req)
	fmt.Printf("Received response with HTTP %s\n", resp.Status)

	if resp != nil {
		defer resp.Body.Close()
	}

	buffer := new(bytes.Buffer)
	if resp != nil {
		count, readBodyErr := buffer.ReadFrom(resp.Body)

		fmt.Printf("Read %v bytes\n", count)

		if readBodyErr != nil {
			fmt.Println("Unable to read body of response")
			return nil, readBodyErr
		}
	}

	err := json.Unmarshal(buffer.Bytes(), &api.JsonObject)

	return api.JsonObject, err
}

func (api *JiraAPI) createUrlWithCredentials(url, username, password string) {
	// Construct the user:password@ string
	usrpw := username + ":" + password + "@"

	// Find the index to which we'll insert said string in url
	usrpwInsertIndex := strings.Index(url, "://") + 3

	// Remove any trailing '/' from the url because if needed they'll be added later on
	// by the specific service executing the api call
	strings.TrimSuffix(url, "/")

	// Build the final url value
	api.url = url[:usrpwInsertIndex] + usrpw + url[usrpwInsertIndex:]
}

/*func ApiCall(ctx configuration.Context) (bool, error) {

	return true, nil

}

func executeApiCall(md configuration.Metadata, op string, body []byte, needDomainObject bool) (*domain.Issue, error) {
	client := http.DefaultClient

	baseUrl := ""
	//baseUrl := md.AuthenticatedUrl()
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
	buffer := new(bytes.Buffer)
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
*/
