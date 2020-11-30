## week2学习笔记

1. `sentinel error`预定义错误，使用特定的值来表示错误，例如 `io.EOF`, 有些情况下会破坏相等性检查，可以使用errors.Cause()获取最底层的error
2. 可以自定义error type, 并通过类型断言来获取更多的上下文信息
3. `opaque error`用于获取
4. `errors.Is` 与sentinel error进行等值判断可能会出现不通过的情况，这时候应该使用Unwrap获取底层的error
5. `errors.As`用于将err转化为自定义的error type
6. 无错误的正常流程代码，将成为一条直线，而不是缩进的代码。
7. Error只处理1一次
8. 对error进行warp，可以添加更多的上下文信息，一般这种处理放在业务代码的底层，也就是与第三方的库交互的地方。
9. 推荐使用github.com/pkg/errors包来代替标准库的errors包, 该包可以通过`errors.Wrap(f)`和`errors.WithMessage`来添加上下文信息
10. error的Unwrap方法可以剥除上下文信息，显示更底层的error
11. 可以使用fmt.Printf("%+v", err)获取堆栈信息。
12. 参考代码：https://github.com/go-kratos/kratos/blob/v2/errors/errors.go
13. 多错误处理参考代码：https://github.com/uber-go/multierr

