package main

import (
	"fmt"

	"github.com/lllllan02/goterm"
)

func main() {
	// 输出不同级别的日志
	fmt.Println(goterm.Error("这是一条错误信息"))
	fmt.Println(goterm.Success("这是一条成功信息"))
	fmt.Println(goterm.Warning("这是一条警告信息"))
	fmt.Println(goterm.Info("这是一条普通信息"))
	fmt.Println(goterm.Remark("这是一条备注信息"))

	// 使用格式化版本
	fmt.Println(goterm.Errorf("格式化错误信息: %d", 404))
	fmt.Println(goterm.Successf("格式化成功信息: %s", "操作完成"))
	fmt.Println(goterm.Warningf("格式化警告信息: %v", "即将超时"))
	fmt.Println(goterm.Infof("格式化普通信息: %f", 3.14))
	fmt.Println(goterm.Remarkf("格式化备注信息: %t", true))
}
