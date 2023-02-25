// author lby
// date 2023/2/24

package errors

import (
	stderrors "errors"
	"runtime"
	"strconv"
	"strings"
)

const (
	// 默认 stack 结构体序列化所需内存大小
	defaultGrowStackSize = len(`{"file":"","line":10,"remark":""}`)
	// 默认 stackError 结构体除 stacks 所需内存大小
	defaultGrowStackErrorSize = len(`{"file":"","stacks":[]}`)
	// 假设默认 remark 大小为 20
	defaultGrowRemarkSize = 20
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

// New 创建一个包含堆栈信息的 error
func New(msg string) error {
	return &stackError{
		cause:  stderrors.New(msg),
		stacks: []*stack{caller("")},
	}
}

func (s *stackError) Error() string {
	return s.cause.Error()
}

// Json 将堆栈错误信息转换成 JSON 格式的字符串
func (s *stackError) Json() string {
	var builder strings.Builder
	// 内存预分配
	cause := s.cause.Error()
	g := len(s.stacks[0].file) + defaultGrowRemarkSize + defaultGrowStackSize
	g *= len(s.stacks)
	g = g + len(cause) + defaultGrowStackErrorSize
	builder.Grow(g)

	// 构建 JSON 字符串
	builder.WriteString(`{"cause":"`)
	// 优化性能
	builder.WriteString(cause)
	builder.WriteString(`","stacks":[`)
	for i := 0; i < len(s.stacks); i++ {
		builder.WriteString(`{"file":"`)
		// 优化性能
		builder.WriteString(s.stacks[i].file)
		builder.WriteString(`","line":` + strconv.Itoa(s.stacks[i].line) + `,` +
			`"remark":"` + s.stacks[i].remark + `"}`)
		if i != len(s.stacks)-1 {
			builder.WriteString(",")
		}
	}
	builder.WriteString(`]}`)

	return builder.String()
}
