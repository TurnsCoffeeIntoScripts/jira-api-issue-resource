package rest

type Metadata struct {
	Url string
	Protocol string
	HttpMethod string
	Username string
	Password string
}

func (m *Metadata) AuthenticatedUrl() string{
	return m.Protocol + "://" + m.Username + ":" + m.Password + "@" + m.Url
}