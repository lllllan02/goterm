package goterm

// 预设样式类型
var (
	// 常用颜色快捷方式
	StyleRed     = New().Red()     // 红色
	StyleGreen   = New().Green()   // 绿色
	StyleYellow  = New().Yellow()  // 黄色
	StyleBlue    = New().Blue()    // 蓝色
	StyleMagenta = New().Magenta() // 洋红色
	StyleCyan    = New().Cyan()    // 青色
	StyleWhite   = New().White()   // 白色
	StyleBlack   = New().Black()   // 黑色

	// 预设样式
	StyleError   = New().Bold().Red()    // 错误
	StyleSuccess = New().Bold().Green()  // 成功
	StyleWarning = New().Bold().Yellow() // 警告
	StyleInfo    = New().Bold().Blue()   // 信息
	StyleRemark  = New().Bold().Cyan()   // 备注
)

// 全局快捷函数 - 基础颜色
func Red(a ...any) string     { return StyleRed.Sprint(a...) }     // 红色
func Green(a ...any) string   { return StyleGreen.Sprint(a...) }   // 绿色
func Yellow(a ...any) string  { return StyleYellow.Sprint(a...) }  // 黄色
func Blue(a ...any) string    { return StyleBlue.Sprint(a...) }    // 蓝色
func Magenta(a ...any) string { return StyleMagenta.Sprint(a...) } // 洋红色
func Cyan(a ...any) string    { return StyleCyan.Sprint(a...) }    // 青色

// 全局快捷函数 - 格式化输出
func Redf(format string, a ...any) string     { return StyleRed.Sprintf(format, a...) }     // 红色格式化
func Greenf(format string, a ...any) string   { return StyleGreen.Sprintf(format, a...) }   // 绿色格式化
func Yellowf(format string, a ...any) string  { return StyleYellow.Sprintf(format, a...) }  // 黄色格式化
func Bluef(format string, a ...any) string    { return StyleBlue.Sprintf(format, a...) }    // 蓝色格式化
func Magentaf(format string, a ...any) string { return StyleMagenta.Sprintf(format, a...) } // 洋红色格式化
func Cyanf(format string, a ...any) string    { return StyleCyan.Sprintf(format, a...) }    // 青色格式化

// 全局快捷函数 - 预设样式
func Error(a ...any) string   { return StyleError.Sprint(a...) }   // 错误
func Success(a ...any) string { return StyleSuccess.Sprint(a...) } // 成功
func Warning(a ...any) string { return StyleWarning.Sprint(a...) } // 警告
func Info(a ...any) string    { return StyleInfo.Sprint(a...) }    // 信息
func Remark(a ...any) string  { return StyleRemark.Sprint(a...) }  // 备注

// 全局快捷函数 - 格式化输出
func Errorf(format string, a ...any) string   { return StyleError.Sprintf(format, a...) }   // 错误格式化
func Successf(format string, a ...any) string { return StyleSuccess.Sprintf(format, a...) } // 成功格式化
func Warningf(format string, a ...any) string { return StyleWarning.Sprintf(format, a...) } // 警告格式化
func Infof(format string, a ...any) string    { return StyleInfo.Sprintf(format, a...) }    // 信息格式化
func Remarkf(format string, a ...any) string  { return StyleRemark.Sprintf(format, a...) }  // 备注格式化
