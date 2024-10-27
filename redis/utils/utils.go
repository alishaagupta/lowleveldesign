package utils

import (
	"fmt"
	"reflect"
	"strconv"
)

func ConvertToString(val interface{}) string {

	switch val.(type) {
	case string:
		return val.(string)
	case bool:
		return strconv.FormatBool(val.(bool))
	case int:
		return fmt.Sprintf("%v", val)
	case float64:

		return fmt.Sprintf("%v", val)
	default:
		fmt.Println("ff ", reflect.TypeOf(val))
		return ""
	}
}
