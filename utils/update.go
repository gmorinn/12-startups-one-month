package utils

func UpdateString(old *string, input *string) string {
	if input == nil || *input == "" {
		if old == nil {
			return ""
		} else {
			return *old
		}
	}
	return *input
}
