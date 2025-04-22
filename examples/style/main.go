// Package main 展示 GoTerm 库样式功能示例
package main

import (
	"fmt"

	"github.com/lllllan02/goterm"
)

func main() {
	fmt.Println("=== GoTerm 样式示例 ===")

	// 预设函数
	fmt.Println("\n预设函数:")
	fmt.Println(goterm.Red("红色文本"))
	fmt.Println(goterm.Green("绿色文本"))
	fmt.Println(goterm.Yellow("黄色文本"))
	fmt.Println(goterm.Blue("蓝色文本"))
	fmt.Println(goterm.Magenta("洋红色文本"))
	fmt.Println(goterm.Cyan("青色文本"))
	fmt.Println(goterm.Error("错误文本"))
	fmt.Println(goterm.Success("成功文本"))
	fmt.Println(goterm.Warning("警告文本"))
	fmt.Println(goterm.Info("信息文本"))
	fmt.Println(goterm.Remark("备注文本"))

	// 预设样式
	fmt.Println("\n预设样式:")
	goterm.StyleRed.Println("红色文本")
	goterm.StyleGreen.Println("绿色文本")
	goterm.StyleYellow.Println("黄色文本")
	goterm.StyleBlue.Println("蓝色文本")
	goterm.StyleMagenta.Println("洋红色文本")
	goterm.StyleCyan.Println("青色文本")

	// 基本样式组合
	fmt.Println("\n基本样式:")
	goterm.New().Bold().Red().Println("粗体红色文本")
	goterm.New().Italic().Green().Println("斜体绿色文本")
	goterm.New().Underline().Blue().Println("下划线蓝色文本")
	goterm.New().White().BgBlue().Println("白色文本蓝色背景")

	// 自定义样式
	fmt.Println("\n自定义样式:")
	errorStyle := goterm.New().Bold().Red()
	warnStyle := goterm.New().Yellow()

	errorStyle.Println("自定义错误样式")
	warnStyle.Println("自定义警告样式")
}
