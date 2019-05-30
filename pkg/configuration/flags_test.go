package configuration

import (
	"os"
	"testing"
)

var (
	// Dummy data to be inserted in the JiraApiResourceFlags struct to test the validation. The value of the data
	// themselves aren't significant, it merely matters whether or not they're present.
	dummyUrl       = "URL"
	dummyUsername  = "USER"
	dummyPassword  = "PASSWORD"
	dummyIssue     = "ISSUE-1234"
	dummyIssueList = "ISSUE-1234,ISSUE-5678"
)

// Local instance of the type JiraApiResourceFlags, all test will be performed with this variable.
var data = JiraApiResourceFlags{}

// The 'TestMain' method is used, so it's made possible to set up our flags, including the parsing operation of the
// 'flag' package.
func TestMain(m *testing.M) {
	data.SetupFlags(false)
	code := m.Run()
	os.Exit(code)
}

// This test represents a successful validation of the JiraApiResourceFlags instance, which means that the 'JiraApiUrl',
// 'Username' and 'Password' are all set. The others are optional.
func TestValidateBaseFlags_Success(t *testing.T) {
	*data.JiraApiUrl = dummyUrl
	*data.Username = dummyUsername
	*data.Password = dummyPassword

	ok := data.ValidateBaseFlags()
	resetData()

	if !ok {
		t.Errorf("JiraApiResourceFlags validation should have been successful but was not. (Expected: %t, got: %t)", true, ok)
	}
}

// This test represents a failed validation of the JiraApiResourceFlags instance. In this case, the 'JiraApiUrl' field
// was omitted.
func TestValidateBaseFlags_FailureURL(t *testing.T) {
	*data.Username = dummyUsername
	*data.Password = dummyPassword

	ok := data.ValidateBaseFlags()
	resetData()

	if ok {
		t.Errorf("JiraApiResourceFlags validation should not have been successful but was. (Expected: %t, got: %t)", false, ok)
	}
}

// This test represents a failed validation of the JiraApiResourceFlags instance. In this case, the 'Username' field
// was omitted.
func TestValidateBaseFlags_FailureUsername(t *testing.T) {
	*data.JiraApiUrl = dummyUrl
	*data.Password = dummyPassword

	ok := data.ValidateBaseFlags()
	resetData()

	if ok {
		t.Errorf("JiraApiResourceFlags validation should not have been successful but was. (Expected: %t, got: %t)", false, ok)
	}
}

// This test represents a failed validation of the JiraApiResourceFlags instance. In this case, the 'Password' field
// was omitted.
func TestValidateBaseFlags_FailurePassword(t *testing.T) {
	*data.JiraApiUrl = dummyUrl
	*data.Username = dummyUsername

	ok := data.ValidateBaseFlags()
	resetData()

	if ok {
		t.Errorf("JiraApiResourceFlags validation should not have been successful but was. (Expected: %t, got: %t)", false, ok)
	}
}

// This test checks that the flags representing the number of issues are set correctly in the context of no Jira issue
// being specified.
func TestValidateBaseFlags_ZeroIssue(t *testing.T) {
	*data.JiraApiUrl = dummyUrl
	*data.Username = dummyUsername
	*data.Password = dummyPassword

	data.ValidateBaseFlags()
	zeroIssue := data.ZeroIssue

	if !zeroIssue {
		t.Errorf("JiraApiRessourceFlags with no 'IssueId' or 'RawIssueList' should have 'ZeroIssue' at true, instead it was %t", zeroIssue)
	}
}

// This test checks that the flags representing the number of issues are set correctly in the context of a single Jira
// issue being specified.
func TestValidateBaseFlags_SingleIssue(t *testing.T) {
	*data.JiraApiUrl = dummyUrl
	*data.Username = dummyUsername
	*data.Password = dummyPassword
	*data.IssueId = dummyIssue

	data.ValidateBaseFlags()
	singleIssue := data.SingleIssue

	if !singleIssue {
		t.Errorf("JiraApiRessourceFlags with an 'IssueId' and no 'RawIssueList' should have 'SingleIssue' at true, instead it was %t", singleIssue)
	}
}

// This test checks that the flags representing the number of issues are set correctly in the context of multiple Jira
// issues being specified.
func TestValidateBaseFlags_MultipleIssue(t *testing.T) {
	*data.JiraApiUrl = dummyUrl
	*data.Username = dummyUsername
	*data.Password = dummyPassword
	*data.RawIssueList = dummyIssueList

	data.ValidateBaseFlags()
	multipleIssue := data.MultipleIssue

	if !multipleIssue {
		t.Errorf("JiraApiRessourceFlags with no 'IssueId' but a 'RawIssueList' should have 'MultipleIssue' at true, instead it was %t", multipleIssue)
	}
}

// Local method that resets the JiraApiResourceFlags fields to their respective zero value.
func resetData() {
	*data.ShowHelp = false
	*data.JiraApiUrl = ""
	*data.Protocol = ""
	*data.Username = ""
	*data.Password = ""
	*data.IssueId = ""
	*data.RawIssueList = ""
	*data.Body = ""
}
