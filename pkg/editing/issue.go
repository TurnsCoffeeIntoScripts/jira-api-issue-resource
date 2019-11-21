// See editing/service.go for this package's comment
package editing

// This struct is a representation of a Jira issue in the context of editing it.
type Issue struct {
	Fields map[string]interface{} `json:"fields"`
}

// This method adds a field in the existing map of an issue before submitting this edit to the Jira API.
func (i *Issue) AddField(key string, val interface{}) {
	if i.Fields == nil {
		i.Fields = make(map[string]interface{})
	}

	i.Fields[key] = val
}
