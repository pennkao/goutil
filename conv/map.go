package conv

func MapValueToString(data map[string]interface{}) map[string]string {
	newMap := make(map[string]string, len(data))
	for key, val := range data {
		valStr := ToString(val)
		newMap[key] = valStr
	}
	return newMap
}
