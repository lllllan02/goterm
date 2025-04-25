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

## 图表功能

GoTerm 支持在终端中绘制简单的图表：

### 条形图

```go
// 创建并配置条形图
bar := goterm.NewBarChart().
    SetTitle("编程语言流行度").
    AddData("Go", 75).
    AddData("Python", 95).
    AddData("Java", 83)

// 自定义样式
bar.SetBarStyle(goterm.New().Green())
bar.SetTitleStyle(goterm.New().Bold().Blue())

// 打印图表
bar.Print()
```

### 饼图

```go
// 创建并配置饼图
pie := goterm.NewPieChart().
    SetTitle("市场份额").
    AddData("产品A", 35).
    AddData("产品B", 25).
    AddData("产品C", 20).
    AddData("其他", 20)

// 自定义样式
pie.SetTitleStyle(goterm.New().Bold().Magenta())
pie.SetStyle("产品A", goterm.New().Red())
pie.SetStyle("产品B", goterm.New().Green())

// 打印图表
pie.Print()
```

### 折线图

```go
// 创建并配置折线图
line := goterm.NewLineChart().
    SetTitle("增长趋势").
    SetAxisTitles("时间", "数值").
    SetWidth(60).
    SetHeight(15)

// 添加数据点
for i := 0; i < 10; i++ {
    line.AddPoint("系列1", float64(i), float64(i*i))
}

// 自定义样式
line.SetLineStyle("系列1", goterm.New().Blue())
line.SetMarkerStyle("系列1", goterm.New().Bold().Blue())

// 打印图表
line.Print()
```

## 详细示例

- 样式示例: [examples/style/main.go](examples/style/main.go)
- 图表示例: [examples/charts.go](examples/charts.go) (包含条形图、饼图和折线图)

**运行图表示例:**

```bash
# 运行所有图表示例
go run examples/charts.go

# 只运行条形图示例
go run examples/charts.go bar

# 只运行饼图示例
go run examples/charts.go pie

# 只运行折线图示例
go run examples/charts.go line
```

## 许可证

MIT 