package configuration

type Metadata struct {
	Url           string
	Protocol      string
	HttpMethod    string
	Username      string
	Password      string
	ResourceFlags JiraAPIResourceFlags
}

func (m *Metadata) AuthenticatedUrl() string {
	return m.Protocol + "://" + m.Username + ":" + m.Password + "@" + m.Url + "/rest/api/2"
}

func (m *Metadata) Initialize(flags JiraAPIResourceFlags) {
	/*m.Url = *flags.JiraApiUrl
	m.Protocol = *flags.Protocol
	m.Username = *flags.Username
	m.Password = *flags.Password*/
	m.ResourceFlags = flags
}
