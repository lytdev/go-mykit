package gnum

import (
	"encoding/json"
	"strconv"
)

// AnyToInt 任意数据转int
func AnyToInt(value interface{}) int {
	if value == nil {
		return 0
	}
	switch val := value.(type) {
	case int:
		return val
	case int8:
		return int(val)
	case int16:
		return int(val)
	case int32:
		return int(val)
	case int64:
		return int(val)
	case uint:
		return int(val)
	case uint8:
		return int(val)
	case uint16:
		return int(val)
	case uint32:
		return int(val)
	case uint64:
		return int(val)
	case *string:
		v, err := strconv.Atoi(*val)
		if err != nil {
			return 0
		}
		return v
	case string:
		v, err := strconv.Atoi(val)
		if err != nil {
			return 0
		}
		return v
	case float32:
		return int(val)
	case float64:
		return int(val)
	case bool:
		if val {
			return 1
		} else {
			return 0
		}
	case json.Number:
		v, _ := val.Int64()
		return int(v)
	}
	return 0
}
