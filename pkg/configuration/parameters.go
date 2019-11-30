/*
Package configuration provides the structs that contains this application's Go flags
and other custom parameters determined from the Go flags. Parsing of said flags is
done via the JiraAPIResourceParameters.Parse() method

To initialize the JiraAPIResourceParameters one only needs to do the following:

	// Short declaration of the variable
	params := configuration.JiraAPIResourceParameters{}

	// Execution of the 'Parse' method
	params.Parse()

After that specific method, the flags.Parse method of the Go lang flags api will have been called.

The package also defines a wide array of constants representing the flags name in the command line.
*/
package configuration

import (
	"flag"
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/log"
	"strings"
)

// Definition of constants that are use for the flags setup
const (
	// Parameters
	jiraAPIURL       = "url"
	username         = "username"
	password         = "password"
	context          = "context"
	issueList        = "issues"
	customFieldName  = "customFieldName"
	customFieldType  = "customFieldType"
	customFieldValue = "customFieldValue"
	loggingLevel     = "loggingLevel"

	// Flags
	secured       = "secured"
	forceOnParent = "forceOnParent"

	// Default values and descriptions for both paramaters and flags
	jiraAPIURLDefault           = ""
	jiraAPIURLDescription       = "The base URL of the Jira API"
	usernameDefault             = ""
	usernameDescription         = "The username used to connect to the Jira API"
	passwordDefault             = ""
	passwordDescription         = "The password needed to connect to the Jira API"
	contextDefault              = ""
	contextDescription          = "The context of execution. {'EditCustomField'}"
	issueListDefault            = ""
	issueListDescription        = "The issue or list of issues to execute the specified context to"
	customFieldNameDefault      = ""
	customFieldNameDescription  = "Certain operation (such as edits) might require the user to specify the name of the custome field so that the resource may find the appropriate custom field"
	customFieldTypeDefault      = "string"
	customFieldTypeDescription  = "The type that is required by the field via the Jira API"
	customFieldValueDefault     = ""
	customFieldValueDescription = "The value of the field that will be updated (in case of update workflow)"
	loggingLevelDefault         = "INFO"
	loggingLevelDescription     = "The level of the loggers of the application {'ALL', 'DEBUG', 'ERROR', 'INFO', 'WARN', 'OFF'}"
	_                           = /*securedDefault*/ true
	securedDescription          = "Flags that indicates if the API calls should be made in session"
	_                           = /*forceOnParentDefault*/ false
	forceOnParentDescription    = "Flags that indicates if we want to force all operation on the parent issue (if there's one)"
)

// JiraAPIResourceParameters is a struct that holds every possible parameters/flags known by the application.
// They are all parsed via the flag.Parse() method of the Go flags api. The struct also contains the meta-parameters
// (see meta-parametes.go).
type JiraAPIResourceParameters struct {
	JiraAPIUrl       *string
	Username         *string
	Password         *string
	Context          Context
	IssueList        []string
	CustomFieldName  *string
	CustomFieldType  *string
	CustomFieldValue *string
	LoggingLevel     *string

	ActiveIssue string         // The **SINGLE** issue that the resource is currently processing
	Meta        MetaParameters //
	Flags       JiraAPIResourceFlags
}

// This struct is used to separate the parameters (which takes values in the command line) of the flags (which don't).
// That being said the values in this struct are still parsed via the Go flags api. They've been put 'aside' for clariry
// purposes.
type JiraAPIResourceFlags struct {
	Secured       *bool
	ForceOnParent *bool
}

// Method that initialize every parameters/flags and makes the actual call the flag.Parse(). A few custom operation are
// performed afterward such as initialization and validation.
func (param *JiraAPIResourceParameters) Parse() {
	var contextString *string
	var issueListString *string

	// Parsing parameters
	param.JiraAPIUrl = flag.String(jiraAPIURL, jiraAPIURLDefault, jiraAPIURLDescription)
	param.Username = flag.String(username, usernameDefault, usernameDescription)
	param.Password = flag.String(password, passwordDefault, passwordDescription)
	contextString = flag.String(context, contextDefault, contextDescription)
	issueListString = flag.String(issueList, issueListDefault, issueListDescription)
	param.CustomFieldName = flag.String(customFieldName, customFieldNameDefault, customFieldNameDescription)
	param.CustomFieldType = flag.String(customFieldType, customFieldTypeDefault, customFieldTypeDescription)
	param.CustomFieldValue = flag.String(customFieldValue, customFieldValueDefault, customFieldValueDescription)
	param.LoggingLevel = flag.String(loggingLevel, loggingLevelDefault, loggingLevelDescription)

	// Parsing flags
	param.Flags.Secured = flag.Bool(secured, true, securedDescription)
	param.Flags.ForceOnParent = flag.Bool(forceOnParent, false, forceOnParentDescription)

	if !param.Meta.parsed {
		flag.Parse()
		param.Meta.parsed = flag.Parsed()
	}

	param.initializeContext(contextString)
	param.initializeIssueList(issueListString)
	param.initLogger()
	param.validate()
}

func (param *JiraAPIResourceParameters) validate() {
	// By default both meta flags are set to true
	param.Meta.mandatoryPresent = true
	param.Meta.valid = true

	if *param.JiraAPIUrl == "" || *param.Username == "" || *param.Password == "" || param.IssueList == nil {
		// In this case we are missing one or more mandatory parameters
		// This also causes the input parameters to not be valid
		param.Meta.mandatoryPresent = false
		param.Meta.valid = false
	} else if len(param.IssueList) == 0 {
		// This is the case where the issue's list was passed but it was empty
		param.Meta.valid = false
	} else if param.Context == Unknown {
		// The specified context wasn't recognized, therefore it isn't valid
		param.Meta.valid = false
	}
}

func (param *JiraAPIResourceParameters) initializeContext(contextString *string) {
	if contextString == nil {
		param.Context = Unknown
	} else {
		param.Context = GetContext(*contextString)
	}
}

func (param *JiraAPIResourceParameters) initializeIssueList(issueListString *string) {
	if *issueListString != "" {
		param.IssueList = strings.Split(*issueListString, ",")

		// More than 1 issue specified will set the 'Multiple' flag to true
		param.Meta.MultipleIssue = len(param.IssueList) > 1
	}
}

func (param *JiraAPIResourceParameters) initLogger() {
	log.Logger = log.ResourceLogger{}
	log.Logger.InitLoggerFromParam(*param.LoggingLevel)
}