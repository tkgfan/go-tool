// author lby
// date 2023/2/24

package errors

import (
	"fmt"
	"github.com/tkgfan/go-tool/structs"
	"strings"
)

// Cause 返回由 errors.New、errors.Wrap、errors.Wrapf 中包裹的 cause error。
// error 如果不是上述方法产生的则会返回其本身
func Cause(err error) error {
	if e, ok := err.(*stackError); ok {
		return e.cause
	}
	return err
}

// Wrap 返回包含堆栈信息的 error
func Wrap(err error) error {
	if structs.IsNil(err) {
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
	if structs.IsNil(err) {
		return nil
	}

	remark := fmt.Sprintf(format, args...)
	if st, ok := err.(*stackError); ok {
		st.stacks = append(st.stacks, caller(remark))
		return st
	}

	return &stackError{
		cause:  err,
		stacks: []*stack{caller(remark)},
	}
}

type Jsoner interface {
	Json() string
}

const (
	defaultCauseJsonSize = len(`{"cause":""}`)
)

var (
	causeJsonPrefix = []byte(`{"cause":"`)
	causeJsonEnd    = []byte(`"}`)
)

// Json 返回 error 序列化为 JSON 的字符串，如果 error 实现了 Json 方法则
// 调用 Json 方法返回字符串。未实现 Json 方法则返回:
//     {"cause":"err.Error()"}
func Json(err error) string {
	// 自定义实现 Json 方法
	if e, ok := err.(Jsoner); ok {
		return e.Json()
	}
	// 默认 Json 序列化
	cause := err.Error()
	builder := strings.Builder{}
	builder.Grow(len(cause) + defaultCauseJsonSize)
	builder.Write(causeJsonPrefix)
	builder.WriteString(err.Error())
	builder.Write(causeJsonEnd)
	return builder.String()
}
