// Package errors
// author gmfan
// date 2023/2/24
package errors

import (
	"encoding/json"
	stderrors "errors"
	"testing"
)

func TestIsStackError(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want bool
	}{
		{
			name: "Nil test",
			err:  nil,
			want: false,
		},
		{
			name: "普通 error",
			err:  stderrors.New("test"),
			want: false,
		},
		{
			name: "stackError",
			err:  New("stackError"),
			want: true,
		},
	}

	for _, tt := range tests {
		_, got := tt.err.(*stackError)
		if got != tt.want {
			t.Errorf("IsStackError(): name: %s, got: %v, want: %v", tt.name, got, tt.want)
		}
	}
}

func TestStackError_Json(t *testing.T) {
	tests := []error{
		New("new"),
		Wrap(New("warp")),
		Wrapf(New("wrapf"), "format %d", 1),
	}

	for _, tt := range tests {
		got := tt.(Jsoner).Json()
		if !json.Valid([]byte(got)) {
			t.Errorf("stackError.Json(): 输出字符串不是 JSON 格式, got: %s", got)
		}
	}
}

//goos: windows
//goarch: amd64
//cpu: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz
//BenchmarkStackError_Json
//BenchmarkStackError_Json-8       1749685               579.2 ns/op
func BenchmarkStackError_Json(b *testing.B) {
	err := Wrapf(New("json test"), "wrap message")
	st, _ := err.(*stackError)
	for i := 0; i < b.N; i++ {
		st.Json()
	}
}
