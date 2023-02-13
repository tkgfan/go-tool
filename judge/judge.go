// author lby
// date 2023/2/13

package judge

import "reflect"

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
