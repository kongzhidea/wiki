
## error处理

### error类型是一个接口类型，这是它的定义：

```
type error interface {
    Error() string
}
```

### 实现 error 接口类型来生成错误信息。

```
errors.New("math: square root of negative number")
```

### demo

```
func Sqrt(f float64) (float64, error) {
    if f < 0 {
        return 0, errors.New("math: square root of negative number")
    }
    return f * f, nil
}

result, err:= Sqrt(-1)

if err != nil {
   fmt.Println(err)
}
```