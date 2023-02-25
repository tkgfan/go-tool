# go-tool

Go 实用工具

## structs

- CopyFields(dst, src any) (err error): 将 src 上的字段复制到 dst 上。
- IsNil(val any) bool: 判断 val 是否为 nil。

## slices

- ToInterfaceSlice(val any) (res []any): 将 val 转换为 interface 切片。

## errors

- New(msg string) error: 创建一个包含堆栈信息的 error。
- Wrap(err error) error: 返回包含栈信息的 error。
- Wrapf(err error, format string, args ...any) error: 返回包含栈信息的 error。
- Cause(err error) error: 返回由 errors.New、errors.Wrap、errors.Wrapf 中包裹的 cause error。普通 error 则返回其本身。
- Json(err error) string: 返回 error 序列化为 JSON 的字符串，如果 error 实现了 Json 方法则调用 Json 方法返回字符串。

## token

- Rand(len int) string: 生成随机 token 字符串，len 为 token 长度。

