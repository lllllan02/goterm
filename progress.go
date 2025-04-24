package goterm

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

// ProgressBarType 表示进度条的类型
type ProgressBarType int

const (
	// BarTypePercent 百分比进度条
	BarTypePercent ProgressBarType = iota
	// BarTypeSpinner 旋转指示器
	BarTypeSpinner
)

// ProgressBar 表示进度条
type ProgressBar struct {
	Total       int64           // 总进度
	Current     int64           // 当前进度
	Width       int             // 进度条宽度
	Type        ProgressBarType // 进度条类型
	ShowPercent bool            // 是否显示百分比
	ShowValue   bool            // 是否显示值
	Fill        string          // 填充字符
	Empty       string          // 空字符
	Spinner     []string        // 旋转指示器字符集
	Prefix      string          // 前缀
	Suffix      string          // 后缀
	Style       *Style          // 样式
	mutex       sync.Mutex      // 互斥锁
	finished    bool            // 是否已完成
	spinnerIdx  int             // 当前旋转指示器索引
	lastPrint   time.Time       // 上次打印时间
}

// NewProgressBar 创建一个新的进度条
func NewProgressBar(total int64) *ProgressBar {
	return &ProgressBar{
		Total:       total,
		Current:     0,
		Width:       50,
		Type:        BarTypePercent,
		ShowPercent: true,
		ShowValue:   true,
		Fill:        "█",
		Empty:       "░",
		Spinner:     []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"},
		Prefix:      "",
		Suffix:      "",
		Style:       New(),
		mutex:       sync.Mutex{},
		finished:    false,
		spinnerIdx:  0,
		lastPrint:   time.Now(),
	}
}

// NewSpinner 创建一个新的旋转指示器
func NewSpinner() *ProgressBar {
	spinner := NewProgressBar(100)
	spinner.Type = BarTypeSpinner
	spinner.ShowPercent = false
	spinner.ShowValue = false
	return spinner
}

// SetFill 设置填充字符
func (p *ProgressBar) SetFill(fill string) *ProgressBar {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.Fill = fill
	return p
}

// SetEmpty 设置空字符
func (p *ProgressBar) SetEmpty(empty string) *ProgressBar {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.Empty = empty
	return p
}

// SetWidth 设置进度条宽度
func (p *ProgressBar) SetWidth(width int) *ProgressBar {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.Width = width
	return p
}

// SetSpinner 设置旋转指示器字符集
func (p *ProgressBar) SetSpinner(spinner []string) *ProgressBar {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.Spinner = spinner
	return p
}

// SetPrefix 设置前缀
func (p *ProgressBar) SetPrefix(prefix string) *ProgressBar {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.Prefix = prefix
	return p
}

// SetSuffix 设置后缀
func (p *ProgressBar) SetSuffix(suffix string) *ProgressBar {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.Suffix = suffix
	return p
}

// SetStyle 设置样式
func (p *ProgressBar) SetStyle(style *Style) *ProgressBar {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.Style = style
	return p
}

// SetShowPercent 设置是否显示百分比
func (p *ProgressBar) SetShowPercent(show bool) *ProgressBar {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.ShowPercent = show
	return p
}

// SetShowValue 设置是否显示值
func (p *ProgressBar) SetShowValue(show bool) *ProgressBar {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.ShowValue = show
	return p
}

// Increment 增加进度
func (p *ProgressBar) Increment() {
	p.Add(1)
}

// Add 增加指定数量的进度
func (p *ProgressBar) Add(n int64) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.Current += n
	if p.Current > p.Total {
		p.Current = p.Total
	}

	// 自动打印进度
	p.print(false)
}

// Set 设置当前进度
func (p *ProgressBar) Set(current int64) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.Current = current
	if p.Current > p.Total {
		p.Current = p.Total
	}

	// 自动打印进度
	p.print(false)
}

// Print 打印进度条
func (p *ProgressBar) Print() {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.print(false)
}

// Finish 完成进度条
func (p *ProgressBar) Finish() {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.Current = p.Total
	p.finished = true
	p.print(true)
	fmt.Println()
}

// Start 启动一个旋转指示器并返回停止函数
func (p *ProgressBar) Start() func() {
	if p.Type != BarTypeSpinner {
		// 只对旋转指示器有效
		return func() {}
	}

	stop := make(chan struct{})

	go func() {
		for {
			select {
			case <-stop:
				return
			default:
				p.mutex.Lock()
				p.spinnerIdx = (p.spinnerIdx + 1) % len(p.Spinner)
				p.print(false)
				p.mutex.Unlock()
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	return func() {
		close(stop)
		// 清理行
		fmt.Print("\r\033[K")
	}
}

// SetPercent 直接设置百分比进度（0-100）
func (p *ProgressBar) SetPercent(percent float64) {
	if percent < 0 {
		percent = 0
	}
	if percent > 100 {
		percent = 100
	}

	current := int64(percent * float64(p.Total) / 100)
	p.Set(current)
}

// print 私有方法，实际打印进度条
func (p *ProgressBar) print(force bool) {
	// 如果上次打印时间间隔小于100毫秒且不是强制打印，则跳过
	if !force && time.Since(p.lastPrint) < 100*time.Millisecond {
		return
	}

	p.lastPrint = time.Now()

	// 清除当前行
	fmt.Print("\r\033[K")

	// 根据类型打印不同的进度条
	switch p.Type {
	case BarTypePercent:
		p.printPercentBar()
	case BarTypeSpinner:
		p.printSpinner()
	}
}

// printPercentBar 打印百分比进度条
func (p *ProgressBar) printPercentBar() {
	percent := float64(p.Current) / float64(p.Total) * 100
	width := p.Width

	// 计算已完成的进度条长度
	completedWidth := int(float64(width) * float64(p.Current) / float64(p.Total))

	// 构建进度条字符串
	var bar strings.Builder

	// 添加前缀
	if p.Prefix != "" {
		bar.WriteString(p.Prefix + " ")
	}

	// 添加进度条
	bar.WriteString("[")
	if completedWidth > 0 {
		bar.WriteString(strings.Repeat(p.Fill, completedWidth))
	}
	if width-completedWidth > 0 {
		bar.WriteString(strings.Repeat(p.Empty, width-completedWidth))
	}
	bar.WriteString("]")

	// 添加百分比
	if p.ShowPercent {
		bar.WriteString(fmt.Sprintf(" %.1f%%", percent))
	}

	// 添加值
	if p.ShowValue {
		bar.WriteString(fmt.Sprintf(" %d/%d", p.Current, p.Total))
	}

	// 添加后缀
	if p.Suffix != "" {
		bar.WriteString(" " + p.Suffix)
	}

	// 打印进度条
	if p.Style != nil {
		fmt.Print(p.Style.Sprint(bar.String()))
	} else {
		fmt.Print(bar.String())
	}
}

// printSpinner 打印旋转指示器
func (p *ProgressBar) printSpinner() {
	var spinner strings.Builder

	// 添加前缀
	if p.Prefix != "" {
		spinner.WriteString(p.Prefix + " ")
	}

	// 添加旋转指示器
	if len(p.Spinner) > 0 {
		spinner.WriteString(p.Spinner[p.spinnerIdx])
	} else {
		spinner.WriteString("-")
	}

	// 添加后缀
	if p.Suffix != "" {
		spinner.WriteString(" " + p.Suffix)
	}

	// 打印旋转指示器
	if p.Style != nil {
		fmt.Print(p.Style.Sprint(spinner.String()))
	} else {
		fmt.Print(spinner.String())
	}
}
