package goterm

import (
	"fmt"
	"strings"
)

// 对齐方式
type Alignment int

const (
	AlignLeft   Alignment = iota // 左对齐
	AlignCenter                  // 居中对齐
	AlignRight                   // 右对齐
)

// TableColumn 表示表格的列
type TableColumn struct {
	Header    string    // 列标题
	Alignment Alignment // 对齐方式
	MinWidth  int       // 最小宽度
	MaxWidth  int       // 最大宽度（0表示不限制）
}

// NewColumn 创建一个新的列
func NewColumn(header string) *TableColumn {
	return &TableColumn{
		Header:    header,
		Alignment: AlignLeft,
		MinWidth:  0,
		MaxWidth:  0,
	}
}

// SetAlignment 设置列的对齐方式
func (c *TableColumn) SetAlignment(alignment Alignment) *TableColumn {
	c.Alignment = alignment
	return c
}

// SetMinWidth 设置列的最小宽度
func (c *TableColumn) SetMinWidth(width int) *TableColumn {
	c.MinWidth = width
	return c
}

// SetMaxWidth 设置列的最大宽度
func (c *TableColumn) SetMaxWidth(width int) *TableColumn {
	c.MaxWidth = width
	return c
}

// Table 表示一个格式化表格
type Table struct {
	Columns         []TableColumn // 列定义
	Rows            [][]string    // 行数据
	Border          *Style        // 边框样式
	Header          *Style        // 表头样式
	Row             *Style        // 行样式
	HasBorder       bool          // 是否显示边框
	HasRowSeparator bool          // 是否显示行之间的分隔线
}

// NewEmptyTable 创建一个没有列的新表格
func NewEmptyTable() *Table {
	return &Table{
		Columns:         make([]TableColumn, 0),
		Rows:            make([][]string, 0),
		Border:          New(),
		Header:          New(),
		Row:             New(),
		HasBorder:       true,
		HasRowSeparator: false,
	}
}

// AddColumn 添加一列到表格
func (t *Table) AddColumn(column *TableColumn) *Table {
	t.Columns = append(t.Columns, *column)
	return t
}

// AddRow 添加一行数据
func (t *Table) AddRow(cells ...string) {
	// 确保cells长度与列数相同
	row := make([]string, len(t.Columns))
	for i := 0; i < len(row); i++ {
		if i < len(cells) {
			row[i] = cells[i]
		}
	}
	t.Rows = append(t.Rows, row)
}

// 计算字符串的显示宽度，考虑全角字符（如中文）占用两个字符位置
func displayWidth(s string) int {
	width := 0
	for _, r := range s {
		if r > 0x7F { // ASCII范围外的字符（如中文）
			width += 2
		} else {
			width += 1
		}
	}
	return width
}

// calculateColumnWidths 计算每列的实际宽度
func (t *Table) calculateColumnWidths() []int {
	// 初始化为每列最小宽度
	widths := make([]int, len(t.Columns))
	for i, col := range t.Columns {
		widths[i] = col.MinWidth
		// 检查表头宽度
		headerWidth := displayWidth(col.Header)
		if headerWidth > widths[i] {
			widths[i] = headerWidth
		}
	}

	// 检查每行中每个单元格的宽度
	for _, row := range t.Rows {
		for i, cell := range row {
			if i >= len(widths) {
				continue
			}
			cellWidth := displayWidth(cell)
			if cellWidth > widths[i] && (t.Columns[i].MaxWidth == 0 || cellWidth <= t.Columns[i].MaxWidth) {
				widths[i] = cellWidth
			}
		}
	}

	// 应用最大宽度限制
	for i, col := range t.Columns {
		if col.MaxWidth > 0 && widths[i] > col.MaxWidth {
			widths[i] = col.MaxWidth
		}
	}

	return widths
}

