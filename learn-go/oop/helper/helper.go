package helper

import "encoding/json"

func StructToMap(data any) map[string]any {
	conv, _ := json.Marshal(data)
	var result map[string]any
	json.Unmarshal(conv, &result)
	return result
}
