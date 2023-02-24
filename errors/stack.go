// author lby
// date 2023/2/24

package errors

import (
	stderrors "errors"
	"fmt"
	"runtime"
)

// 堆栈信息结构
type stack struct {
	pc     uintptr
	file   string
	line   int
	remark string
}

func caller(remark string) *stack {
	st := &stack{
		remark: remark,
	}
	st.pc, st.file, st.line, _ = runtime.Caller(2)
	return st
}

// stackError 包含错误的堆栈信息
type stackError struct {
	cause  error
	stacks []*stack
}

func New(msg string) error {
	return &stackError{
		cause:  stderrors.New(msg),
		stacks: []*stack{caller("")},
	}
}

func (s *stackError) Error() string {
	return s.cause.Error()
}

func (s *stackError) Format(state fmt.State, verb rune) {

}
