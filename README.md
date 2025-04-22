# GoTerm

GoTerm 是一个简单的 Go 语言终端样式库，用于输出彩色和样式化的文本。

## 安装

```bash
go get github.com/lllllan02/goterm
```

## 基本用法

```go
fmt.Println(goterm.Red("红色文本"))

goterm.StyleRed.Println("红色文本")

goterm.New().Bold().Red().Println("粗体红色文本")
```

## 可用样式

- 文本样式: `Bold()`, `Italic()`, `Underline()`
- 颜色: `Red()`, `Green()`, `Blue()`, `Yellow()`, `Cyan()`, `Magenta()`, `White()`, `Black()`
- 背景色: `BgRed()`, `BgGreen()`, 等同颜色名称

## 详细示例

查看 [examples/style/main.go](examples/style/main.go) 获取更多使用示例。

## 许可证

MIT 