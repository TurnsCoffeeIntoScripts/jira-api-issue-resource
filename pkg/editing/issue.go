package editing

type Issue struct {
	Fields map[string]interface{} `json:"fields"`
}

func (i *Issue) AddField(key string, val interface{}) {
	if i.Fields == nil {
		i.Fields = make(map[string]interface{})
	}

	i.Fields[key] = val
}
