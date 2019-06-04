package configuration

import (
	"errors"
	"reflect"
)

// ValidateBaseParameters checks if the mandatory parameters aren't specified, in which case a new error is appended
// to a list which is then returned along the success status (bool)
func (conf *JiraAPIResourceConfiguration) ValidateBaseParameters() bool {
	conf.Errors = checkEmptyConfigElement(conf.Parameters.JiraAPIUrl, "missing api --url", conf.Errors)
	conf.Errors = checkEmptyConfigElement(conf.Parameters.Username, "missing api --username", conf.Errors)
	conf.Errors = checkEmptyConfigElement(conf.Parameters.Password, "missing api --password", conf.Errors)
	conf.Errors = checkOneInManyNotEmtpy("Need only one of '--issue-id', '--issue-list' and '--issue-script'", conf.Errors,
		conf.Parameters.IssueID, conf.Parameters.IssueList, conf.Parameters.IssueScript)

	return len(conf.Errors) == 0
}

func (conf *JiraAPIResourceConfiguration) ValidateContextParameters() bool {
	callValidateContextMethod(conf, conf.Flags.ContextFlags.CtxAddLabel)
	callValidateContextMethod(conf, conf.Flags.ContextFlags.CtxComment)

	return len(conf.Errors) == 0
}

func (conf *JiraAPIResourceConfiguration) ValidateShowHelpContext() {}

func (conf *JiraAPIResourceConfiguration) ValidateAddLabelContext() {
	conf.Errors = checkEmptyConfigElement(conf.Parameters.Label, "missing --label", conf.Errors)
}

func (conf *JiraAPIResourceConfiguration) ValidateCommentContext() {
	conf.Errors = checkEmptyConfigElement(conf.Parameters.Body, "missing --body", conf.Errors)
}

func callValidateContextMethod(conf *JiraAPIResourceConfiguration, def JiraAPIResourceContextFlagDefinition) {
	if *def.Value {
		method := reflect.ValueOf(conf).MethodByName(def.Func)
		method.Call(nil)
	}
}

func checkEmptyConfigElement(elem *string, msg string, errList []error) []error {
	if *elem == "" {
		return append(errList, errors.New(msg))
	}

	return errList
}

func checkOneInManyNotEmtpy(msg string, errList []error, elements ...*string) []error {
	foundOne := false
	foundMany := false
	for _, elem := range elements {
		if *elem != "" {
			if !foundOne {
				foundOne = true
				continue
			}

			if foundOne {
				foundMany = true
				continue
			}
		}
	}

	if foundMany || (!foundOne && !foundMany) {
		return append(errList, errors.New(msg))
	}

	return errList
}
