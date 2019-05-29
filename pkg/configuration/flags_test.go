package configuration

import (
	"os"
	"testing"
)

var dummyUrl = "URL"
var dummyUsername = "USER"
var dummyPassword = "PASSWORD"
var dummyIssue = "ISSUE-1234"

var data = JiraApiResourceFlags{}

func TestMain(m *testing.M) {
	data.SetupFlags(false)
	code := m.Run()
	os.Exit(code)
}

func TestValidateBaseFlags_Success(t *testing.T) {
	*data.JiraApiUrl = dummyUrl
	*data.Username = dummyUsername
	*data.Password = dummyPassword
	*data.IssueId = dummyIssue

	ok := data.ValidateBaseFlags()
	resetData()

	if !ok {
		t.Errorf("JiraApiResourceFlags validation should have been successful but was not. (Expected: %t, got: %t)", true, ok)
	}
}

func TestValidateBaseFlags_FailureURL(t *testing.T) {
	ok := data.ValidateBaseFlags()
	resetData()

	if ok {
		t.Errorf("JiraApiResourceFlags validation should not have been successful but was. (Expected: %t, got: %t)", false, ok)
	}
}

func resetData()  {
	*data.ShowHelp = false
	*data.JiraApiUrl = ""
	*data.Protocol = ""
	*data.Username = ""
	*data.Password = ""
	*data.IssueId = ""
	*data.Body = ""
}
