package configuration

type Metadata struct {
	Url                   string
	Protocol              string
	HttpMethod            string
	Username              string
	Password              string
	ResourceConfiguration JiraAPIResourceConfiguration
}

func (m *Metadata) AuthenticatedUrl() string {
	return m.Protocol + "://" + m.Username + ":" + m.Password + "@" + m.Url + "/rest/api/2"
}

func (m *Metadata) Initialize(conf JiraAPIResourceConfiguration) {
	m.Url = *conf.Parameters.JiraAPIUrl
	m.Protocol = *conf.Parameters.Protocol
	m.Username = *conf.Parameters.Username
	m.Password = *conf.Parameters.Password
	m.ResourceConfiguration = conf
}
