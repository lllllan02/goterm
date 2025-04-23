package main

import (
	"fmt"
	"time"

	"github.com/lllllan02/goterm"
)

func main() {
	cursor := goterm.NewCursor()

	// 清除屏幕
	cursor.ClearScreen()

	// 移动到屏幕中央
	cursor.MoveTo(10, 20)
	fmt.Print("Hello, World!")

	// 等待1秒
	time.Sleep(time.Second)

	// 清除当前行
	cursor.ClearLine()

	// 保存当前位置
	cursor.SavePosition()
	fmt.Print("First line")

	// 移动到下一行
	cursor.MoveDown(1)
	fmt.Print("Second line")

	// 恢复之前保存的位置
	cursor.RestorePosition()
	fmt.Print(" (overwritten)")

	// 隐藏光标
	cursor.HideCursor()
	time.Sleep(time.Second)

	// 显示光标
	cursor.ShowCursor()

	// 演示光标移动
	cursor.MoveTo(15, 0)
	fmt.Print("Moving cursor: ")
	for i := 0; i < 5; i++ {
		cursor.MoveRight(1)
		fmt.Print(".")
		time.Sleep(200 * time.Millisecond)
	}
	fmt.Println()
}
