// Package token
// author gmfan
// date 2023/2/25
package token

import (
	"testing"
)

func TestRand(t *testing.T) {
	set := make(map[string]struct{})
	for j := 0; j < 100000; j++ {
		token := Rand(8)
		if len(token) != 8 {
			t.Error("GenRandToken 函数生成 Token 长度应为", 8, "实际为", len(token))
		}
		if _, ok := set[token]; ok {
			t.Error("GenRandToken 函数生成 Token 冲突")
		}
		set[token] = struct{}{}
	}
}

//goos: windows
//goarch: amd64
//cpu: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz
//BenchmarkRand
//BenchmarkRand-8          1594503               638.5 ns/op
func BenchmarkRand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Rand(18)
	}
}
