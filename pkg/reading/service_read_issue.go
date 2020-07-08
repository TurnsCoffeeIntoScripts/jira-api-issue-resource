package reading

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/TurnsCoffeeIntoScripts/jira-api-issue-resource/pkg/assets"
	"github.com/TurnsCoffeeIntoScripts/jira-api-issue-resource/pkg/configuration"
	"github.com/TurnsCoffeeIntoScripts/jira-api-issue-resource/pkg/helpers"
	"github.com/TurnsCoffeeIntoScripts/jira-api-issue-resource/pkg/http/rest"
	resulthelper "github.com/TurnsCoffeeIntoScripts/jira-api-issue-resource/pkg/result"
	"github.com/TurnsCoffeeIntoScripts/jira-api-issue-resource/pkg/service"
	"net/http"
	"os"
)

type ServiceReadIssue struct {
	SkipCustomKeyRetrieval bool

	issueId     string
	fieldKey    string
	fieldName   string
	statusName  string
	destination string

	invokingContext configuration.Context
}

func (s *ServiceReadIssue) InitJiraAPI(params configuration.JiraAPIResourceParameters) (rest.JiraAPI, error) {
	s.issueId = params.ActiveIssue
	s.fieldName = *params.EditCustomFieldParam.CustomFieldName
	s.destination = *params.Destination
	s.invokingContext = params.Context

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
		// Match custom field name if it was set
		if !s.SkipCustomKeyRetrieval && s.fieldName != "" {
			s.fieldKey = helpers.FindCustomName(issue.Names.CustomFields, s.fieldName)

			if s.fieldKey == "" {
				return errors.New("failed to retrieve field key from specified custom field name")
			}
		}

		// Find the id of the status of the current Jira issue
		if issue.Fields.Status != nil {
			s.statusName = issue.Fields.Status.Name
		}

		if s.destination != "" {
			if file, err := resulthelper.CreateDestination(s.destination+"_"+s.issueId, "json"); err != nil {
				return errors.New("failed to create destination file")
			} else {
				switch s.invokingContext {
				case configuration.ReadStatus:
					if err := writeStatusToFile(file, s.issueId, s.statusName); err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}

func (s *ServiceReadIssue) Name() string {
	return "ServiceReadIssue"
}

func (s *ServiceReadIssue) ExecuteAsLastStep(params configuration.JiraAPIResourceParameters) error {
	//if file, err := resulthelper.CreateDestination(s.destination, "json"); err != nil {
	//	return err
	//} else {
	//	ctx := params.Context
	//
	//	switch ctx {
	//	case configuration.ReadIssue:
	//		vi := VersionReadIssueResponse{Issues: helpers.SliceToCommaSeparatedString(params.IssueList)}
	//		err := json.NewEncoder(file).Encode(InResponseIssue{
	//			Issues: vi,
	//		})
	//
	//		if err != nil {
	//			return err
	//		}
	//	}
	//}
	//
	return nil
}

func writeStatusToFile(file *os.File, issueId, statusName string) error {
	vi := VersionReadIssueResponse{Issues: issueId}
	mdf := assets.MetadataField{Name: "Status", Value: statusName}
	md := assets.Metadata{}
	md = append(md, mdf)
	err := json.NewEncoder(file).Encode(InResponseIssue{
		Issues:   vi,
		Metadata: md,
	})

	if err != nil {
		return err
	}

	return nil
}
