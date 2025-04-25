package goterm

import (
	"fmt"
	"time"
)

// Animation 提供终端动画效果
type Animation struct {
	cursor *Cursor
}

// NewAnimation 创建一个新的Animation实例
func NewAnimation() *Animation {
	return &Animation{
		cursor: NewCursor(),
	}
}

// 增强型动画样式常量
const (
	// 打字效果
	TypewriterModeNormal   = 0 // 普通打字效果
	TypewriterModeFadeIn   = 1 // 渐入效果
	TypewriterModeBlinking = 2 // 闪烁效果
)

// Typewriter 打字效果
type Typewriter struct {
	text      string
	delay     time.Duration
	cursor    *Cursor
	isRunning bool
	mode      int
}

// NewTypewriter 创建一个新的打字效果
func (a *Animation) NewTypewriter(text string) *Typewriter {
	return &Typewriter{
		text:      text,
		delay:     50 * time.Millisecond,
		cursor:    a.cursor,
		isRunning: false,
		mode:      TypewriterModeNormal,
	}
}

// SetDelay 设置字符间的延迟时间
func (t *Typewriter) SetDelay(delay time.Duration) *Typewriter {
	t.delay = delay
	return t
}

// SetMode 设置打字效果模式
func (t *Typewriter) SetMode(mode int) *Typewriter {
	t.mode = mode
	return t
}

// Play 播放打字效果
func (t *Typewriter) Play() {
	if t.isRunning {
		return
	}

	t.isRunning = true
	t.cursor.HideCursor()

	switch t.mode {
	case TypewriterModeNormal:
		t.playNormal()
	case TypewriterModeFadeIn:
		t.playFadeIn()
	case TypewriterModeBlinking:
		t.playBlinking()
	default:
		t.playNormal()
	}

	t.isRunning = false
	t.cursor.ShowCursor()
	fmt.Println() // 打字效果结束后换行
}

// playNormal 普通打字效果
func (t *Typewriter) playNormal() {
	runes := []rune(t.text)
	for i := 0; i < len(runes); i++ {
		fmt.Print(string(runes[i]))
		time.Sleep(t.delay)
	}
}

// playFadeIn 渐入效果
func (t *Typewriter) playFadeIn() {
	runes := []rune(t.text)
	for i := 0; i < len(runes); i++ {
		// 显示到当前位置的文本
		t.cursor.ClearLine()
		fmt.Print("\r")
		fmt.Print(string(runes[:i+1]))
		time.Sleep(t.delay)
	}
}

// playBlinking 闪烁效果
func (t *Typewriter) playBlinking() {
	runes := []rune(t.text)
	for i := 0; i < len(runes); i++ {
		// 先显示字符
		fmt.Print(string(runes[i]))

		// 如果不是最后一个字符，添加闪烁光标
		if i < len(runes)-1 {
			fmt.Print("▋")
			time.Sleep(t.delay / 2)
			fmt.Print("\b \b") // 删除光标
			time.Sleep(t.delay / 2)
		}
	}
}

// RainbowText 彩虹文字动画
func (a *Animation) RainbowText(text string, cycles int, delay time.Duration) {
	colors := []string{
		"\033[31m", // 红
		"\033[33m", // 黄
		"\033[32m", // 绿
		"\033[36m", // 青
		"\033[34m", // 蓝
		"\033[35m", // 紫
	}

	reset := "\033[0m"
	runes := []rune(text)

	a.cursor.HideCursor()

	for cycle := 0; cycle < cycles; cycle++ {
		for colorIndex := 0; colorIndex < len(colors); colorIndex++ {
			fmt.Print("\r")
			for i := 0; i < len(runes); i++ {
				currentColor := colors[(colorIndex+i)%len(colors)]
				fmt.Print(currentColor + string(runes[i]) + reset)
			}
			time.Sleep(delay)
		}
	}

	a.cursor.ShowCursor()
	fmt.Println()
}
