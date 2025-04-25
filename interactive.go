package goterm

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Interactive 提供终端交互式组件
type Interactive struct {
	cursor *Cursor
	style  *Style
}

// NewInteractive 创建一个新的交互式组件实例
func NewInteractive() *Interactive {
	return &Interactive{
		cursor: NewCursor(),
		style:  New(),
	}
}

// InputField 输入框
type InputField struct {
	prompt      string
	defaultText string
	interactive *Interactive
}

// NewInputField 创建一个新的输入框
func (i *Interactive) NewInputField(prompt string) *InputField {
	return &InputField{
		prompt:      prompt,
		defaultText: "",
		interactive: i,
	}
}

// WithDefault 设置默认文本
func (input *InputField) WithDefault(defaultText string) *InputField {
	input.defaultText = defaultText
	return input
}

// ReadString 读取用户输入的字符串
func (input *InputField) ReadString() string {
	promptStyle := input.interactive.style.Green().Bold()
	fmt.Print(promptStyle.Sprint(input.prompt + ": "))

	if input.defaultText != "" {
		fmt.Printf("[%s] ", input.defaultText)
	}

	reader := bufio.NewReader(os.Stdin)
	result, _ := reader.ReadString('\n')
	result = strings.TrimSpace(result)

	if result == "" && input.defaultText != "" {
		return input.defaultText
	}

	return result
}

// SelectOption 选择项
type SelectOption struct {
	Value string
	Label string
}

// SelectField 选择框
type SelectField struct {
	prompt      string
	options     []SelectOption
	selected    int
	interactive *Interactive
}

// NewSelectField 创建一个新的选择框
func (i *Interactive) NewSelectField(prompt string, options []SelectOption) *SelectField {
	return &SelectField{
		prompt:      prompt,
		options:     options,
		selected:    0,
		interactive: i,
	}
}

// Render 渲染选择框并获取用户选择
func (selectField *SelectField) Render() SelectOption {
	// 保存初始位置
	promptStyle := selectField.interactive.style.Blue().Bold()
	fmt.Println(promptStyle.Sprint(selectField.prompt))

	// 初始化一次选项显示
	for i, option := range selectField.options {
		prefix := "  "
		if i == selectField.selected {
			prefix = "> "
			fmt.Println(selectField.interactive.style.Green().Sprint(prefix + option.Label))
		} else {
			fmt.Println(prefix + option.Label)
		}
	}

	// 提示信息
	fmt.Println(selectField.interactive.style.Faint().Sprint("(使用↑↓选择，回车确认，ESC取消)"))

	// 将光标移到选项区域的起始位置
	selectField.interactive.cursor.MoveUp(len(selectField.options) + 1) // +1 是为了提示行

	// 设置终端为原始模式
	selectField.interactive.cursor.HideCursor()
	defer selectField.interactive.cursor.ShowCursor()

	// 将终端设置为原始模式，以便可以读取单个字符
	oldState, err := makeRaw(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println("无法设置终端为原始模式:", err)
		return selectField.options[selectField.selected]
	}
	defer restoreTerminal(int(os.Stdin.Fd()), oldState)

	// 主循环：处理用户输入
	for {
		// 清除并重绘所有选项
		for i, option := range selectField.options {
			// 移动到当前选项行
			if i > 0 {
				selectField.interactive.cursor.MoveDown(1)
			}

			// 清除当前行
			fmt.Print("\r")
			selectField.interactive.cursor.ClearLine()

			// 显示选项
			prefix := "  "
			if i == selectField.selected {
				prefix = "> "
				fmt.Print(selectField.interactive.style.Green().Sprint(prefix + option.Label))
			} else {
				fmt.Print(prefix + option.Label)
			}
		}

		// 移动到提示行并更新
		selectField.interactive.cursor.MoveDown(1)
		fmt.Print("\r")
		selectField.interactive.cursor.ClearLine()
		fmt.Print(selectField.interactive.style.Faint().Sprint("(使用↑↓选择，回车确认，ESC取消)"))

		// 移回第一个选项的位置
		selectField.interactive.cursor.MoveUp(len(selectField.options))

		// 读取按键
		key := readKey()

		// 处理按键
		switch key {
		case keyEsc:
			// 选择第一个选项并退出
			selectField.selected = 0

			// 清除选择区域和提示行
			fmt.Print("\r")
			for i := 0; i < len(selectField.options); i++ {
				selectField.interactive.cursor.ClearLine()
				selectField.interactive.cursor.MoveDown(1)
			}
			selectField.interactive.cursor.ClearLine() // 清除提示行

			// 返回到选项区域开始处并显示取消信息
			selectField.interactive.cursor.MoveUp(len(selectField.options) + 1)
			fmt.Printf("%s: %s\n", selectField.prompt,
				selectField.interactive.style.Red().Sprint("已取消选择"))

			// 如果没有选项，创建一个默认选项
			if len(selectField.options) > 0 {
				return selectField.options[0]
			}
			return SelectOption{Value: "", Label: ""}

		case keyArrowUp:
			if selectField.selected > 0 {
				selectField.selected--
			}
		case keyArrowDown:
			if selectField.selected < len(selectField.options)-1 {
				selectField.selected++
			}
		case keyEnter:
			// 清除选择区域和提示行
			fmt.Print("\r")
			for i := 0; i < len(selectField.options); i++ {
				selectField.interactive.cursor.ClearLine()
				if i < len(selectField.options)-1 {
					selectField.interactive.cursor.MoveDown(1)
				}
			}
			selectField.interactive.cursor.MoveDown(1)
			selectField.interactive.cursor.ClearLine() // 清除提示行

			// 返回到选项区域开始处并显示结果
			selectField.interactive.cursor.MoveUp(len(selectField.options))
			fmt.Printf("%s: %s\n", selectField.prompt,
				selectField.interactive.style.Green().Sprint(
					selectField.options[selectField.selected].Label))

			return selectField.options[selectField.selected]
		}
	}
}

