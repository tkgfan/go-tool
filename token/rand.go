// Package token
// author gmfan
// date 2023/2/25
package token

import (
	"crypto/rand"
	"io"
)

const CharacterSequence = "0123456789qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"
const N = 62

// Rand 生成随机 token 字符串，len 为 token 长度。
// 随机字符从 CharacterSequence 选取。CharacterSequence 含有 62 个字符
func Rand(len int) string {
	// 获取随机序列
	ridx := make([]byte, len)
	_, err := io.ReadFull(rand.Reader, ridx)
	if err != nil {
		panic(err)
	}

	// 随机序列转换成 byte 切片
	bs := make([]byte, len)
	for i := 0; i < len; i++ {
		bs[i] = CharacterSequence[ridx[i]%N]
	}
	return string(bs)
}
