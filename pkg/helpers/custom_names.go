package helpers

func FindCustomName(customFields map[string]interface{}, searchVal string) string {
	for key := range customFields {
		if customFields[key] == searchVal {
			return key
		}
	}

	return ""
}
