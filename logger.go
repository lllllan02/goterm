package goterm

import (
	"fmt"
	"time"
)

// 日志级别前缀
var (
	PrefixError   = New().Bold().Red().Sprint("ERROR")
	PrefixSuccess = New().Bold().RGB(0, 128, 0).Sprint("SUCCESS")
	PrefixWarning = New().Bold().Yellow().Sprint("WARN")
	PrefixInfo    = New().Bold().Blue().Sprint("INFO")
	PrefixRemark  = New().Bold().Cyan().Sprint("REMARK")
)

// 获取不带ANSI转义码的字符串长度
func getPlainLength(s string) int {
	length := 0
	inEscape := false
	for _, c := range s {
		if c == '\x1b' {
			inEscape = true
			continue
		}
		if inEscape {
			if c == 'm' {
				inEscape = false
			}
			continue
		}
		length++
	}
	return length
}

// 格式化日志消息
func formatLog(prefix string, message string) string {
	now := time.Now().Format("2006-01-02 15:04:05")
	// 计算前缀的纯文本长度
	prefixLength := getPlainLength(prefix)
	// 计算需要填充的空格数（SUCCESS是最长的，长度为7）
	totalWidth := 9 // 设置总宽度为9，确保有足够的空间
	leftPadding := (totalWidth - prefixLength) / 2
	rightPadding := totalWidth - prefixLength - leftPadding
	// 添加对齐空格
	alignedPrefix := fmt.Sprintf("%*s%s%*s", leftPadding, "", prefix, rightPadding, "")
	return fmt.Sprintf("[%s] %s %s", now, alignedPrefix, message)
}

// 全局快捷函数 - 预设样式
func Error(a ...any) string {
	return formatLog(PrefixError, fmt.Sprint(a...))
}

func Success(a ...any) string {
	return formatLog(PrefixSuccess, fmt.Sprint(a...))
}

func Warning(a ...any) string {
	return formatLog(PrefixWarning, fmt.Sprint(a...))
}

func Info(a ...any) string {
	return formatLog(PrefixInfo, fmt.Sprint(a...))
}

func Remark(a ...any) string {
	return formatLog(PrefixRemark, fmt.Sprint(a...))
}

// 全局快捷函数 - 格式化输出
func Errorf(format string, a ...any) string {
	return formatLog(PrefixError, fmt.Sprintf(format, a...))
}

func Successf(format string, a ...any) string {
	return formatLog(PrefixSuccess, fmt.Sprintf(format, a...))
}

func Warningf(format string, a ...any) string {
	return formatLog(PrefixWarning, fmt.Sprintf(format, a...))
}

func Infof(format string, a ...any) string {
	return formatLog(PrefixInfo, fmt.Sprintf(format, a...))
}

func Remarkf(format string, a ...any) string {
	return formatLog(PrefixRemark, fmt.Sprintf(format, a...))
}
