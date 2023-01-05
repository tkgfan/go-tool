// Package structs
// author lby
// date 2022/8/10

package structs

import (
	"fmt"
	"reflect"
	"testing"
)

type S1 struct {
	Name  string
	S1Age int
}

type S2 struct {
	Name string
}

type S3 struct {
	S3Name string
}

type S4 struct {
	S3Arr []S3
}

type S5 struct {
	S3Arr []S3
}

func TestDeepCopy(t *testing.T) {
	src := []S1{{"s1", 11}, {"s11", 111}}
	//var dst []S2
	//err := CopyFields(&dst, src)
	//if err != nil || !equals(dst, src) {
	//	t.Error(err)
	//}

	src = append(src, S1{Name: "s3"})
	dst2 := [3]S2{}
	err2 := CopyFields(&dst2, src)
	if err2 != nil {
		t.Error(err2)
	}

	err3 := CopyFields(&dst2, src)
	if err3 != nil {
		t.Error(err3)
	}
	for i, _ := range dst2 {
		if dst2[i].Name != src[i].Name {
			t.Error(src[i], dst2[i])
		}
	}

	s4 := S4{
		S3Arr: []S3{{"n1"}, {"n2"}},
	}
	s5 := &S5{}
	err5 := CopyFields(s5, s4)
	if err5 != nil {
		fmt.Println(err5)
	}
}

func equals(dst []S2, src []S1) bool {
	for i, _ := range dst {
		if dst[i].Name != src[i].Name {
			return false
		}
	}
	return true
}

func TestCopyValue(t *testing.T) {
	s1 := S1{"s1Name", 11}
	s2 := S2{}
	err := copyValue(reflect.ValueOf(&s2), reflect.ValueOf(s1))
	if err != nil || s2.Name != s1.Name {
		t.Error(err, s1, s2)
	}
}

func TestCheckIfArrayOrSlice(t *testing.T) {
	src := make([]S2, 2)

	dst := [2]S1{}
	if isArrayOrSlice, _ := typeCheck(dst, src); !isArrayOrSlice {
		t.Error("isArrayOrSlice 应该为true，实际为false")
	}
	if isArrayOrSlice, _ := typeCheck(&dst, src); !isArrayOrSlice {
		t.Error("isArrayOrSlice 应该为true，实际为false")
	}
	dst2 := [1]S1{}
	if isArrayOrSlice, _ := typeCheck(dst2, src); !isArrayOrSlice {
		t.Error("isArrayOrSlice 应该为true，实际为false")
	}
	if isArrayOrSlice, _ := typeCheck(&dst2, src); !isArrayOrSlice {
		t.Error("isArrayOrSlice 应该为true，实际为false")
	}

	dst3 := S1{}
	if isArrayOrSlice, _ := typeCheck(dst3, src); isArrayOrSlice {
		t.Error("isArrayOrSlice 应该为false，实际为true")
	}
	if isArrayOrSlice, _ := typeCheck(&dst3, src); isArrayOrSlice {
		t.Error("isArrayOrSlice 应该为false，实际为true")
	}

	dst4 := [2]S2{}
	if _, err := typeCheck(dst4, src); err == nil {
		t.Error("error 不应该为空")
	}
	if _, err := typeCheck(&dst4, src); err != nil {
		t.Error("error 应该为空，但实际error为：", err)
	}
	dst5 := [1]S2{}
	if _, err := typeCheck(dst5, src); err == nil {
		t.Error("error 不应该为空")
	}
	if _, err := typeCheck(&dst5, src); err == nil {
		t.Error("error 不应该为空")
	}

}
