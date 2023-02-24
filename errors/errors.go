// author lby
// date 2023/2/24

package errors

import (
	"fmt"
)

// Wrap 返回包含堆栈信息的 error
func Wrap(err error) error {
	if err == nil {
		return nil
	}
	if se, ok := err.(*stackError); ok {
		se.stacks = append(se.stacks, caller(""))
		return se
	}
	return &stackError{
		cause:  err,
		stacks: []*stack{caller("")},
	}
}

// Wrapf 返回包含堆栈信息的 error。format 格式化信息会保存到堆栈信息中
func Wrapf(err error, format string, args ...any) error {
	if err == nil {
		return nil
	}

	remark := fmt.Sprintf(format, args)
	if st, ok := err.(*stackError); ok {
		st.stacks = append(st.stacks, caller(remark))
		return st
	}

	return &stackError{
		cause:  err,
		stacks: []*stack{caller(remark)},
	}
}
