package configuration

import (
	"bytes"
	"strings"
)

const IssuePlaceholder = "??ISSUE??"

func BuildJsonBodyFromString(strs ...string) []byte {
	var buffer bytes.Buffer
	buffer.WriteString("{")
	for idx := 0; idx < len(strs) - 1; idx += 2 {
		buffer.WriteString("\"" + strs[idx] + "\"")
		buffer.WriteString(":")
		buffer.WriteString("\"" + strs[idx + 1] + "\",")
	}

	stringBuffer := strings.TrimRight(buffer.String(), ",")
	buffer = *bytes.NewBufferString(stringBuffer)

	buffer.WriteString("}")

	return buffer.Bytes()
}
