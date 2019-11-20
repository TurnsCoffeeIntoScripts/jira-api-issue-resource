package reading

import (
	"encoding/json"
	"errors"
)

type Names struct {
	CustomFields map[string]interface{}
}

func (n *Names) UnmarshalJSON(data []byte) error {
	var namesMap map[string]interface{}
	n.CustomFields = make(map[string]interface{})

	if n == nil {
		return errors.New("rawString: UnmarshalJSON on nil pointer")
	}

	if err := json.Unmarshal(data, &namesMap); err != nil {
		return err
	}

	for key, val := range namesMap {
		n.CustomFields[key] = val
	}

	return nil
}
