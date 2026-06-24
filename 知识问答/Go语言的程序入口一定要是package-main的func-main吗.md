# Go 语言的程序入口一定要是 package main 的 func main 吗？

## 简短结论

如果要写一个可以直接运行的 Go 程序，入口必须在 `package main` 中，并且函数必须是：

```go
func main()
```

也就是说，可执行程序需要：

```go
package main

func main() {
}
```

## 详细说明

Go 中有两类常见包：

- `package main`：用于构建可执行程序。
- 普通包：用于被其他代码导入和复用。

只有 `package main` 会被 Go 工具链当作可执行程序入口包。这个包中需要定义无参数、无返回值的 `func main()`，程序运行时会从这里开始执行。

普通包不需要 `func main()`。例如工具函数、业务函数、公共库代码，都可以放在其他包中。

测试文件也不需要自己写 `func main()`。执行 `go test` 时，Go 测试工具会负责生成测试入口，并运行符合规则的测试函数。

## 易误解点

不是所有 Go 文件都需要 `func main()`。

只有“要直接运行成程序”的 `package main` 需要入口函数。写普通包、测试代码、库代码时，不需要 `func main()`。

## 记忆点

- 可执行程序：`package main` + `func main()`。
- 普通包：不需要 `func main()`。
- 测试代码：不需要自己写 `func main()`。
- `func main()` 不能有参数，也不能有返回值。
