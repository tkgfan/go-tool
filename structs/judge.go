// author lby
// date 2023/2/15

package structs

import "reflect"

// IsNil 判断 val 是否为 nil
func IsNil(val any) bool {
	if val == nil {
		return true
	}

	vv := reflect.ValueOf(val)
	if vv.Kind() == reflect.Pointer {
		return vv.IsNil()
	}

	return false
}
