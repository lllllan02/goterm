package main

import (
	"fmt"

	"github.com/lllllan02/goterm"
)

func main() {
	// 创建交互式组件
	interactive := goterm.NewInteractive()

	fmt.Println("==== 交互式组件演示 ====")
	fmt.Println()

	// 输入框示例
	fmt.Println("1. 输入框示例:")
	name := interactive.NewInputField("请输入您的名字").ReadString()
	age := interactive.NewInputField("请输入您的年龄").WithDefault("18").ReadString()
	fmt.Printf("您好，%s！您的年龄是 %s。\n\n", name, age)

	// 下拉菜单示例
	fmt.Println("2. 下拉菜单示例:")
	options := []string{
		"选项 A",
		"选项 B",
		"选项 C",
		"选项 D",
	}
	selected := interactive.NewDropdownMenu("请选择一个选项", options).Show()
	fmt.Printf("您选择了: %s (索引值: %d)\n\n", options[selected], selected)

	// 选择框示例
	fmt.Println("3. 选择框示例:")
	colorOptions := []goterm.SelectOption{
		{Value: "red", Label: "红色"},
		{Value: "green", Label: "绿色"},
		{Value: "blue", Label: "蓝色"},
		{Value: "yellow", Label: "黄色"},
	}
	colorResult := interactive.NewSelectField("请选择一个颜色", colorOptions).Render()
	fmt.Printf("您选择了颜色: %s (值: %s)\n\n", colorResult.Label, colorResult.Value)

	// 结束
	fmt.Println("交互式组件演示结束！")
}