// formatCell 格式化单元格内容，根据对齐方式和宽度
func formatCell(content string, width int, align Alignment) string {
	// 计算内容的实际显示宽度（考虑中文字符占用2个宽度的情况）
	displayWidth := 0
	for _, r := range content {
		if r > 0x7F { // ASCII范围外的字符（如中文）
			displayWidth += 2
		} else {
			displayWidth += 1
		}
	}

	// 如果内容宽度超过列宽度，截断内容
	if displayWidth > width {
		// 安全截取UTF-8字符串
		result := ""
		currentWidth := 0

		// 预留省略号的宽度（3个ASCII字符）
		targetWidth := width - 3
		if targetWidth < 0 {
			targetWidth = 0
		}

		for _, r := range content {
			charWidth := 1
			if r > 0x7F {
				charWidth = 2
			}

			// 如果添加这个字符会超出目标宽度，停止添加
			if currentWidth+charWidth > targetWidth {
				break
			}

			result += string(r)
			currentWidth += charWidth
		}

		// 添加省略号
		result += "..."
		currentWidth += 3

		// 根据对齐方式填充空格到指定宽度
		remainingSpace := width - currentWidth
		if remainingSpace > 0 {
			switch align {
			case AlignLeft:
				return result + strings.Repeat(" ", remainingSpace)
			case AlignRight:
				return strings.Repeat(" ", remainingSpace) + result
			case AlignCenter:
				left := remainingSpace / 2
				right := remainingSpace - left
				return strings.Repeat(" ", left) + result + strings.Repeat(" ", right)
			default:
				return result + strings.Repeat(" ", remainingSpace)
			}
		}

		return result
	}

	// 根据对齐方式填充空格
	padding := width - displayWidth
	if padding < 0 {
		padding = 0
	}

	switch align {
	case AlignLeft:
		return content + strings.Repeat(" ", padding)
	case AlignRight:
		return strings.Repeat(" ", padding) + content
	case AlignCenter:
		left := padding / 2
		right := padding - left
		return strings.Repeat(" ", left) + content + strings.Repeat(" ", right)
	default:
		return content + strings.Repeat(" ", padding)
	}
}

// String 返回表格的字符串表示
func (t *Table) String() string {
	if len(t.Columns) == 0 {
		return ""
	}

	// 计算每列的宽度
	widths := t.calculateColumnWidths()

	var sb strings.Builder

	// 绘制顶部边框
	if t.HasBorder {
		sb.WriteString("┌")
		for i, width := range widths {
			sb.WriteString(strings.Repeat("─", width+2))
			if i < len(widths)-1 {
				sb.WriteString("┬")
			}
		}
		sb.WriteString("┐\n")
	}

	// 绘制表头
	if t.HasBorder {
		sb.WriteString("│ ")
	}
	for i, col := range t.Columns {
		header := formatCell(col.Header, widths[i], AlignCenter)
		sb.WriteString(header)
		if i < len(t.Columns)-1 {
			if t.HasBorder {
				sb.WriteString(" │ ")
			} else {
				sb.WriteString("  ")
			}
		}
	}
	if t.HasBorder {
		sb.WriteString(" │")
	}
	sb.WriteString("\n")

	// 绘制表头与数据之间的分隔线
	if t.HasBorder {
		sb.WriteString("├")
		for i, width := range widths {
			sb.WriteString(strings.Repeat("─", width+2))
			if i < len(widths)-1 {
				sb.WriteString("┼")
			}
		}
		sb.WriteString("┤\n")
	}

	// 绘制数据行
	for rowIndex, row := range t.Rows {
		if t.HasBorder {
			sb.WriteString("│ ")
		}
		for i, cell := range row {
			if i >= len(t.Columns) {
				continue
			}
			formattedCell := formatCell(cell, widths[i], t.Columns[i].Alignment)
			sb.WriteString(formattedCell)
			if i < len(t.Columns)-1 {
				if t.HasBorder {
					sb.WriteString(" │ ")
				} else {
					sb.WriteString("  ")
				}
			}
		}
		if t.HasBorder {
			sb.WriteString(" │")
		}
		sb.WriteString("\n")

		// 添加行之间的分隔线（除了最后一行之后）
		if t.HasRowSeparator && t.HasBorder && rowIndex < len(t.Rows)-1 {
			sb.WriteString("├")
			for i, width := range widths {
				sb.WriteString(strings.Repeat("─", width+2))
				if i < len(widths)-1 {
					sb.WriteString("┼")
				}
			}
			sb.WriteString("┤\n")
		}
	}

	// 绘制底部边框
	if t.HasBorder {
		sb.WriteString("└")
		for i, width := range widths {
			sb.WriteString(strings.Repeat("─", width+2))
			if i < len(widths)-1 {
				sb.WriteString("┴")
			}
		}
		sb.WriteString("┘\n")
	}

	return sb.String()
}

