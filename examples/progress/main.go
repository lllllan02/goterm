package main

import (
	"fmt"
	"time"

	"github.com/lllllan02/goterm"
)

func main() {
	// 演示基本百分比进度条
	fmt.Println("基本百分比进度条：")
	showBasicProgressBar()
	fmt.Println()

	// 演示自定义样式的进度条
	fmt.Println("自定义样式的进度条：")
	showStyledProgressBar()
	fmt.Println()

	// 演示旋转指示器
	fmt.Println("旋转指示器：")
	showSpinner()
	fmt.Println()

	// 演示不同填充样式的进度条
	fmt.Println("不同填充样式的进度条：")
	showCustomFillProgressBar()
	fmt.Println()
}

// 基本百分比进度条
func showBasicProgressBar() {
	// 创建一个总量为100的进度条
	bar := goterm.NewProgressBar(100)
	bar.SetPrefix("下载中")

	// 模拟进度
	for i := 0; i <= 100; i += 5 {
		bar.Set(int64(i))
		time.Sleep(100 * time.Millisecond)
	}

	// 完成进度条
	bar.Finish()
}

// 自定义样式的进度条
func showStyledProgressBar() {
	// 创建一个总量为150的进度条
	bar := goterm.NewProgressBar(150)
	bar.SetPrefix("处理文件")
	bar.SetSuffix("请稍候...")
	bar.SetWidth(30)
	bar.SetStyle(goterm.New().Green())

	// 模拟进度
	for i := 0; i <= 150; i += 10 {
		bar.Set(int64(i))
		time.Sleep(100 * time.Millisecond)
	}

	// 完成进度条
	bar.Finish()
}

// 旋转指示器
func showSpinner() {
	// 创建一个旋转指示器
	spinner := goterm.NewSpinner()
	spinner.SetPrefix("加载中")
	spinner.SetStyle(goterm.New().Blue())

	// 启动旋转指示器
	stop := spinner.Start()

	// 模拟处理过程
	time.Sleep(3 * time.Second)

	// 停止旋转指示器
	stop()
	fmt.Println("加载完成！")
}

// 不同填充样式的进度条
func showCustomFillProgressBar() {
	// 使用不同填充字符的进度条
	fills := []struct {
		name  string
		fill  string
		empty string
	}{
		{"方块", "█", "░"},
		{"散列", "#", "-"},
		{"星号", "*", " "},
		{"箭头", "▶", "◀"},
	}

	for _, f := range fills {
		// 创建进度条
		bar := goterm.NewProgressBar(100)
		bar.SetFill(f.fill)
		bar.SetEmpty(f.empty)
		bar.SetPrefix(f.name)
		bar.SetWidth(30)
		bar.SetShowValue(false)

		// 模拟进度
		for i := 0; i <= 100; i += 20 {
			bar.Set(int64(i))
			time.Sleep(100 * time.Millisecond)
		}

		// 完成进度条
		bar.Finish()
	}
}
