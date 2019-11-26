package authentication

type SessionInfo struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Session struct {
	Info SessionInfo `json:"session"`
}
