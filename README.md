## gutil
功能包函数

## time
```go
// 输入混合时间，可正确解析，例如1h1m1s
// 输出time.Time类型 等于 time.Hour + time.Minute + time.Second
time.ParseTime("1h1m1s")
```
## file
```go
// 输入混合大小，可正确解析，例如1B1Kb1Mb1G
// 输出字节数
file.ParseSize("1B1Kb1Mb1G")
```
## jwt
```go
// 生成token
secret := "12345"
token, err := GenToken(time.Second, "test", secret)

// 解析token
j, err := ParseToken(token, secret)
```
