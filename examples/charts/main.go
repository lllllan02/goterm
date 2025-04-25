package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/lllllan02/goterm"
)

func main() {
	if len(os.Args) > 1 {
		arg := strings.ToLower(os.Args[1])
		switch arg {
		case "bar", "条形图":
			showBarCharts()
		case "pie", "饼图":
			showPieCharts()
		case "line", "折线图":
			showLineCharts()
		default:
			showAllCharts()
		}
	} else {
		showAllCharts()
	}
}

func showAllCharts() {
	fmt.Println("GoTerm 图表库示例")
	fmt.Println("===============================")

	// 条形图
	fmt.Println("\n条形图示例")
	fmt.Println("-------------------------------")
	showBarCharts()

	// 饼图
	fmt.Println("\n\n饼图示例")
	fmt.Println("-------------------------------")
	showPieCharts()

	// 折线图
	fmt.Println("\n\n折线图示例")
	fmt.Println("-------------------------------")
	showLineCharts()
}

// ==================== 条形图示例 ====================

func showBarCharts() {
	// 基本条形图
	basicBarChart()

	fmt.Println()
	// 自定义样式的条形图
	customStyledBarChart()
}

func basicBarChart() {
	fmt.Println("基本条形图:")
	fmt.Println("----------------")

	// 创建条形图
	bar := goterm.NewBarChart().
		SetTitle("编程语言流行度").
		AddData("Go", 75).
		AddData("Python", 95).
		AddData("Java", 83).
		AddData("JavaScript", 91).
		AddData("Rust", 62).
		AddData("C++", 70)

	// 打印条形图
	bar.Print()
}

func customStyledBarChart() {
	fmt.Println("自定义样式的条形图:")
	fmt.Println("----------------")

	// 创建条形图
	bar := goterm.NewBarChart().
		SetTitle("季度销售额").
		AddData("Q1", 120).
		AddData("Q2", 240).
		AddData("Q3", 180).
		AddData("Q4", 310)

	// 自定义样式
	bar.SetTitleStyle(goterm.New().Bold().Magenta())
	bar.SetLabelStyle(goterm.New().Bold().Yellow())
	bar.SetValueStyle(goterm.New().Bold().Green())
	bar.SetBarStyle(goterm.New().Cyan())
	bar.SetBarChar("■")

	// 打印条形图
	bar.Print()
}

// ==================== 饼图示例 ====================

func showPieCharts() {
	// 基本饼图
	basicPieChart()

	fmt.Println()
	// 小尺寸饼图
	smallPieChart()
}

func basicPieChart() {
	fmt.Println("基本饼图:")
	fmt.Println("----------------")

	// 创建饼图
	pie := goterm.NewPieChart().
		SetTitle("市场份额").
		AddData("产品A", 35).
		AddData("产品B", 25).
		AddData("产品C", 20).
		AddData("产品D", 15).
		AddData("其他", 5)

	// 自定义样式
	pie.SetTitleStyle(goterm.New().Bold().Underline().Magenta())
	pie.SetStyle("产品A", goterm.New().Red())
	pie.SetStyle("产品B", goterm.New().Green())
	pie.SetStyle("产品C", goterm.New().Blue())
	pie.SetStyle("产品D", goterm.New().Yellow())
	pie.SetStyle("其他", goterm.New().Cyan())

	// 打印饼图
	pie.Print()
}

func smallPieChart() {
	fmt.Println("小尺寸饼图:")
	fmt.Println("----------------")

	// 创建小尺寸饼图
	pie := goterm.NewPieChart().
		SetTitle("月度支出").
		SetSize(10). // 更小的尺寸
		AddData("住房", 30).
		AddData("食物", 25).
		AddData("交通", 15).
		AddData("娱乐", 10).
		AddData("储蓄", 20)

	// 打印饼图
	pie.Print()
}

// ==================== 折线图示例 ====================

func showLineCharts() {
	// 简单折线图
	simpleLineChart()

	fmt.Println()
	// 多系列折线图
	multiSeriesLineChart()

	fmt.Println()
	// 三角函数图
	trigFunctionsChart()
}

func simpleLineChart() {
	fmt.Println("简单折线图:")
	fmt.Println("----------------")

	// 创建折线图
	line := goterm.NewLineChart().
		SetTitle("简单的增长曲线").
		SetAxisTitles("时间", "数值").
		SetHeight(12)

	// 添加数据点
	for i := 0; i < 20; i++ {
		x := float64(i)
		y := float64(i * i / 8)
		line.AddPoint("增长", x, y)
	}

	// 自定义样式
	line.SetTitleStyle(goterm.New().Bold().Underline().Blue())
	line.SetLineStyle("增长", goterm.New().Green())
	line.SetMarkerStyle("增长", goterm.New().Bold().Green())

	// 打印折线图
	line.Print()
}

func multiSeriesLineChart() {
	fmt.Println("多系列折线图:")
	fmt.Println("----------------")

	// 创建折线图
	line := goterm.NewLineChart().
		SetTitle("不同产品的销售趋势").
		SetAxisTitles("月份", "销售量").
		SetWidth(70).
		SetHeight(15)

	// 产品A的销售数据
	data := []float64{12, 19, 15, 8, 22, 30, 25, 19, 23, 15, 18, 26}
	for i, value := range data {
		line.AddPoint("产品A", float64(i+1), value)
	}

	// 产品B的销售数据
	dataB := []float64{8, 5, 10, 14, 18, 21, 19, 26, 28, 22, 16, 12}
	for i, value := range dataB {
		line.AddPoint("产品B", float64(i+1), value)
	}

	// 产品C的销售数据
	dataC := []float64{5, 11, 14, 18, 21, 25, 16, 18, 15, 13, 12, 10}
	for i, value := range dataC {
		line.AddPoint("产品C", float64(i+1), value)
	}

	// 自定义样式
	line.SetLineStyle("产品A", goterm.New().Red())
	line.SetLineStyle("产品B", goterm.New().Blue())
	line.SetLineStyle("产品C", goterm.New().Green())

	// 打印折线图
	line.Print()
}

func trigFunctionsChart() {
	fmt.Println("三角函数图:")
	fmt.Println("----------------")

	// 创建折线图
	line := goterm.NewLineChart().
		SetTitle("正弦和余弦函数").
		SetAxisTitles("x", "y").
		SetWidth(80).
		SetHeight(20)

	// 生成正弦曲线数据
	for i := 0; i <= 60; i++ {
		x := float64(i) * 0.1 * math.Pi
		y := math.Sin(x)
		line.AddPoint("sin(x)", x, y)
	}

	// 生成余弦曲线数据
	for i := 0; i <= 60; i++ {
		x := float64(i) * 0.1 * math.Pi
		y := math.Cos(x)
		line.AddPoint("cos(x)", x, y)
	}

	// 自定义样式
	line.SetTitleStyle(goterm.New().Bold().Magenta())
	line.SetLineStyle("sin(x)", goterm.New().Red())
	line.SetLineStyle("cos(x)", goterm.New().Cyan())
	line.SetMarkerStyle("sin(x)", goterm.New().Bold().Red())
	line.SetMarkerStyle("cos(x)", goterm.New().Bold().Cyan())

	// 打印折线图
	line.Print()
}
