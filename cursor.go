package goterm

import (
	"fmt"
)

// Cursor 提供终端光标控制功能
type Cursor struct{}

// NewCursor 创建一个新的Cursor实例
func NewCursor() *Cursor {
	return &Cursor{}
}

// MoveUp 将光标向上移动n行
func (c *Cursor) MoveUp(n int) {
	fmt.Printf("\033[%dA", n)
}

// MoveDown 将光标向下移动n行
func (c *Cursor) MoveDown(n int) {
	fmt.Printf("\033[%dB", n)
}

// MoveRight 将光标向右移动n列
func (c *Cursor) MoveRight(n int) {
	fmt.Printf("\033[%dC", n)
}

// MoveLeft 将光标向左移动n列
func (c *Cursor) MoveLeft(n int) {
	fmt.Printf("\033[%dD", n)
}

// MoveTo 将光标移动到指定位置
func (c *Cursor) MoveTo(row, col int) {
	fmt.Printf("\033[%d;%dH", row, col)
}

// SavePosition 保存当前光标位置
func (c *Cursor) SavePosition() {
	fmt.Print("\033[s")
}

// RestorePosition 恢复之前保存的光标位置
func (c *Cursor) RestorePosition() {
	fmt.Print("\033[u")
}

// ClearScreen 清除整个屏幕
func (c *Cursor) ClearScreen() {
	fmt.Print("\033[2J")
}

// ClearLine 清除从光标位置到行尾的内容
func (c *Cursor) ClearLine() {
	fmt.Print("\033[K")
}

// HideCursor 隐藏光标
func (c *Cursor) HideCursor() {
	fmt.Print("\033[?25l")
}

// ShowCursor 显示光标
func (c *Cursor) ShowCursor() {
	fmt.Print("\033[?25h")
}
