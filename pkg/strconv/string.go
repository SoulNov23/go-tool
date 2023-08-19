package strconv

import (
	"fmt"
	"strings"
	"unsafe"

	"github.com/soulnov23/go-tool/pkg/json"
)

// https://github.com/golang/go/issues/53003
func BytesToString(b []byte) string {
	return unsafe.String(&b[0], len(b))
}

func StringToBytes(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

func StringToMap(data string, fieldSep string, valueSep string) map[string]string {
	recordMap := map[string]string{}
	fieldSlice := strings.Split(data, fieldSep)
	for _, kv := range fieldSlice {
		valueSlice := strings.Split(kv, valueSep)
		if len(valueSlice) == 2 {
			recordMap[valueSlice[0]] = valueSlice[1]
		} else if len(valueSlice) == 1 && strings.Count(kv, valueSep) == 1 {
			recordMap[valueSlice[0]] = ""
		}
	}
	return recordMap
}

func MapToString(recordMap map[string]string) string {
	var builder strings.Builder
	for key, value := range recordMap {
		builder.WriteString(key + "=" + value + "&")
	}
	builder.Len()
	return builder.String()[0 : builder.Len()-1]
}

func AnyToString(row any) string {
	switch v := row.(type) {
	case nil:
		return ""
	case *string:
		return fmt.Sprintf("%v", *v)
	case *bool:
		return fmt.Sprintf("%v", *v)
	case *uint8:
		return fmt.Sprintf("%v", *v)
	case *uint16:
		return fmt.Sprintf("%v", *v)
	case *uint32:
		return fmt.Sprintf("%v", *v)
	case *uint64:
		return fmt.Sprintf("%v", *v)
	case *int8:
		return fmt.Sprintf("%v", *v)
	case *int16:
		return fmt.Sprintf("%v", *v)
	case *int32:
		return fmt.Sprintf("%v", *v)
	case *int64:
		return fmt.Sprintf("%v", *v)
	case *float32:
		return fmt.Sprintf("%v", *v)
	case *float64:
		return fmt.Sprintf("%v", *v)
	case *int:
		return fmt.Sprintf("%v", *v)
	case *uint:
		return fmt.Sprintf("%v", *v)
	case *[]byte:
		return fmt.Sprintf("%v", *v)
	case string, bool, uint8, uint16, uint32, uint64, int8, int16, int32, int64, float32, float64, int, uint:
		return fmt.Sprintf("%v", v)
	case []byte:
		return string(v)
	case *struct{}:
		result, err := json.Marshal(*v)
		if err != nil {
			return ""
		}
		return string(result)
	case *any:
		return AnyToString(*v)
	case any:
		switch vv := v.(type) {
		case string, bool, uint8, uint16, uint32, uint64, int8, int16, int32, int64, float32, float64, int, uint:
			return fmt.Sprintf("%v", vv)
		case []byte:
			return string(vv)
		default:
			return fmt.Sprintf("%v", vv)
		}
	default:
		return fmt.Sprintf("%v", v)
	}
}
