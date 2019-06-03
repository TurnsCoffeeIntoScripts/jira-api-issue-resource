package configuration

import "errors"

func (conf *JiraAPIResourceConfiguration) ValidateBaseParameters() (bool, []error) {
	var errList []error
	var ok bool
	success := true

	ok, errList = checkEmptyConfigElement(conf.Parameters.JiraApiUrl, "missing api URL", errList)
	success = success && ok

	ok, errList = checkEmptyConfigElement(conf.Parameters.Username, "missing api username", errList)
	success = success && ok

	ok, errList = checkEmptyConfigElement(conf.Parameters.Password, "missing api password", errList)
	success = success && ok

	return success, errList
}

func checkEmptyConfigElement(elem *string, msg string, errList []error) (bool, []error) {
	if *elem == "" {
		return false, append(errList, errors.New(msg))
	}

	return true, errList
}
