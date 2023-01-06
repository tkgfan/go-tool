// Package structs
// author lby
// date 2022/8/10

package structs

import (
	"reflect"
	"testing"
)

type (
	copyFieldsTest struct {
		name     string
		src      any
		dst      any
		expect   any
		hasError bool
	}

	people struct {
		Name          string
		Age           int
		Relationships []string
	}

	man struct {
		Name          string
		Sex           int8
		Age           int
		Relationships []string
	}

	woman struct {
		people
		Sex int8
	}

	animal struct {
		Name          int
		Age           int
		Relationships []int
	}
)

func TestCopyFields(t *testing.T) {
	tests := []copyFieldsTest{
		{
			name:     "值类型用例",
			src:      people{Name: "people", Age: 18},
			dst:      man{},
			hasError: true,
		},
		{
			name:     "正常复制结构体测试",
			src:      people{Name: "people", Age: 18},
			dst:      &man{},
			expect:   &man{Name: "people", Age: 18},
			hasError: false,
		},
		{
			name:     "切片复制到切片",
			src:      []people{{Name: "p1", Age: 11}, {Name: "p2", Age: 22}},
			dst:      &[]man{},
			expect:   &[]man{{Name: "p1", Age: 11}, {Name: "p2", Age: 22}},
			hasError: false,
		},
		{
			name:     "切片复制到数组",
			src:      []people{{Name: "p1", Age: 11}, {Name: "p2", Age: 22}},
			dst:      &[2]man{},
			expect:   &[2]man{{Name: "p1", Age: 11}, {Name: "p2", Age: 22}},
			hasError: false,
		},
		{
			name:     "切片复制到数组 2",
			src:      []people{{Name: "p1", Age: 11}, {Name: "p2", Age: 22}},
			dst:      &[3]man{},
			expect:   &[3]man{{Name: "p1", Age: 11}, {Name: "p2", Age: 22}},
			hasError: false,
		},
		{
			name:     "切片复制到容量不够的数组中",
			src:      []people{{Name: "p1", Age: 11}, {Name: "p2", Age: 22}},
			dst:      &[1]man{},
			hasError: true,
		},
		{
			name:     "切片复制到接口数组中",
			src:      []people{{Name: "p1", Age: 11}, {Name: "p2", Age: 22}},
			dst:      &[3]any{},
			expect:   &[3]any{people{Name: "p1", Age: 11}, people{Name: "p2", Age: 22}},
			hasError: false,
		},
		{
			name:     "类型不一致测试",
			src:      []people{{Name: "p1", Age: 11}},
			dst:      &people{},
			hasError: true,
		},
		{
			name:     "非结构体复制",
			src:      []int{1, 2},
			dst:      &[]int{},
			hasError: true,
		},
		{
			name:     "内嵌结构体复制",
			src:      woman{people: people{Name: "woman", Age: 11}, Sex: 1},
			dst:      &people{},
			expect:   &people{Name: "woman", Age: 11},
			hasError: false,
		},
		{
			name:     "字段名相同类型不一致",
			src:      people{Name: "people", Age: 11},
			dst:      &animal{},
			expect:   &animal{Age: 11},
			hasError: false,
		},
		{
			name:     "被复制实例不是结构体",
			src:      11,
			dst:      &man{},
			hasError: true,
		},
		{
			name:     "字段存在数组类型",
			src:      people{Relationships: []string{"alice", "bob"}},
			dst:      &man{},
			expect:   &man{Relationships: []string{"alice", "bob"}},
			hasError: false,
		},
		{
			name:     "字段存在类型不相同的数组类型",
			src:      people{Relationships: []string{"alice", "bob"}},
			dst:      &animal{},
			expect:   &animal{},
			hasError: false,
		},
	}

	executeTests(tests, t)
}

func executeTests(tests []copyFieldsTest, t *testing.T) {
	for _, test := range tests {
		err := CopyFields(test.dst, test.src)
		if test.hasError {
			if err == nil {
				t.Errorf("%s ,需要抛出异常", test.name)
			}
		} else {
			if err != nil {
				t.Errorf("%s , %+v", test.name, err)
			}
			if !reflect.DeepEqual(test.expect, test.dst) {
				t.Errorf("%s , expect=%+v, got=%+v", test.name, test.expect, test.dst)
			}
		}
	}
}

func BenchmarkCopyFields(b *testing.B) {
	src := people{Name: "people", Age: 18, Relationships: []string{"alice", "bob"}}
	dst := &man{}
	for i := 0; i < b.N; i++ {
		CopyFields(dst, src)
	}
}
