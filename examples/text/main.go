package main

import (
	"fmt"

	"github.com/lllllan02/goterm"
)

func main() {
	// 创建文本格式化器
	formatter := goterm.NewTextFormatter()

	// 示例2：段落格式化
	fmt.Println("示例2：段落格式化")
	paragraph := `这是一个段落示例。这是一个段落示例。这是一个段落示例。这是一个段落示例。这是一个段落示例。这是一个段落示例。这是一个段落示例。这是一个段落示例。这是一个段落示例。这是一个段落示例。这是一个段落示例。
它包含多行文本。
每行都会自动缩进。`
	fmt.Println(formatter.Paragraph(paragraph, 4))
	fmt.Println()

	// 示例3：无序列表
	fmt.Println("示例3：无序列表")
	items := []string{
		"第一项",
		"第二项",
		"第三项",
	}
	fmt.Println(formatter.List(items, "•", 2))
	fmt.Println()

	// 示例4：有序列表
	fmt.Println("示例4：有序列表")
	fmt.Println(formatter.NumberedList(items, 1, 2))
	fmt.Println()

	// 示例5：组合使用
	fmt.Println("示例5：组合使用")
	fmt.Println(formatter.Paragraph("这是一个示例文档，展示了各种文本格式化功能。", 0))
	fmt.Println()
	fmt.Println(formatter.List([]string{
		"标题格式化",
		"段落格式化",
		"列表格式化",
	}, "→", 4))
}
