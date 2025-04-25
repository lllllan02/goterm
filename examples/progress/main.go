package main

import (
	"fmt"
	"time"

	"github.com/lllllan02/goterm"
)

func main() {
	// 演示基本百分比进度条
	fmt.Println("基本百分比进度条：")
	showBasicProgressBar()
	fmt.Println()

	// 演示自定义样式的进度条
	fmt.Println("自定义样式的进度条：")
	showStyledProgressBar()
	fmt.Println()

	// 演示旋转指示器
	fmt.Println("旋转指示器：")
	showSpinner()
	fmt.Println()

	// 演示不同填充样式的进度条
	fmt.Println("不同填充样式的进度条：")
	showCustomFillProgressBar()
	fmt.Println()

	// 演示固定在底部的进度条
	fmt.Println("固定在底部的进度条（带日志信息）：")
	fmt.Println("按Ctrl+C退出演示")
	showStickyProgressBar()

	// 演示使用现有logger的固定进度条
	fmt.Println("\n使用现有logger的固定进度条：")
	fmt.Println("按Ctrl+C退出演示")
	showStickyProgressBarWithExistingLogger()

	// 演示更简单的方式使用logger与进度条结合
	fmt.Println("\n使用更简单方式的日志与进度条结合：")
	fmt.Println("按Ctrl+C退出演示")
	showSimplifiedLoggerWithProgressBar()
}

// 基本百分比进度条
func showBasicProgressBar() {
	// 创建一个总量为100的进度条
	bar := goterm.NewProgressBar(100)
	bar.SetPrefix("下载中")

	// 模拟进度
	for i := 0; i <= 100; i += 5 {
		bar.Set(int64(i))
		time.Sleep(100 * time.Millisecond)
	}

	// 完成进度条
	bar.Finish()
}

// 自定义样式的进度条
func showStyledProgressBar() {
	// 创建一个总量为150的进度条
	bar := goterm.NewProgressBar(150)
	bar.SetPrefix("处理文件")
	bar.SetSuffix("请稍候...")
	bar.SetWidth(30)
	bar.SetStyle(goterm.New().Green())

	// 模拟进度
	for i := 0; i <= 150; i += 10 {
		bar.Set(int64(i))
		time.Sleep(100 * time.Millisecond)
	}

	// 完成进度条
	bar.Finish()
}

// 旋转指示器
func showSpinner() {
	// 创建一个旋转指示器
	spinner := goterm.NewSpinner()
	spinner.SetPrefix("加载中")
	spinner.SetStyle(goterm.New().Blue())

	// 启动旋转指示器
	stop := spinner.Start()

	// 模拟处理过程
	time.Sleep(3 * time.Second)

	// 停止旋转指示器
	stop()
	fmt.Println("加载完成！")
}

// 不同填充样式的进度条
func showCustomFillProgressBar() {
	// 使用不同填充字符的进度条
	fills := []struct {
		name  string
		fill  string
		empty string
	}{
		{"方块", "█", "░"},
		{"散列", "#", "-"},
		{"星号", "*", " "},
		{"箭头", "▶", "◀"},
	}

	for _, f := range fills {
		// 创建进度条
		bar := goterm.NewProgressBar(100)
		bar.SetFill(f.fill)
		bar.SetEmpty(f.empty)
		bar.SetPrefix(f.name)
		bar.SetWidth(30)
		bar.SetShowValue(false)

		// 模拟进度
		for i := 0; i <= 100; i += 20 {
			bar.Set(int64(i))
			time.Sleep(100 * time.Millisecond)
		}

		// 完成进度条
		bar.Finish()
	}
}

// 演示固定在底部的进度条
func showStickyProgressBar() {
	// 创建一个固定在底部的进度条
	bar := goterm.NewStickyProgressBar(100)
	bar.SetPrefix("数据处理")
	bar.SetStyle(goterm.New().Yellow())
	bar.SetWidth(40)

	// 设置最大显示日志行数
	bar.SetMaxLogLines(20)

	// 启动进度条
	stop := bar.Start()
	defer stop()

	// 模拟一个长时间任务，同时输出日志
	tasks := []string{
		"初始化系统",
		"连接数据库",
		"加载用户数据",
		"处理文件",
		"验证结果",
		"生成报告",
		"发送通知",
		"清理临时文件",
		"更新缓存",
		"完成任务",
	}

	for i, task := range tasks {
		// 记录开始任务
		bar.Log("[%s] 开始: %s", time.Now().Format("15:04:05"), task)

		// 模拟任务处理过程
		subTasks := 5
		for j := 0; j < subTasks; j++ {
			time.Sleep(300 * time.Millisecond)

			// 记录子任务进度
			bar.Log("  - 子任务 %d/%d: %s 进行中...", j+1, subTasks, task)

			// 更新总进度
			progress := int64(i*10 + (j+1)*10/subTasks)
			bar.Set(progress)
		}

		// 记录任务完成
		bar.Log("[%s] 完成: %s ✓", time.Now().Format("15:04:05"), task)
	}

	// 完成进度条
	bar.Finish()

	fmt.Println("所有任务处理完成！")
}

