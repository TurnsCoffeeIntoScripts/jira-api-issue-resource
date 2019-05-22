package rest

import "github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/configuration"

type Metadata struct {
	Url        string
	Protocol   string
	HttpMethod string
	Username   string
	Password   string
}

func (m *Metadata) AuthenticatedUrl() string {
	return m.Protocol + "://" + m.Username + ":" + m.Password + "@" + m.Url
}

func (m *Metadata) Initialize(flags configuration.JiraApiResourceFlags) {
	m.Url = *flags.JiraApiUrl
	m.Protocol = *flags.Protocol
	m.Username = *flags.Username
	m.Password = *flags.Password
}