// DropdownMenu 下拉菜单
type DropdownMenu struct {
	title       string
	items       []string
	interactive *Interactive
}

// NewDropdownMenu 创建一个新的下拉菜单
func (i *Interactive) NewDropdownMenu(title string, items []string) *DropdownMenu {
	return &DropdownMenu{
		title:       title,
		items:       items,
		interactive: i,
	}
}

// Show 显示下拉菜单并返回用户选择的索引
func (menu *DropdownMenu) Show() int {
	options := make([]SelectOption, len(menu.items))
	for i, item := range menu.items {
		options[i] = SelectOption{
			Value: fmt.Sprintf("%d", i),
			Label: item,
		}
	}

	selectField := menu.interactive.NewSelectField(menu.title, options)
	result := selectField.Render()

	index := 0
	fmt.Sscanf(result.Value, "%d", &index)
	return index
}

// 键盘按键常量
const (
	keyArrowUp    = 1000
	keyArrowDown  = 1001
	keyArrowRight = 1002
	keyArrowLeft  = 1003
	keyEnter      = 1004
	keyEsc        = 1005
)

// readKey 读取一个按键
func readKey() int {
	buffer := make([]byte, 3)

	// 读取第一个字节
	n, _ := os.Stdin.Read(buffer[:1])
	if n != 1 {
		return 0
	}

	// 检查是否为回车键 (10是换行符LF, 13是回车符CR)
	if buffer[0] == 10 || buffer[0] == 13 {
		return keyEnter
	}

	if buffer[0] == 27 { // ESC 键
		// 检查是否有后续字节可读
		n, _ = os.Stdin.Read(buffer[1:2])
		if n != 1 { // 单独的ESC键
			return keyEsc
		}

		if buffer[1] != 91 { // 不是 [
			return keyEsc
		}

		// 读取第三个字节 (表示箭头键的方向)
		n, _ = os.Stdin.Read(buffer[2:3])
		if n != 1 {
			return keyEsc
		}

		// 根据第三个字节确定具体的箭头键
		switch buffer[2] {
		case 65:
			return keyArrowUp // 上
		case 66:
			return keyArrowDown // 下
		case 67:
			return keyArrowRight // 右
		case 68:
			return keyArrowLeft // 左
		}

		return keyEsc
	}

	// 其他按键
	return int(buffer[0])
}

// 终端状态
type terminalState struct {
	state []byte
}

// 将终端设置为原始模式
func makeRaw(fd int) (*terminalState, error) {
	var oldState terminalState
	// 使用 stty 命令保存当前设置
	cmd := exec.Command("stty", "-g")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	oldState.state = out

	// 设置为原始模式
	rawCmd := exec.Command("stty", "cbreak", "-echo")
	rawCmd.Stdin = os.Stdin
	err = rawCmd.Run()
	if err != nil {
		return nil, err
	}

	return &oldState, nil
}

// 恢复终端设置
func restoreTerminal(fd int, oldState *terminalState) error {
	if oldState == nil {
		return nil
	}
	cmd := exec.Command("stty", string(oldState.state))
	cmd.Stdin = os.Stdin
	return cmd.Run()
}
