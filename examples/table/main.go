package main

import (
	"fmt"

	"github.com/lllllan02/goterm"
)

func main() {
	fmt.Println("基本表格示例：")
	showBasicTable()

	fmt.Println("\n无边框表格：")
	showBorderlessTable()

	fmt.Println("\n蓝色边框表格：")
	showBlueTable()

	fmt.Println("\n绿色内容表格：")
	showGreenContentTable()

	fmt.Println()
	showTitledTable()

	fmt.Println("\n带有最大宽度限制的表格：")
	showMaxWidthTable()

	fmt.Println("\n完全分离样式的表格示例：")
	showStyledTable()

	fmt.Println("\n带行分隔线的表格示例：")
	showRowSeparatorTable()
}

// 创建基本表格
func createBasicTable() *goterm.Table {
	table := goterm.NewEmptyTable()
	table.AddColumn(goterm.NewColumn("姓名").SetAlignment(goterm.AlignLeft).SetMinWidth(6))
	table.AddColumn(goterm.NewColumn("年龄").SetAlignment(goterm.AlignRight).SetMinWidth(4))
	table.AddColumn(goterm.NewColumn("城市").SetAlignment(goterm.AlignCenter).SetMinWidth(8))
	table.AddColumn(goterm.NewColumn("职业").SetAlignment(goterm.AlignLeft).SetMinWidth(12))

	// 添加数据
	table.AddRow("张三", "28", "北京", "软件工程师")
	table.AddRow("李四", "32", "上海", "产品经理")
	table.AddRow("王五", "25", "广州", "UI设计师")
	table.AddRow("赵六", "30", "深圳", "数据分析师")
	table.AddRow("陈七", "35", "杭州", "架构师")

	return table
}

// 显示基本表格
func showBasicTable() {
	table := createBasicTable()
	table.Print()
}

// 显示无边框表格
func showBorderlessTable() {
	table := createBasicTable()
	table.HasBorder = false
	table.Print()
}

// 显示蓝色边框表格
func showBlueTable() {
	table := createBasicTable()
	table.Border.Blue()
	table.PrintWithStyle()
}

// 显示绿色内容表格
func showGreenContentTable() {
	table := createBasicTable()
	table.Border.Blue()
	table.Header.Green().Bold()
	table.PrintWithStyle()
}

// 显示带标题的表格
func showTitledTable() {
	table := createBasicTable()
	title := "员工信息表"
	titleStyle := goterm.New().Magenta().Bold().Underline()

	// 设置青色边框
	table.Border.Cyan()

	goterm.PrintStyledTable(table, title, titleStyle, nil)
}

// 显示带有最大宽度限制的表格
func showMaxWidthTable() {
	// 创建新表格并添加带最大宽度限制的列
	table := goterm.NewEmptyTable()
	table.AddColumn(goterm.NewColumn("标题").SetAlignment(goterm.AlignLeft).SetMinWidth(15).SetMaxWidth(15))
	table.AddColumn(goterm.NewColumn("描述").SetAlignment(goterm.AlignLeft).SetMinWidth(25).SetMaxWidth(25))
	table.AddColumn(goterm.NewColumn("发布日期").SetAlignment(goterm.AlignCenter).SetMinWidth(10))
	table.AddColumn(goterm.NewColumn("状态").SetAlignment(goterm.AlignRight).SetMinWidth(8))

	// 添加一些长文本数据
	table.AddRow(
		"这是一个很长的标题，会被截断",
		"这是一段非常长的描述文本，演示最大宽度限制功能，超出部分会被截断",
		"2023-10-15",
		"已发布",
	)
	table.AddRow(
		"短标题",
		"短描述",
		"2023-10-16",
		"草稿",
	)
	table.AddRow(
		"中等长度的标题",
		"这是一个中等长度的描述，不会被截断",
		"2023-10-17",
		"审核中",
	)
	table.AddRow(
		"另一个长标题示例",
		"另一个很长的描述，测试文本截断和对齐效果是否正确",
		"2023-10-18",
		"已完成",
	)

	// 设置样式并打印
	table.Border.RGB(100, 100, 255)
	table.Header.RGB(255, 100, 200).Bold()
	table.PrintWithStyle()
}

// 显示带样式的表格
func showStyledTable() {
	table := goterm.NewEmptyTable()
	table.AddColumn(goterm.NewColumn("ID").SetAlignment(goterm.AlignCenter).SetMinWidth(5))
	table.AddColumn(goterm.NewColumn("名称").SetAlignment(goterm.AlignLeft).SetMinWidth(15))
	table.AddColumn(goterm.NewColumn("状态").SetAlignment(goterm.AlignCenter).SetMinWidth(10))

	table.AddRow("001", "示例项目 A", "进行中")
	table.AddRow("002", "示例项目 B", "已完成")
	table.AddRow("003", "示例项目 C", "未开始")

	// 设置不同部分的样式
	table.Border.Blue()          // 蓝色边框
	table.Header.Yellow().Bold() // 黄色粗体表头
	table.Row.Green()            // 绿色行数据

	// 打印表格
	table.PrintWithStyle()
}

// 显示带行分隔线的表格
func showRowSeparatorTable() {
	table := goterm.NewEmptyTable()
	table.AddColumn(goterm.NewColumn("ID").SetAlignment(goterm.AlignCenter).SetMinWidth(5))
	table.AddColumn(goterm.NewColumn("名称").SetAlignment(goterm.AlignLeft).SetMinWidth(15))
	table.AddColumn(goterm.NewColumn("状态").SetAlignment(goterm.AlignCenter).SetMinWidth(10))

	table.AddRow("001", "示例项目 A", "进行中")
	table.AddRow("002", "示例项目 B", "已完成")
	table.AddRow("003", "示例项目 C", "未开始")
	table.AddRow("004", "示例项目 D", "规划中")
	table.AddRow("005", "示例项目 E", "已暂停")

	// 设置样式
	table.Border.Red()             // 红色边框
	table.Header.Cyan().Bold()     // 青色表头
	table.SetHasRowSeparator(true) // 显示行分隔线

	// 打印表格
	table.PrintWithStyle()
}