// 使用现有logger与固定进度条结合的示例
func showStickyProgressBarWithExistingLogger() {
	// 创建一个固定在底部的进度条
	bar := goterm.NewStickyProgressBar(100)
	bar.SetPrefix("处理任务")
	bar.SetStyle(goterm.New().Yellow())
	bar.SetWidth(40)

	// 设置最大显示日志行数
	bar.SetMaxLogLines(20)

	// 获取日志写入器
	logWriter := bar.GetLogWriter()

	// 启动进度条
	stop := bar.Start()
	defer stop()

	// 模拟一个长时间任务，使用现有logger输出到进度条
	tasks := []struct {
		name     string
		subtasks int
		status   string
	}{
		{name: "系统初始化", subtasks: 3, status: "成功"},
		{name: "数据加载", subtasks: 5, status: "成功"},
		{name: "用户认证", subtasks: 2, status: "警告"},
		{name: "资源分配", subtasks: 4, status: "成功"},
		{name: "网络检查", subtasks: 3, status: "错误"},
		{name: "配置更新", subtasks: 4, status: "成功"},
		{name: "缓存清理", subtasks: 2, status: "成功"},
	}

	totalProgress := 0
	progressStep := 100 / (len(tasks) * 5) // 每个子任务5%的进度

	for i, task := range tasks {
		// 记录开始任务
		fmt.Fprintf(logWriter, "%s\n", goterm.Infof("开始任务: %s (#%d/%d)", task.name, i+1, len(tasks)))
		time.Sleep(500 * time.Millisecond)

		// 任务进行中
		for j := 1; j <= task.subtasks; j++ {
			time.Sleep(300 * time.Millisecond)

			totalProgress += progressStep
			bar.Set(int64(totalProgress))

			// 根据任务状态使用不同日志级别
			switch {
			case j == task.subtasks && task.status == "错误":
				fmt.Fprintf(logWriter, "%s\n", goterm.Errorf("子任务 %d/%d 失败: %s", j, task.subtasks, task.name))
			case j == task.subtasks && task.status == "警告":
				fmt.Fprintf(logWriter, "%s\n", goterm.Warningf("子任务 %d/%d 完成但有警告: %s", j, task.subtasks, task.name))
			case j == 1:
				fmt.Fprintf(logWriter, "%s\n", goterm.Remarkf("子任务 %d/%d 开始: %s", j, task.subtasks, task.name))
			default:
				fmt.Fprintf(logWriter, "%s\n", goterm.Infof("子任务 %d/%d 进行中: %s", j, task.subtasks, task.name))
			}
		}

		// 记录任务完成
		time.Sleep(300 * time.Millisecond)
		switch task.status {
		case "成功":
			fmt.Fprintf(logWriter, "%s\n", goterm.Successf("任务完成: %s ✓", task.name))
		case "警告":
			fmt.Fprintf(logWriter, "%s\n", goterm.Warningf("任务完成但有警告: %s ⚠", task.name))
		case "错误":
			fmt.Fprintf(logWriter, "%s\n", goterm.Errorf("任务失败: %s ✗", task.name))
		}
	}

	// 完成进度条
	bar.Finish()

	fmt.Println("所有任务处理完成！")
}

// 使用更简单方式的日志与进度条结合示例
func showSimplifiedLoggerWithProgressBar() {
	// 创建一个固定在底部的进度条
	bar := goterm.NewStickyProgressBar(100)
	bar.SetPrefix("处理任务")
	bar.SetStyle(goterm.New().Yellow())
	bar.SetWidth(40)
	bar.SetMaxLogLines(20)

	// 将进度条设置为活跃进度条
	bar.SetAsActive()

	// 启动进度条
	stop := bar.Start()
	defer func() {
		stop()
		// 清除活跃进度条
		bar.ClearActive()
	}()

	// 模拟一个长时间任务，直接使用goterm的日志函数
	tasks := []struct {
		name     string
		subtasks int
		status   string
	}{
		{name: "系统初始化", subtasks: 3, status: "成功"},
		{name: "数据加载", subtasks: 5, status: "成功"},
		{name: "用户认证", subtasks: 2, status: "警告"},
		{name: "资源分配", subtasks: 4, status: "成功"},
		{name: "网络检查", subtasks: 3, status: "错误"},
		{name: "配置更新", subtasks: 4, status: "成功"},
		{name: "缓存清理", subtasks: 2, status: "成功"},
	}

	totalProgress := 0
	progressStep := 100 / (len(tasks) * 5) // 每个子任务5%的进度

	for i, task := range tasks {
		// 直接使用goterm.Infof记录开始任务
		goterm.Infof("开始任务: %s (#%d/%d)", task.name, i+1, len(tasks))
		time.Sleep(500 * time.Millisecond)

		// 任务进行中
		for j := 1; j <= task.subtasks; j++ {
			time.Sleep(300 * time.Millisecond)

			totalProgress += progressStep
			bar.Set(int64(totalProgress))

			// 根据任务状态使用不同日志级别
			switch {
			case j == task.subtasks && task.status == "错误":
				goterm.Errorf("子任务 %d/%d 失败: %s", j, task.subtasks, task.name)
			case j == task.subtasks && task.status == "警告":
				goterm.Warningf("子任务 %d/%d 完成但有警告: %s", j, task.subtasks, task.name)
			case j == 1:
				goterm.Remarkf("子任务 %d/%d 开始: %s", j, task.subtasks, task.name)
			default:
				goterm.Infof("子任务 %d/%d 进行中: %s", j, task.subtasks, task.name)
			}
		}

		// 记录任务完成
		time.Sleep(300 * time.Millisecond)
		switch task.status {
		case "成功":
			goterm.Successf("任务完成: %s ✓", task.name)
		case "警告":
			goterm.Warningf("任务完成但有警告: %s ⚠", task.name)
		case "错误":
			goterm.Errorf("任务失败: %s ✗", task.name)
		}
	}

	// 完成进度条
	bar.Finish()

	fmt.Println("所有任务处理完成！")
}