// Print 打印表格
func (t *Table) Print() {
	fmt.Print(t.String())
}

// PrintWithStyle 使用样式打印表格
func (t *Table) PrintWithStyle() {
	// 如果没有样式，直接使用Print
	if t.Border == nil && t.Header == nil && t.Row == nil {
		t.Print()
		return
	}

	if len(t.Columns) == 0 {
		return
	}

	// 计算每列的宽度
	widths := t.calculateColumnWidths()

	// 绘制顶部边框
	if t.HasBorder {
		var borderLine strings.Builder
		borderLine.WriteString("┌")
		for i, width := range widths {
			borderLine.WriteString(strings.Repeat("─", width+2))
			if i < len(widths)-1 {
				borderLine.WriteString("┬")
			}
		}
		borderLine.WriteString("┐")

		// 打印顶部边框
		if t.Border != nil {
			fmt.Println(t.Border.Sprint(borderLine.String()))
		} else {
			fmt.Println(borderLine.String())
		}
	}

	// 绘制表头
	if t.HasBorder {
		if t.Border != nil {
			fmt.Print(t.Border.Sprint("│"))
		} else {
			fmt.Print("│")
		}
		fmt.Print(" ")
	}

	// 打印表头内容
	for i, col := range t.Columns {
		cellContent := formatCell(col.Header, widths[i], AlignCenter)

		// 应用表头样式
		if t.Header != nil {
			fmt.Print(t.Header.Sprint(cellContent))
		} else {
			fmt.Print(cellContent)
		}

		if i < len(t.Columns)-1 {
			if t.HasBorder {
				if t.Border != nil {
					fmt.Print(" " + t.Border.Sprint("│") + " ")
				} else {
					fmt.Print(" │ ")
				}
			} else {
				fmt.Print("  ")
			}
		}
	}

	if t.HasBorder {
		fmt.Print(" ")
		if t.Border != nil {
			fmt.Println(t.Border.Sprint("│"))
		} else {
			fmt.Println("│")
		}
	} else {
		fmt.Println()
	}

	// 绘制表头与数据之间的分隔线
	if t.HasBorder {
		var borderLine strings.Builder
		borderLine.WriteString("├")
		for i, width := range widths {
			borderLine.WriteString(strings.Repeat("─", width+2))
			if i < len(widths)-1 {
				borderLine.WriteString("┼")
			}
		}
		borderLine.WriteString("┤")

		// 打印分隔线
		if t.Border != nil {
			fmt.Println(t.Border.Sprint(borderLine.String()))
		} else {
			fmt.Println(borderLine.String())
		}
	}

	// 绘制数据行
	for rowIndex, row := range t.Rows {
		if t.HasBorder {
			if t.Border != nil {
				fmt.Print(t.Border.Sprint("│"))
			} else {
				fmt.Print("│")
			}
			fmt.Print(" ")
		}

		for i, cell := range row {
			if i >= len(t.Columns) {
				continue
			}

			cellContent := formatCell(cell, widths[i], t.Columns[i].Alignment)

			// 应用行样式
			if t.Row != nil {
				fmt.Print(t.Row.Sprint(cellContent))
			} else {
				fmt.Print(cellContent)
			}

			if i < len(t.Columns)-1 {
				if t.HasBorder {
					if t.Border != nil {
						fmt.Print(" " + t.Border.Sprint("│") + " ")
					} else {
						fmt.Print(" │ ")
					}
				} else {
					fmt.Print("  ")
				}
			}
		}

		if t.HasBorder {
			fmt.Print(" ")
			if t.Border != nil {
				fmt.Println(t.Border.Sprint("│"))
			} else {
				fmt.Println("│")
			}
		} else {
			fmt.Println()
		}

		// 添加行之间的分隔线（除了最后一行之后）
		if t.HasRowSeparator && t.HasBorder && rowIndex < len(t.Rows)-1 {
			var borderLine strings.Builder
			borderLine.WriteString("├")
			for i, width := range widths {
				borderLine.WriteString(strings.Repeat("─", width+2))
				if i < len(widths)-1 {
					borderLine.WriteString("┼")
				}
			}
			borderLine.WriteString("┤")

			// 打印行分隔线
			if t.Border != nil {
				fmt.Println(t.Border.Sprint(borderLine.String()))
			} else {
				fmt.Println(borderLine.String())
			}
		}
	}

	// 绘制底部边框
	if t.HasBorder {
		var borderLine strings.Builder
		borderLine.WriteString("└")
		for i, width := range widths {
			borderLine.WriteString(strings.Repeat("─", width+2))
			if i < len(widths)-1 {
				borderLine.WriteString("┴")
			}
		}
		borderLine.WriteString("┘")

		// 打印底部边框
		if t.Border != nil {
			fmt.Println(t.Border.Sprint(borderLine.String()))
		} else {
			fmt.Println(borderLine.String())
		}
	}
}

