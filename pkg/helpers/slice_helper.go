package helpers

import "strings"

func SliceToCommaSeparatedString(slice []string) string {
	if slice == nil || len(slice) == 0 {
		return "."
	}

	var result string
	for _, s := range slice {
		result += s
		result += ","
	}

	// Remove trailing comma for clarity
	return strings.TrimSuffix(result, ",")
}
