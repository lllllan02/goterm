# GoTerm

GoTerm 是一个功能丰富的 Go 语言终端库，用于创建美观、交互式的命令行应用程序。它提供了文本样式、图表、进度条、表格、交互式组件等多种功能。

## 安装

```bash
go get github.com/lllllan02/goterm
```

## 主要功能

### 1. 文本样式

支持多种文本样式和颜色：

```go
// 简单颜色文本
fmt.Println(goterm.Red("红色文本"))

// 使用样式对象
goterm.StyleRed.Println("红色文本")

// 链式调用组合样式
goterm.New().Bold().Red().Println("粗体红色文本")
```

可用样式：
- 文本样式: `Bold()`, `Italic()`, `Underline()`
- 文本颜色: `Red()`, `Green()`, `Blue()`, `Yellow()`, `Cyan()`, `Magenta()`, `White()`, `Black()`
- 背景颜色: `BgRed()`, `BgGreen()`, `BgBlue()` 等

### 2. 图表功能

#### 条形图

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

#### 饼图

```go
pie := goterm.NewPieChart().
    SetTitle("市场份额").
    AddData("产品A", 35).
    AddData("产品B", 25).
    AddData("产品C", 40)

// 自定义样式
pie.SetTitleStyle(goterm.New().Bold().Magenta())
pie.Print()
```

#### 折线图

```go
line := goterm.NewLineChart().
    SetTitle("增长趋势").
    SetAxisTitles("时间", "数值").
    SetWidth(60).
    SetHeight(15)

// 添加数据点
for i := 0; i < 10; i++ {
    line.AddPoint("系列1", float64(i), float64(i*i))
}

line.Print()
```

### 3. 进度条

```go
// 创建进度条
bar := goterm.NewProgressBar().
    SetTotal(100).
    SetWidth(50).
    SetStyle(goterm.New().Green())

// 更新进度
for i := 0; i <= 100; i++ {
    bar.Update(i)
    time.Sleep(50 * time.Millisecond)
}
```

### 4. 表格

```go
// 创建表格
table := goterm.NewTable().
    SetHeaders("ID", "姓名", "年龄").
    AddRow("1", "张三", "25").
    AddRow("2", "李四", "30").
    AddRow("3", "王五", "28")

// 自定义样式
table.SetHeaderStyle(goterm.New().Bold().Green())
table.SetBorderStyle(goterm.New().Blue())

// 打印表格
table.Print()
```

### 5. 交互式组件

```go
// 选择菜单
options := []string{"选项1", "选项2", "选项3"}
selected, _ := goterm.Select("请选择一个选项:", options)
fmt.Println("你选择了:", selected)

// 输入框
name, _ := goterm.Input("请输入你的名字:", "")
fmt.Println("你好,", name)

// 确认对话框
confirm, _ := goterm.Confirm("确定要继续吗?", false)
if confirm {
    fmt.Println("继续操作...")
} else {
    fmt.Println("操作已取消")
}
```

### 6. 动画效果

```go
// 打字效果
goterm.NewTyper().
    SetText("这是一段带有打字效果的文本").
    SetDelay(100 * time.Millisecond).
    SetStyle(goterm.New().Green()).
    Start()

// 彩虹文字效果
goterm.NewRainbow().
    SetText("彩虹文字效果").
    Start()

// 加载动画
spinner := goterm.NewSpinner().
    SetStyle(goterm.New().Cyan()).
    Start()
time.Sleep(3 * time.Second)
spinner.Stop()
```

### 7. 树形结构

```go
// 创建树形结构
tree := goterm.NewTree().
    SetRoot("根节点").
    AddBranch("根节点", "分支1").
    AddBranch("根节点", "分支2").
    AddBranch("分支1", "叶子1").
    AddBranch("分支1", "叶子2")

// 设置样式
tree.SetRootStyle(goterm.New().Bold().Green())
tree.SetBranchStyle(goterm.New().Blue())
tree.SetLeafStyle(goterm.New().Yellow())

// 打印树
tree.Print()
```

### 8. 日志系统

```go
// 创建日志器
logger := goterm.NewLogger()

// 输出不同级别的日志
logger.Info("这是一条信息日志")
logger.Warning("这是一条警告日志")
logger.Error("这是一条错误日志")
logger.Debug("这是一条调试日志")
```

### 9. 光标控制

```go
// 移动光标
goterm.MoveCursor(5, 10) // 移动到第5行，第10列
goterm.ClearScreen()     // 清屏
goterm.CursorUp(2)       // 上移2行
goterm.CursorDown(3)     // 下移3行
```

## 示例代码

查看完整示例代码：

- 文本样式: [examples/style/](examples/style/)
- 图表: [examples/charts/](examples/charts/)
- 进度条: [examples/progress/](examples/progress/)
- 表格: [examples/table/](examples/table/)
- 交互式组件: [examples/interactive/](examples/interactive/)
- 动画效果: [examples/animation/](examples/animation/)
- 树形结构: [examples/tree/](examples/tree/)
- 日志系统: [examples/logger/](examples/logger/)
- 光标控制: [examples/cursor/](examples/cursor/)

## 许可证

MIT 