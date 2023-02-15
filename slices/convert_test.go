// author lby
// date 2023/2/7

package slices

import (
	"reflect"
	"testing"
)

type (
	toInterfaceSliceTest struct {
		name   string
		arg    any
		expect any
	}

	cat struct {
		Name string
		Age  int
	}
)

func TestToInterfaceSlice(t *testing.T) {
	tests := []toInterfaceSliceTest{
		{
			name:   "单个元素测试",
			arg:    cat{Name: "alice", Age: 16},
			expect: []any{cat{Name: "alice", Age: 16}},
		},
		{
			name:   "Nil 测试",
			arg:    nil,
			expect: []any{},
		},
		{
			name:   "切片转换",
			arg:    []cat{{Name: "alice", Age: 16}, {Name: "tkg", Age: 10}},
			expect: []any{cat{Name: "alice", Age: 16}, cat{Name: "tkg", Age: 10}},
		},
		{
			name:   "数组转换",
			arg:    [2]cat{{Name: "alice", Age: 16}, {Name: "tkg", Age: 10}},
			expect: []any{cat{Name: "alice", Age: 16}, cat{Name: "tkg", Age: 10}},
		},
	}

	for _, test := range tests {
		action := ToInterfaceSlice(test.arg)
		if !reflect.DeepEqual(action, test.expect) {
			t.Errorf("%s,arg=%+v, expect=%+v, got=%+v", test.name, test.arg, test.expect, action)
		}
	}
}