// PrintStyled 使用提供的样式打印表格
func (t *Table) PrintStyled(borderStyle, headerStyle, rowStyle *Style) {
	// 保存当前样式
	oldBorderStyle := t.Border
	oldHeaderStyle := t.Header
	oldRowStyle := t.Row

	// 应用提供的样式
	t.Border = borderStyle
	t.Header = headerStyle
	t.Row = rowStyle

	// 打印表格
	t.PrintWithStyle()

	// 恢复原来的样式
	t.Border = oldBorderStyle
	t.Header = oldHeaderStyle
	t.Row = oldRowStyle
}

// PrintStyledTable 以更灵活的方式打印带样式的表格
func PrintStyledTable(table *Table, title string, titleStyle, borderStyle *Style) {
	if titleStyle != nil && title != "" {
		fmt.Println(titleStyle.Sprint(title))
	} else if title != "" {
		fmt.Println(title)
	}

	if borderStyle != nil {
		table.Border = borderStyle
	}

	table.PrintWithStyle()
}

// SetBorderStyle 设置表格边框样式
func (t *Table) SetBorderStyle(style *Style) *Table {
	t.Border = style
	return t
}

// SetHeaderStyle 设置表格表头样式
func (t *Table) SetHeaderStyle(style *Style) *Table {
	t.Header = style
	return t
}

// SetRowStyle 设置表格行样式
func (t *Table) SetRowStyle(style *Style) *Table {
	t.Row = style
	return t
}

// SetHasBorder 设置表格是否显示边框
func (t *Table) SetHasBorder(hasBorder bool) *Table {
	t.HasBorder = hasBorder
	return t
}

// SetHasRowSeparator 设置表格是否显示行之间的分隔线
func (t *Table) SetHasRowSeparator(hasRowSeparator bool) *Table {
	t.HasRowSeparator = hasRowSeparator
	return t
}

// ResetStyles 重置所有样式设置
func (t *Table) ResetStyles() *Table {
	t.Border = nil
	t.Header = nil
	t.Row = nil
	return t
}

// WithColumns 设置表格列
func (t *Table) WithColumns(columns []TableColumn) *Table {
	t.Columns = columns
	return t
}

// WithRows 设置表格行数据
func (t *Table) WithRows(rows [][]string) *Table {
	t.Rows = rows
	return t
}

// Clone 复制表格
func (t *Table) Clone() *Table {
	newTable := &Table{
		Columns:         make([]TableColumn, len(t.Columns)),
		Rows:            make([][]string, len(t.Rows)),
		Border:          t.Border,
		Header:          t.Header,
		Row:             t.Row,
		HasBorder:       t.HasBorder,
		HasRowSeparator: t.HasRowSeparator,
	}

	// 复制列定义
	copy(newTable.Columns, t.Columns)

	// 复制行数据
	for i, row := range t.Rows {
		newRow := make([]string, len(row))
		copy(newRow, row)
		newTable.Rows[i] = newRow
	}

	return newTable
}
