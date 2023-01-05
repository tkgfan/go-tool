// Package structs
// author gmfan
// date 2022/8/31

package structs

import (
	"errors"
	"fmt"
	"reflect"
)

// CopyFields 复制结构体 src 到 dst 中，会尽可能复制名称相同的字段。
// 如果 src 为切片或数组，则 dst 必须也为切片或数组。需要注意的是 dst
// 必须为指针类型
func CopyFields(dst, src any) (err error) {
	// 防止意外的panic
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(fmt.Sprintf("%+v", e))
		}
	}()

	// dst 与 src 为数组或切片
	if isArrOrSlice, e := typeCheck(dst, src); isArrOrSlice || e != nil {
		if isArrOrSlice && e == nil {
			e = copyArrayOrSlice(dst, src)
		}

		return e
	}

	srcVal := reflect.ValueOf(src)
	dstVal := reflect.ValueOf(dst)

	return copyValue(dstVal, srcVal)
}

// 复制数组与切片
func copyArrayOrSlice(dst, src any) (err error) {
	dstValRes := reflect.ValueOf(dst)
	dstVal := elemIfPointer(dstValRes)
	srcVal := elemIfPointer(reflect.ValueOf(src))

	elementType := dstVal.Type().Elem()

	for i := 0; i < srcVal.Len(); i++ {
		// 扩容
		if dstVal.Len() == i {
			newElem := reflect.New(elementType).Elem()
			dstVal = reflect.Append(dstVal, newElem)
		}

		if err = copyValue(dstVal.Index(i), srcVal.Index(i)); err != nil {
			return
		}
	}

	if dstVal.Kind() == reflect.Array {
		dstValRes.Elem().Set(dstVal)
	} else {
		dstValRes.Elem().Set(dstVal.Slice(0, srcVal.Len()))
	}
	return
}

// 复制src的值到dst
func copyValue(dst, src reflect.Value) (err error) {
	dst = elemIfPointer(dst)
	if dst.Kind() != reflect.Struct {
		return errors.New("dst 不为结构体")
	}

	src = elemIfPointer(src)
	if src.Kind() != reflect.Struct {
		return errors.New("src不为结构体")
	}

	srcType := src.Type()
	for i := 0; i < src.NumField(); i++ {
		stf := srcType.Field(i)
		df := dst.FieldByName(stf.Name)
		if ok := df.IsValid(); !ok {
			// stf 可能是匿名结构体
			if stf.Type.Kind() == reflect.Struct {
				err = copyValue(dst, src.Field(i))
			}
			continue
		}
		// 如果 dst 是 interface 则直接复制
		if df.Kind() != reflect.Interface {
			// 类型不一致
			if df.Kind() != src.Field(i).Kind() {
				continue
			}

			// 如果字段是数组或切片需要检查元素类型是否一致，否则跳过
			if src.Field(i).Kind() == reflect.Array || src.Field(i).Kind() == reflect.Slice {
				if src.Field(i).Type() != df.Type() {
					continue
				}
			}
		}

		df.Set(src.Field(i))
	}

	return nil
}

// 类型检查，如果 dst 与 src 都为数组或切片则 isArrOrSlice 为 true。
// 如果两个类型不可复制则抛出异常
func typeCheck(dst, src any) (isArrOrSlice bool, err error) {
	// dst 必须为指针类型
	dstType := reflect.ValueOf(dst)
	if dstType.Kind() != reflect.Pointer {
		return false, errors.New("dst 必须为 Pointer")
	}

	dstType = dstType.Elem()
	srcType := reflect.TypeOf(src)

	// src 为数组或切片
	if srcType.Kind() == reflect.Array || srcType.Kind() == reflect.Slice {
		if dstType.Kind() != reflect.Array && dstType.Kind() != reflect.Slice {
			return false, errors.New("src 为数组或切片，dst 也必须是数组或切片")
		}

		if dstType.Kind() == reflect.Array {
			srcVal := elemIfPointer(reflect.ValueOf(src))
			if dstType.Len() < srcVal.Len() {
				return true, errors.New("dst 为数组时长度必须大于等于 src 的长度")
			}
		}
		return true, nil
	}

	return false, nil
}

// 如果v是Pointer则返回其指向的值
func elemIfPointer(v reflect.Value) reflect.Value {
	if v.Kind() == reflect.Pointer {
		return v.Elem()
	}
	return v
}
