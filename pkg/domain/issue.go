package domain

type Issue struct {
    Key    string `json:"key"`
    Id     string `json:"id"`
    Fields Fields `json:"fields"`
}
