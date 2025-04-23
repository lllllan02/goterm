package goterm

import (
	"fmt"
	"strings"
)

// TextFormatter 提供文本格式化功能
type TextFormatter struct {
	style *Style
}

// NewTextFormatter 创建一个新的文本格式化器
func NewTextFormatter() *TextFormatter {
	return &TextFormatter{
		style: New(),
	}
}

// Paragraph 格式化段落
// text: 段落文本
// indent: 缩进空格数
func (tf *TextFormatter) Paragraph(text string, indent int) string {
	// 添加缩进
	indentStr := strings.Repeat(" ", indent)

	// 处理换行，确保每行都有正确的缩进
	lines := strings.Split(text, "\n")
	for i, line := range lines {
		if line != "" {
			lines[i] = indentStr + line
		}
	}

	return strings.Join(lines, "\n")
}

// List 格式化列表
// items: 列表项
// bullet: 列表符号（默认为"•"）
// indent: 缩进空格数
func (tf *TextFormatter) List(items []string, bullet string, indent int) string {
	if bullet == "" {
		bullet = "•"
	}

	// 添加缩进
	indentStr := strings.Repeat(" ", indent)

	// 格式化每个列表项
	formattedItems := make([]string, len(items))
	for i, item := range items {
		formattedItems[i] = fmt.Sprintf("%s%s %s", indentStr, bullet, item)
	}

	return strings.Join(formattedItems, "\n")
}

// NumberedList 格式化有序列表
// items: 列表项
// start: 起始编号
// indent: 缩进空格数
func (tf *TextFormatter) NumberedList(items []string, start int, indent int) string {
	// 添加缩进
	indentStr := strings.Repeat(" ", indent)

	// 格式化每个列表项
	formattedItems := make([]string, len(items))
	for i, item := range items {
		formattedItems[i] = fmt.Sprintf("%s%d. %s", indentStr, start+i, item)
	}

	return strings.Join(formattedItems, "\n")
}
