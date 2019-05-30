package domain

type Issue struct {
	Key    string `json:"key"`
	Id     string `json:"id"`
	Fields Fields `json:"fields"`
}

func (i *Issue) HasParent() bool {
	return i.Fields.Parent != nil
}

func (i *Issue) GetParent() *Issue {
	return i.Fields.Parent
}
