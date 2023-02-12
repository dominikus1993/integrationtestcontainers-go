package common

func GetValueOrDefault(value, defaultValue string) string {
	if value == "" {
		return defaultValue
	}
	return value
}
