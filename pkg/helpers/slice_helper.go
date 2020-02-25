package helpers

func CleanStringSlice(slice []string) []string {
	var r = make([]string, 0)
	for _, str := range slice {
		if str != "" {
			r = append(r, str)
		}
	}

	return r
}

func SliceToCommaSeparatedString(slice []string) string {
	if slice == nil || len(slice) == 0 {
		return "."
	}

	var result string
	for _, s := range slice {
		result += s
		result += ","
	}

	return result
}
