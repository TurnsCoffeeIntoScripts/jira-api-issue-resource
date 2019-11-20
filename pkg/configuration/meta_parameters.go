package configuration

type MetaParameters struct {
	parsed           bool // Indicates whether the flag.Parse() method was called or not
	mandatoryPresent bool // Indicates whether the mandatory parameters are persent or not
	valid            bool // Indicates if the values received for each flags are valid or not

	MultipleIssue bool // Indicates whether the resource received multiple Jira issue to process
}

func (meta *MetaParameters) AllMandatoryValuesPresent() bool {
	return meta.parsed && meta.mandatoryPresent
}

func (meta *MetaParameters) Ready() bool {
	return meta.parsed && meta.valid
}
