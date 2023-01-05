// Package structs
// author gmfan
// date 2022/8/31

package structs

import (
	"errors"
	"fmt"
	"reflect"
)

// CopyFields 深度复制结构体src到dst，会复制名称相同且类型一致的部分的
// src字段到dst中。如果src为切片或数组，则dst必须也为切片或数组且长度要
// 一致。
func CopyFields(dst, src any) (err error) {
	// 防止意外的panic
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(fmt.Sprintf("%#v", e))
		}
	}()

	// dst与src为数组或切片
	if isArrOrSlice, e := checkIfArrayOrSlice(dst, src); isArrOrSlice || e != nil {
		if isArrOrSlice && e == nil {
			e = copyArrayOrSlice(dst, src)
		}

		return e
	}

	// 普通结构体
	dstVal := reflect.ValueOf(dst)
	if dstVal.Kind() != reflect.Pointer {
		return errors.New("dst必须为Pointer")
	}

	srcVal := reflect.ValueOf(src)

	return copyValue(dstVal, srcVal)
}

// 复制数组与切片
func copyArrayOrSlice(dst, src any) (err error) {
	dstVal := elemIfPointer(reflect.ValueOf(dst))
	srcVal := elemIfPointer(reflect.ValueOf(src))

	for i := 0; i < srcVal.Len(); i++ {
		if err = copyValue(dstVal.Index(i), srcVal.Index(i)); err != nil {
			return
		}
	}

	return
}

// 复制src的值到dst
func copyValue(dst, src reflect.Value) (err error) {
	dst = elemIfPointer(dst)
	if dst.Kind() != reflect.Struct {
		return errors.New("dst不为结构体")
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
			// stf可能是匿名结构体
			if stf.Type.Kind() == reflect.Struct {
				err = copyValue(dst, src.Field(i))
			}
			continue
		}
		// 如果dst是interface则直接复制
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

// 类型检查，如果dst与src是数组或切片两者长度必须一致
func checkIfArrayOrSlice(dst, src any) (isArrOrSlice bool, err error) {
	srcVal := reflect.ValueOf(src)
	srcVal = elemIfPointer(srcVal)
	if srcVal.Kind() != reflect.Array && srcVal.Kind() != reflect.Slice {
		dstVal := reflect.ValueOf(dst)
		dstVal = elemIfPointer(dstVal)
		if dstVal.Kind() == reflect.Array || dstVal.Kind() == reflect.Slice {
			return false, errors.New("src不为数组或切片，但dst为数组或切片")
		}
		return
	}

	dstVal := reflect.ValueOf(dst)
	if dstVal.Kind() == reflect.Array {
		return true, errors.New("dst必须为数组的Pointer")
	}

	dstVal = elemIfPointer(dstVal)
	if dstVal.Kind() != reflect.Array && dstVal.Kind() != reflect.Slice {
		return false, errors.New("src为数组或切片，但是dst不为数组或切片")
	}

	if dstVal.Len() != srcVal.Len() {
		return true, errors.New("src长度与dst长度不一致")
	}

	return true, nil
}

// 如果v是Pointer则返回其指向的值
func elemIfPointer(v reflect.Value) reflect.Value {
	if v.Kind() == reflect.Pointer {
		return v.Elem()
	}
	return v
}
