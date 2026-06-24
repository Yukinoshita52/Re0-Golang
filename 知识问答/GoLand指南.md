# GoLand 指南

本文件用于记录学习 Go 时遇到的 GoLand 使用问题和解决经验。

## `main_test.go` 中同包函数爆红，但 `go test` 可以通过

### 问题现象

Go 的测试文件，也就是 `*_test.go` 文件，异常爆红，提示某个方法找不到，例如：

```text
Unresolved reference 'greeting'
```

但是命令行执行 `go test` 又可以通过。

### 解决方式

```text
GoLand 可能抽风了，go test 都没问题但是就是爆红，那就重启 GoLand。
```
