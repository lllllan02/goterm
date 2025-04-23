package goterm

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/mattn/go-isatty"
)

// ANSI 颜色代码
const (
	Reset     = "\033[0m" // 重置所有样式
	Bold      = "\033[1m" // 粗体
	Faint     = "\033[2m" // 弱化
	Italic    = "\033[3m" // 斜体
	Underline = "\033[4m" // 下划线

	// 前景色（文本颜色）
	FgBlack   = "\033[30m" // 黑色
	FgRed     = "\033[31m" // 红色
	FgGreen   = "\033[32m" // 绿色
	FgYellow  = "\033[33m" // 黄色
	FgBlue    = "\033[34m" // 蓝色
	FgMagenta = "\033[35m" // 洋红色
	FgCyan    = "\033[36m" // 青色
	FgWhite   = "\033[37m" // 白色

	// 背景色
	BgBlack   = "\033[40m" // 黑色
	BgRed     = "\033[41m" // 红色
	BgGreen   = "\033[42m" // 绿色
	BgYellow  = "\033[43m" // 黄色
	BgBlue    = "\033[44m" // 蓝色
	BgMagenta = "\033[45m" // 洋红色
	BgCyan    = "\033[46m" // 青色
	BgWhite   = "\033[47m" // 白色
)

// 全局配置
var (
	// 是否启用彩色输出：
	// 1. 设置 NO_COLOR 环境变量
	// 2. 终端为 dumb 终端
	// 3. 不是终端
	NoColor = os.Getenv("NO_COLOR") != "" || os.Getenv("TERM") == "dumb" ||
		(!isatty.IsTerminal(os.Stdout.Fd()) && !isatty.IsCygwinTerminal(os.Stdout.Fd()))

	// 默认输出目标
	Output io.Writer = os.Stdout
)

// Style 表示一个带有样式的字符串
type Style struct {
	codes []string
}

// New 创建一个新的样式
func New() *Style {
	return &Style{}
}

// add 添加一个样式代码
func (s *Style) add(code string) *Style {
	s.codes = append(s.codes, code)
	return s
}

// 以下是各种样式方法
func (s *Style) Bold() *Style      { return s.add(Bold) }      // 粗体
func (s *Style) Faint() *Style     { return s.add(Faint) }     // 弱化
func (s *Style) Italic() *Style    { return s.add(Italic) }    // 斜体
func (s *Style) Underline() *Style { return s.add(Underline) } // 下划线

// Foreground colors
func (s *Style) Black() *Style   { return s.add(FgBlack) }   // 黑色
func (s *Style) Red() *Style     { return s.add(FgRed) }     // 红色
func (s *Style) Green() *Style   { return s.add(FgGreen) }   // 绿色
func (s *Style) Yellow() *Style  { return s.add(FgYellow) }  // 黄色
func (s *Style) Blue() *Style    { return s.add(FgBlue) }    // 蓝色
func (s *Style) Magenta() *Style { return s.add(FgMagenta) } // 洋红色
func (s *Style) Cyan() *Style    { return s.add(FgCyan) }    // 青色
func (s *Style) White() *Style   { return s.add(FgWhite) }   // 白色

// Background colors
func (s *Style) BgBlack() *Style   { return s.add(BgBlack) }   // 黑色
func (s *Style) BgRed() *Style     { return s.add(BgRed) }     // 红色
func (s *Style) BgGreen() *Style   { return s.add(BgGreen) }   // 绿色
func (s *Style) BgYellow() *Style  { return s.add(BgYellow) }  // 黄色
func (s *Style) BgBlue() *Style    { return s.add(BgBlue) }    // 蓝色
func (s *Style) BgMagenta() *Style { return s.add(BgMagenta) } // 洋红色
func (s *Style) BgCyan() *Style    { return s.add(BgCyan) }    // 青色
func (s *Style) BgWhite() *Style   { return s.add(BgWhite) }   // 白色

// RGB 添加自定义RGB颜色
func (s *Style) RGB(r, g, b int) *Style {
	return s.add(RGB(r, g, b))
}

// BgRGB 添加自定义RGB背景色
func (s *Style) BgRGB(r, g, b int) *Style {
	return s.add(BgRGB(r, g, b))
}

// Sprint 返回带有样式的字符串
func (s *Style) Sprint(a ...any) string {
	if NoColor || len(s.codes) == 0 {
		return fmt.Sprint(a...)
	}

	// 应用所有样式代码
	return strings.Join(s.codes, "") + fmt.Sprint(a...) + Reset
}

// Sprintf 返回带有样式和格式的字符串
func (s *Style) Sprintf(format string, a ...any) string {
	if NoColor || len(s.codes) == 0 {
		return fmt.Sprintf(format, a...)
	}

	return strings.Join(s.codes, "") + fmt.Sprintf(format, a...) + Reset
}

// Print 输出带样式的文本
func (s *Style) Print(a ...any) (n int, err error) {
	return fmt.Fprint(Output, s.Sprint(a...))
}

// Println 输出带样式的文本并换行
func (s *Style) Println(a ...any) (n int, err error) {
	return fmt.Fprintln(Output, s.Sprint(a...))
}

// Printf 输出带样式和格式的文本
func (s *Style) Printf(format string, a ...any) (n int, err error) {
	return fmt.Fprint(Output, s.Sprintf(format, a...))
}

// Fprint 输出带样式的文本到指定的writer
func (s *Style) Fprint(w io.Writer, a ...any) (n int, err error) {
	return fmt.Fprint(w, s.Sprint(a...))
}

// Fprintln 输出带样式的文本到指定的 writer 并换行
func (s *Style) Fprintln(w io.Writer, a ...any) (n int, err error) {
	return fmt.Fprintln(w, s.Sprint(a...))
}

// Fprintf 输出带样式和格式的文本到指定的 writer
func (s *Style) Fprintf(w io.Writer, format string, a ...any) (n int, err error) {
	return fmt.Fprint(w, s.Sprintf(format, a...))
}

// RGB 创建自定义RGB颜色代码（仅支持真彩色终端）
func RGB(r, g, b int) string {
	return "\033[38;2;" + strconv.Itoa(r) + ";" + strconv.Itoa(g) + ";" + strconv.Itoa(b) + "m"
}

// BgRGB 创建自定义RGB背景色代码（仅支持真彩色终端）
func BgRGB(r, g, b int) string {
	return "\033[48;2;" + strconv.Itoa(r) + ";" + strconv.Itoa(g) + ";" + strconv.Itoa(b) + "m"
}

// SetWriter 设置全局输出目标
func SetWriter(w io.Writer) {
	Output = w
}
