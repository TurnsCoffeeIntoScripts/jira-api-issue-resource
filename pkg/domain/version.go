/*
Package 'domain' provides struct that corresponds to Jira REST API objects
*/
package domain

// Simplified representation of the 'version' object from the Jira REST API.
type Version struct {
	Name string `json:"name"`
}
