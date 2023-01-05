package helper

import (
	"encoding/json"
	"strconv"
)

// ToStr converts any value to string.
func ToStr(v interface{}) string {
	if v == nil {
		return ""
	}
	switch v := v.(type) {
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	case int32:
		return strconv.Itoa(int(v))
	case int64:
		return strconv.FormatInt(v, 10)
	case bool:
		return strconv.FormatBool(v)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case []uint8:
		return string(v)
	default:
		resultJSON, err := json.Marshal(v)
		if err != nil {
			return ""
		}
		return string(resultJSON)
	}
}
