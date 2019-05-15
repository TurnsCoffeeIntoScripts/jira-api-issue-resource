package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/domain"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetIssue(md Metadata, prefix string, id string) error {
	return apiCall(md, "issue/" + prefix + "-" + string(id))
}

func apiCall(md Metadata, op string) error {
	baseUrl := md.Url
	if !strings.HasSuffix(baseUrl, "/") {
		baseUrl = baseUrl + "/"
	}

	if !strings.HasSuffix(op, "/") {
		op = op + "/"
	}

	url := baseUrl + op
	response, err := http.Get(url)

	if err != nil {
		errMsg := fmt.Sprintf("http request failed with error %s", err)
		return errors.New(errMsg)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		issue := domain.Issue{}
		unmarshalErr := json.Unmarshal(data, issue)

		if unmarshalErr != nil {
			return errors.New("unmarshalling failed")
		} else {
			fmt.Println(issue.Id)
		}
	}

	return nil
}
