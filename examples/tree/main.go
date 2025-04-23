package main

import (
	"fmt"

	"github.com/lllllan02/goterm"
)

func main() {
	fmt.Println("目录结构示例：")
	// 创建目录结构
	dirTree := goterm.NewTree("project", nil)
	src := dirTree.Root.AddChild("src", nil)
	src.AddChild("main.go", nil)
	utils := src.AddChild("utils", nil)
	utils.AddChild("logger.go", nil)
	utils.AddChild("style.go", nil)
	examples := src.AddChild("examples", nil)
	examples.AddChild("main.go", nil)
	dirTree.Root.AddChild("README.md", nil)
	dirTree.Root.AddChild("go.mod", nil)

	// 使用绿色打印目录结构
	dirTree.PrintWithStyle(goterm.New().Bold().Green())

	fmt.Println("\n系统状态示例：")
	// 创建系统状态
	sysTree := goterm.NewTree("系统状态", nil)
	cpu := sysTree.Root.AddChild("CPU使用率", nil)
	cpu.AddChild("用户态", "45%")
	cpu.AddChild("系统态", "15%")
	cpu.AddChild("空闲", "40%")
	mem := sysTree.Root.AddChild("内存使用", nil)
	mem.AddChild("已用", "4.2GB")
	mem.AddChild("可用", "2.8GB")
	disk := sysTree.Root.AddChild("磁盘空间", nil)
	disk.AddChild("已用", "120GB")
	disk.AddChild("可用", "80GB")

	// 使用蓝色打印系统状态
	sysTree.PrintWithStyle(goterm.New().Bold().Blue())

	fmt.Println("\n自定义样式示例：")
	// 创建自定义树
	customTree := goterm.NewTree("自定义树", nil)
	level1 := customTree.Root.AddChild("第一层", nil)
	level2 := level1.AddChild("第二层", "值1")
	level2.AddChild("第三层", "值2")
	level2.AddChild("第三层", "值3")
	level1.AddChild("第二层", "值4")

	// 使用不同的颜色打印不同层级
	customTree.PrintWithStyle(goterm.New().Bold().RGB(128, 0, 128)) // 紫色
}
