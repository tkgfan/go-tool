// Package errors
// author gmfan
// date 2023/2/25
package errors

import (
	stderrors "errors"
	"testing"
)

func TestCause(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want error
	}{
		{
			name: "普通 error",
			err:  stderrors.New("t1"),
			want: stderrors.New("t1"),
		},
		{
			name: "errors.New",
			err:  New("errors.New"),
			want: stderrors.New("errors.New"),
		},
		{
			name: "errors.Wrap",
			err:  Wrap(stderrors.New("errors.Wrap")),
			want: stderrors.New("errors.Wrap"),
		},
		{
			name: "errors.Wrapf",
			err:  Wrapf(stderrors.New("errors.Wrapf"), "format"),
			want: stderrors.New("errors.Wrapf"),
		},
	}

	for _, tt := range tests {
		got := Cause(tt.err)
		if got.Error() != tt.want.Error() {
			t.Errorf("Cause(): name: %s, got: %s, want: %s", tt.name, got, tt.want)
		}
	}
}

func TestWrap(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want error
	}{
		{
			name: "Nil 用例",
			err:  nil,
			want: nil,
		},
		{
			name: "普通 error",
			err:  stderrors.New("foo"),
			want: stderrors.New("foo"),
		},
		{
			name: "stackError",
			err:  New("stackError"),
			want: New("stackError"),
		},
	}

	for _, tt := range tests {
		got := Wrap(tt.err)
		if got == nil || tt.want == nil {
			if got != tt.want {
				t.Errorf("Wrap(): name: %s, got: %v, want: %v", tt.name, got, tt.want)
			}
			continue
		}
		if got.Error() != tt.want.Error() {
			t.Errorf("Wrap(): name: %q, got: %q, want: %q", tt.name, got, tt.want)
		}
	}
}

func TestWrapf(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want error
	}{
		{
			name: "Nil 用例",
			err:  nil,
			want: nil,
		},
		{
			name: "普通 error",
			err:  stderrors.New("foo"),
			want: stderrors.New("foo"),
		},
		{
			name: "stackError",
			err:  New("stackError"),
			want: New("stackError"),
		},
	}

	for _, tt := range tests {
		got := Wrapf(tt.err, "")
		if got == nil || tt.want == nil {
			if got != tt.want {
				t.Errorf("Wrap(): name: %s, got: %v, want: %v", tt.name, got, tt.want)
			}
			continue
		}
		if got.Error() != tt.want.Error() {
			t.Errorf("Wrap(): name: %q, got: %q, want: %q", tt.name, got, tt.want)
		}
	}
}

func TestJson(t *testing.T) {
	st := New("implemented Jsoner interface")
	tests := []struct {
		name string
		err  error
		want string
	}{
		{
			name: "自定义实现了 Json 方法",
			err:  st,
			want: st.(Jsoner).Json(),
		},
		{
			name: "没有实现 Json 方法",
			err:  stderrors.New("not implemented Jsoner interface"),
			want: `{"cause":"not implemented Jsoner interface"}`,
		},
	}

	for _, tt := range tests {
		got := Json(tt.err)
		if got != tt.want {
			t.Errorf("Json(): name: %q, got: %q, want: %q", tt.name, got, tt.want)
		}
	}
}
