// author lby
// date 2023/1/30

package slices

import "reflect"

// ToInterfaceSlice 将任意类型的 val 转换为 interface 切片。
// 1. 如果 val 为切片或数组，则转换为相同长度的 interface 切片。
// 2. 如果 val 非切片与数组类型，则转换为长度为1的 interface 切片。
func ToInterfaceSlice(val any) (res []any) {
	vv := reflect.ValueOf(val)
	if vv.Kind() != reflect.Array && vv.Kind() != reflect.Slice {
		return []any{val}
	}

	res = make([]any, vv.Len())
	for i := 0; i < len(res); i++ {
		res[i] = vv.Index(i).Interface()
	}
	return
}
