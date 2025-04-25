package main

import (
	"fmt"
	"time"

	"github.com/lllllan02/goterm"
)

func main() {
	// 创建动画控制器
	animation := goterm.NewAnimation()

	// 示例1: 普通打字效果
	fmt.Println("示例1: 普通打字效果")
	typewriter := animation.NewTypewriter("这是一个打字效果的演示，文字会一个字符一个字符地显示出来，就像有人在实时打字一样。")
	typewriter.SetDelay(80 * time.Millisecond) // 调整打字速度
	typewriter.Play()
	time.Sleep(500 * time.Millisecond)

	// 示例2: 渐入打字效果
	fmt.Println("\n示例2: 渐入打字效果")
	fadeTypewriter := animation.NewTypewriter("这是渐入效果，整个文本会逐渐出现。")
	fadeTypewriter.SetDelay(100 * time.Millisecond)
	fadeTypewriter.SetMode(goterm.TypewriterModeFadeIn)
	fadeTypewriter.Play()
	time.Sleep(500 * time.Millisecond)

	// 示例3: 闪烁打字效果
	fmt.Println("\n示例3: 闪烁打字效果")
	blinkTypewriter := animation.NewTypewriter("这是闪烁效果，文字打印时会有闪烁光标。")
	blinkTypewriter.SetDelay(100 * time.Millisecond)
	blinkTypewriter.SetMode(goterm.TypewriterModeBlinking)
	blinkTypewriter.Play()
	time.Sleep(500 * time.Millisecond)

	// 示例4: 彩虹文字效果
	fmt.Println("\n示例4: 彩虹文字效果")
	animation.RainbowText("这是彩虹效果文本，颜色会不断变化！", 3, 100*time.Millisecond)

	fmt.Println("\n所有动画效果演示完成！")
}
