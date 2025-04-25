package goterm

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

// Chart 是图表接口
type Chart interface {
	// String 返回图表的字符串表示
	String() string
	// Print 打印图表
	Print()
}

// BarChart 表示条形图
type BarChart struct {
	Title       string         // 图表标题
	Data        map[string]int // 数据（标签 -> 值）
	Width       int            // 条形图宽度
	MaxBarWidth int            // 最大条形宽度
	BarStyle    *Style         // 条形样式
	LabelStyle  *Style         // 标签样式
	ValueStyle  *Style         // 值样式
	TitleStyle  *Style         // 标题样式
	BarChar     string         // 用于绘制条形的字符
}

// NewBarChart 创建新的条形图
func NewBarChart() *BarChart {
	return &BarChart{
		Title:       "",
		Data:        make(map[string]int),
		Width:       50,
		MaxBarWidth: 30,
		BarStyle:    New().Cyan(),
		LabelStyle:  New(),
		ValueStyle:  New().Bold(),
		TitleStyle:  New().Bold().Underline(),
		BarChar:     "█",
	}
}

// SetTitle 设置图表标题
func (c *BarChart) SetTitle(title string) *BarChart {
	c.Title = title
	return c
}

// AddData 添加数据项
func (c *BarChart) AddData(label string, value int) *BarChart {
	c.Data[label] = value
	return c
}

// SetData 设置所有数据
func (c *BarChart) SetData(data map[string]int) *BarChart {
	c.Data = data
	return c
}

// SetWidth 设置图表宽度
func (c *BarChart) SetWidth(width int) *BarChart {
	c.Width = width
	return c
}

// SetMaxBarWidth 设置最大条形宽度
func (c *BarChart) SetMaxBarWidth(width int) *BarChart {
	c.MaxBarWidth = width
	return c
}

// SetBarStyle 设置条形样式
func (c *BarChart) SetBarStyle(style *Style) *BarChart {
	c.BarStyle = style
	return c
}

// SetLabelStyle 设置标签样式
func (c *BarChart) SetLabelStyle(style *Style) *BarChart {
	c.LabelStyle = style
	return c
}

// SetValueStyle 设置值样式
func (c *BarChart) SetValueStyle(style *Style) *BarChart {
	c.ValueStyle = style
	return c
}

// SetTitleStyle 设置标题样式
func (c *BarChart) SetTitleStyle(style *Style) *BarChart {
	c.TitleStyle = style
	return c
}

// SetBarChar 设置条形字符
func (c *BarChart) SetBarChar(char string) *BarChart {
	c.BarChar = char
	return c
}

// String 返回条形图的字符串表示
func (c *BarChart) String() string {
	if len(c.Data) == 0 {
		return "空图表（没有数据）"
	}

	var result strings.Builder

	// 添加标题
	if c.Title != "" {
		titleStr := c.TitleStyle.Sprint(c.Title)
		result.WriteString(titleStr + "\n\n")
	}

	// 获取最长标签和最大值
	maxLabelLength := 0
	maxValue := 0
	for label, value := range c.Data {
		if len(label) > maxLabelLength {
			maxLabelLength = len(label)
		}
		if value > maxValue {
			maxValue = value
		}
	}

	// 确保标签长度至少为4个字符
	if maxLabelLength < 4 {
		maxLabelLength = 4
	}

	// 遍历每个数据项
	for label, value := range c.Data {
		// 绘制标签
		labelStr := c.LabelStyle.Sprint(fmt.Sprintf("%-*s", maxLabelLength, label))
		result.WriteString(labelStr)
		result.WriteString(" │ ")

		// 计算条形长度
		barWidth := 0
		if maxValue > 0 {
			barWidth = int((float64(value) / float64(maxValue)) * float64(c.MaxBarWidth))
		}

		// 绘制条形
		bar := strings.Repeat(c.BarChar, barWidth)
		barStr := c.BarStyle.Sprint(bar)
		result.WriteString(barStr)

		// 在条形后显示值
		valueStr := c.ValueStyle.Sprint(fmt.Sprintf(" %d", value))
		result.WriteString(valueStr + "\n")
	}

	return result.String()
}

// Print 打印条形图
func (c *BarChart) Print() {
	fmt.Print(c.String())
}

// PieChart 表示饼图
type PieChart struct {
	Title       string            // 图表标题
	Data        map[string]int    // 数据（标签 -> 值）
	Size        int               // 饼图大小（直径）
	Styles      map[string]*Style // 每个部分的样式
	TitleStyle  *Style            // 标题样式
	LegendStyle *Style            // 图例样式
	ValueStyle  *Style            // 值样式
}

// NewPieChart 创建新的饼图
func NewPieChart() *PieChart {
	return &PieChart{
		Title:       "",
		Data:        make(map[string]int),
		Size:        15,
		Styles:      make(map[string]*Style),
		TitleStyle:  New().Bold().Underline(),
		LegendStyle: New(),
		ValueStyle:  New().Bold(),
	}
}

// SetTitle 设置图表标题
func (p *PieChart) SetTitle(title string) *PieChart {
	p.Title = title
	return p
}

// AddData 添加数据项
func (p *PieChart) AddData(label string, value int) *PieChart {
	p.Data[label] = value
	return p
}

// SetData 设置所有数据
func (p *PieChart) SetData(data map[string]int) *PieChart {
	p.Data = data
	return p
}

// SetSize 设置饼图大小
func (p *PieChart) SetSize(size int) *PieChart {
	if size < 5 {
		size = 5 // 最小大小
	}
	if size > 25 {
		size = 25 // 最大大小
	}
	p.Size = size
	return p
}

// SetStyle 设置某个部分的样式
func (p *PieChart) SetStyle(label string, style *Style) *PieChart {
	p.Styles[label] = style
	return p
}

// SetTitleStyle 设置标题样式
func (p *PieChart) SetTitleStyle(style *Style) *PieChart {
	p.TitleStyle = style
	return p
}

// SetLegendStyle 设置图例样式
func (p *PieChart) SetLegendStyle(style *Style) *PieChart {
	p.LegendStyle = style
	return p
}

// SetValueStyle 设置值样式
func (p *PieChart) SetValueStyle(style *Style) *PieChart {
	p.ValueStyle = style
	return p
}

// String 返回饼图的字符串表示
func (p *PieChart) String() string {
	if len(p.Data) == 0 {
		return "空图表（没有数据）"
	}

	var result strings.Builder

	// 添加标题
	if p.Title != "" {
		titleStr := p.TitleStyle.Sprint(p.Title)
		result.WriteString(titleStr + "\n\n")
	}

	// 计算总值
	total := 0
	for _, value := range p.Data {
		total += value
	}

	// 如果总值为0，不能绘制饼图
	if total == 0 {
		return "无法绘制饼图：所有值都为0"
	}

	// 创建默认样式
	defaultStyles := []*Style{
		New().Red(),
		New().Green(),
		New().Yellow(),
		New().Blue(),
		New().Magenta(),
		New().Cyan(),
		New().White(),
	}

	// 准备数据和角度
	type pieSection struct {
		label      string
		value      int
		percentage float64
		startAngle float64
		endAngle   float64
		style      *Style
	}

	sections := make([]pieSection, 0, len(p.Data))
	currentAngle := 0.0

	i := 0
	for label, value := range p.Data {
		percentage := float64(value) / float64(total)
		sectionAngle := percentage * 360.0
		endAngle := currentAngle + sectionAngle

		// 获取样式
		var style *Style
		if s, ok := p.Styles[label]; ok {
			style = s
		} else {
			style = defaultStyles[i%len(defaultStyles)]
		}

		sections = append(sections, pieSection{
			label:      label,
			value:      value,
			percentage: percentage,
			startAngle: currentAngle,
			endAngle:   endAngle,
			style:      style,
		})

		currentAngle = endAngle
		i++
	}

	// 绘制饼图
	// 增大饼图尺寸，使其看起来更圆
	width := p.Size * 2 // 宽度是高度的2倍，补偿终端字体高宽比
	height := p.Size
	radiusX := float64(width) / 2.0
	radiusY := float64(height) / 2.0
	centerX := radiusX
	centerY := radiusY

	// 创建二维网格
	grid := make([][]string, height+1)
	for i := range grid {
		grid[i] = make([]string, width+1)
		for j := range grid[i] {
			grid[i][j] = " "
		}
	}

	// 计算椭圆填充像素
	for y := 0; y <= height; y++ {
		for x := 0; x <= width; x++ {
			// 计算点到中心的标准化距离（椭圆方程）
			relX := float64(x) - centerX
			relY := float64(y) - centerY

			// 使用椭圆方程计算是否在圆内
			// (x/a)² + (y/b)² <= 1
			normalizedDistance := (relX*relX)/(radiusX*radiusX) + (relY*relY)/(radiusY*radiusY)

			if normalizedDistance > 1.0 {
				continue // 在圆外
			}

			// 计算角度（弧度）
			angle := math.Atan2(relY, relX) * 180 / math.Pi
			if angle < 0 {
				angle += 360 // 转换为0-360度
			}

			// 确定所属扇区并选择合适的字符
			for _, section := range sections {
				if angle >= section.startAngle && angle < section.endAngle {
					// 根据到中心的距离选择不同字符，使边缘更圆滑
					var char string

					// 边缘判定，使饼图边缘更柔和
					if normalizedDistance > 0.85 {
						char = "·" // 边缘使用较小的点
					} else {
						char = "●" // 中间部分使用实心圆形
					}

					grid[y][x] = section.style.Sprint(char)
					break
				}
			}
		}
	}

	// 输出饼图
	for _, row := range grid {
		for _, cell := range row {
			result.WriteString(cell)
		}
		result.WriteString("\n")
	}

	result.WriteString("\n图例:\n")

	// 创建图例
	for _, section := range sections {
		legendSymbol := section.style.Sprint("●")
		labelText := p.LegendStyle.Sprint(section.label)
		percentage := p.ValueStyle.Sprint(fmt.Sprintf("%.1f%%", section.percentage*100))
		value := p.ValueStyle.Sprint(fmt.Sprintf("(%d)", section.value))

		result.WriteString(fmt.Sprintf("%s %s: %s %s\n", legendSymbol, labelText, percentage, value))
	}

	return result.String()
}

// Print 打印饼图
func (p *PieChart) Print() {
	fmt.Print(p.String())
}

// LineChart 表示折线图
type LineChart struct {
	Title      string             // 图表标题
	Data       map[string][]Point // 数据（系列名 -> 点集合）
	Width      int                // 图表宽度
	Height     int                // 图表高度
	XAxisTitle string             // X轴标题
	YAxisTitle string             // Y轴标题
	ShowGrid   bool               // 是否显示网格

	// 样式设置
	TitleStyle   *Style            // 标题样式
	AxisStyle    *Style            // 坐标轴样式
	GridStyle    *Style            // 网格样式
	LegendStyle  *Style            // 图例样式
	LineStyles   map[string]*Style // 每条线的样式
	MarkerStyles map[string]*Style // 每条线标记点的样式
}

// Point 表示折线图中的点
type Point struct {
	X float64
	Y float64
}

// NewLineChart 创建新的折线图
func NewLineChart() *LineChart {
	return &LineChart{
		Title:        "",
		Data:         make(map[string][]Point),
		Width:        60,
		Height:       15,
		XAxisTitle:   "",
		YAxisTitle:   "",
		ShowGrid:     true,
		TitleStyle:   New().Bold().Underline(),
		AxisStyle:    New(),
		GridStyle:    New().Faint(),
		LegendStyle:  New(),
		LineStyles:   make(map[string]*Style),
		MarkerStyles: make(map[string]*Style),
	}
}

// SetTitle 设置图表标题
func (c *LineChart) SetTitle(title string) *LineChart {
	c.Title = title
	return c
}

// AddPoint 添加单个数据点到指定系列
func (c *LineChart) AddPoint(series string, x, y float64) *LineChart {
	if _, ok := c.Data[series]; !ok {
		c.Data[series] = []Point{}
	}
	c.Data[series] = append(c.Data[series], Point{X: x, Y: y})
	return c
}

// AddSeries 添加完整的数据系列
func (c *LineChart) AddSeries(series string, points []Point) *LineChart {
	c.Data[series] = points
	return c
}

// SetWidth 设置图表宽度
func (c *LineChart) SetWidth(width int) *LineChart {
	c.Width = width
	return c
}

// SetHeight 设置图表高度
func (c *LineChart) SetHeight(height int) *LineChart {
	c.Height = height
	return c
}

// SetAxisTitles 设置坐标轴标题
func (c *LineChart) SetAxisTitles(x, y string) *LineChart {
	c.XAxisTitle = x
	c.YAxisTitle = y
	return c
}

// SetShowGrid 设置是否显示网格
func (c *LineChart) SetShowGrid(show bool) *LineChart {
	c.ShowGrid = show
	return c
}

// SetTitleStyle 设置标题样式
func (c *LineChart) SetTitleStyle(style *Style) *LineChart {
	c.TitleStyle = style
	return c
}

// SetAxisStyle 设置坐标轴样式
func (c *LineChart) SetAxisStyle(style *Style) *LineChart {
	c.AxisStyle = style
	return c
}

// SetGridStyle 设置网格样式
func (c *LineChart) SetGridStyle(style *Style) *LineChart {
	c.GridStyle = style
	return c
}

// SetLineStyle 设置指定系列的线条样式
func (c *LineChart) SetLineStyle(series string, style *Style) *LineChart {
	c.LineStyles[series] = style
	return c
}

// SetMarkerStyle 设置指定系列的标记点样式
func (c *LineChart) SetMarkerStyle(series string, style *Style) *LineChart {
	c.MarkerStyles[series] = style
	return c
}

// String 返回折线图的字符串表示
func (c *LineChart) String() string {
	if len(c.Data) == 0 {
		return "空图表（没有数据）"
	}

	var result strings.Builder

	// 添加标题
	if c.Title != "" {
		titleStr := c.TitleStyle.Sprint(c.Title)
		result.WriteString(titleStr + "\n\n")
	}

	// 设置默认线型和标记
	defaultLineStyles := []*Style{
		New().Red(),
		New().Green(),
		New().Yellow(),
		New().Blue(),
		New().Magenta(),
		New().Cyan(),
	}

	defaultMarkers := []string{"●", "■", "▲", "◆", "✦", "✱"}

	// 找出所有点的X和Y的最大最小值
	var xMin, xMax, yMin, yMax float64
	var allXValues []float64
	first := true

	for series, points := range c.Data {
		if len(points) == 0 {
			continue
		}

		// 为每个系列设置默认样式
		if _, ok := c.LineStyles[series]; !ok {
			seriesIndex := len(c.LineStyles) % len(defaultLineStyles)
			c.LineStyles[series] = defaultLineStyles[seriesIndex]
		}

		// 为每个系列设置默认标记
		if _, ok := c.MarkerStyles[series]; !ok {
			seriesIndex := len(c.MarkerStyles) % len(defaultLineStyles)
			markerStyle := defaultLineStyles[seriesIndex]
			c.MarkerStyles[series] = markerStyle
		}

		for _, p := range points {
			if first {
				xMin, xMax = p.X, p.X
				yMin, yMax = p.Y, p.Y
				first = false
			} else {
				xMin = math.Min(xMin, p.X)
				xMax = math.Max(xMax, p.X)
				yMin = math.Min(yMin, p.Y)
				yMax = math.Max(yMax, p.Y)
			}
			allXValues = append(allXValues, p.X)
		}
	}

	// 确保Y轴从0开始，除非有负值
	if yMin > 0 {
		yMin = 0
	}

	// 确保X和Y的范围不为0
	if xMax == xMin {
		xMax = xMin + 1
	}
	if yMax == yMin {
		yMax = yMin + 1
	}

	// 轴留出的边距
	leftMargin := 8   // 左边距，用于Y轴标签
	bottomMargin := 2 // 底部边距，用于X轴标签

	// 图表实际绘制区域
	plotWidth := c.Width - leftMargin
	plotHeight := c.Height - bottomMargin

	// 创建绘图区域（包括坐标轴）
	gridChars := make([][]string, c.Height)
	for i := range gridChars {
		gridChars[i] = make([]string, c.Width)
		for j := range gridChars[i] {
			gridChars[i][j] = " "
		}
	}

	// 绘制Y轴
	for i := 0; i < c.Height-bottomMargin; i++ {
		yPos := c.Height - bottomMargin - 1 - i
		gridChars[yPos][leftMargin-1] = c.AxisStyle.Sprint("│")

		// 绘制Y轴刻度和标签
		if i%(plotHeight/4) == 0 || i == plotHeight-1 {
			gridChars[yPos][leftMargin-1] = c.AxisStyle.Sprint("┤")

			// 计算当前Y值
			y := yMin + (float64(i)/float64(plotHeight-1))*(yMax-yMin)
			yLabel := fmt.Sprintf("%5.1f", y)

			// 添加Y轴标签
			for j := 0; j < len(yLabel) && j < leftMargin-1; j++ {
				gridChars[yPos][j] = c.AxisStyle.Sprint(string(yLabel[j]))
			}
		}

		// 绘制水平网格线
		if c.ShowGrid && i > 0 && i < plotHeight {
			for j := 0; j < plotWidth; j++ {
				if gridChars[yPos][leftMargin+j] == " " {
					if i%(plotHeight/4) == 0 {
						gridChars[yPos][leftMargin+j] = c.GridStyle.Sprint("─")
					} else {
						gridChars[yPos][leftMargin+j] = c.GridStyle.Sprint("·")
					}
				}
			}
		}
	}

	// 绘制X轴
	for j := 0; j < plotWidth; j++ {
		xPos := leftMargin + j
		gridChars[c.Height-bottomMargin][xPos] = c.AxisStyle.Sprint("─")

		// 绘制X轴刻度和标签
		if j%(plotWidth/5) == 0 || j == plotWidth-1 {
			gridChars[c.Height-bottomMargin][xPos] = c.AxisStyle.Sprint("┬")

			// 计算当前X值
			x := xMin + (float64(j)/float64(plotWidth-1))*(xMax-xMin)
			xLabel := fmt.Sprintf("%.1f", x)

			// 添加X轴标签（居中对齐）
			for k := 0; k < len(xLabel) && xPos+k-len(xLabel)/2 < c.Width && xPos+k-len(xLabel)/2 >= 0; k++ {
				charPos := xPos + k - len(xLabel)/2
				if charPos < c.Width {
					gridChars[c.Height-bottomMargin+1][charPos] = c.AxisStyle.Sprint(string(xLabel[k]))
				}
			}
		}

		// 绘制垂直网格线
		if c.ShowGrid && j > 0 && j < plotWidth {
			for i := 0; i < plotHeight; i++ {
				yPos := c.Height - bottomMargin - 1 - i
				if gridChars[yPos][xPos] == " " || gridChars[yPos][xPos] == "·" {
					if j%(plotWidth/5) == 0 {
						gridChars[yPos][xPos] = c.GridStyle.Sprint("│")
					} else if gridChars[yPos][xPos] != "─" {
						gridChars[yPos][xPos] = c.GridStyle.Sprint("·")
					}
				}
			}
		}
	}

	// 标记轴的交点
	gridChars[c.Height-bottomMargin][leftMargin-1] = c.AxisStyle.Sprint("┼")

	// 绘制每个数据系列的线条
	seriesIndex := 0
	for series, points := range c.Data {
		if len(points) < 2 {
			continue
		}

		// 获取线条和标记样式
		lineStyle := c.LineStyles[series]
		markerStyle := c.MarkerStyles[series]
		marker := defaultMarkers[seriesIndex%len(defaultMarkers)]

		// 按X值排序点
		sort.Slice(points, func(i, j int) bool {
			return points[i].X < points[j].X
		})

		// 绘制线条和标记点
		var lastX, lastY int
		first := true

		for _, p := range points {
			// 将数据点映射到绘图区域坐标
			x := int(math.Round(float64(plotWidth-1) * (p.X - xMin) / (xMax - xMin)))
			y := int(math.Round(float64(plotHeight-1) * (1 - (p.Y-yMin)/(yMax-yMin))))

			// 确保点在绘图区域内
			if x < 0 {
				x = 0
			} else if x >= plotWidth {
				x = plotWidth - 1
			}

			if y < 0 {
				y = 0
			} else if y >= plotHeight {
				y = plotHeight - 1
			}

			x += leftMargin
			y = c.Height - bottomMargin - 1 - y

			// 绘制标记点
			gridChars[y][x] = markerStyle.Sprint(marker)

			// 绘制线段连接当前点和上一个点
			if !first {
				// 简单的线段绘制算法（Bresenham算法的简化版）
				dx := x - lastX
				dy := y - lastY
				steps := int(math.Max(math.Abs(float64(dx)), math.Abs(float64(dy))))

				for i := 1; i < steps; i++ {
					tx := lastX + i*dx/steps
					ty := lastY + i*dy/steps

					if tx >= 0 && tx < c.Width && ty >= 0 && ty < c.Height {
						// 使用适当的字符来表示线条方向
						lineChar := "─"
						if dx == 0 {
							lineChar = "│"
						} else if dy == 0 {
							lineChar = "─"
						} else if (dx > 0 && dy > 0) || (dx < 0 && dy < 0) {
							lineChar = "\\"
						} else {
							lineChar = "/"
						}

						// 只在空白处绘制线条
						if gridChars[ty][tx] == " " ||
							gridChars[ty][tx] == "·" ||
							gridChars[ty][tx] == "─" ||
							gridChars[ty][tx] == "│" {
							gridChars[ty][tx] = lineStyle.Sprint(lineChar)
						}
					}
				}
			}

			lastX, lastY = x, y
			first = false
		}

		seriesIndex++
	}

	// 渲染图表
	for i := 0; i < c.Height; i++ {
		for j := 0; j < c.Width; j++ {
			result.WriteString(gridChars[i][j])
		}
		result.WriteString("\n")
	}

	// 添加图例
	result.WriteString("\n图例: ")
	seriesIndex = 0
	for series := range c.Data {
		lineStyle := c.LineStyles[series]
		marker := defaultMarkers[seriesIndex%len(defaultMarkers)]
		result.WriteString(lineStyle.Sprint(marker + " " + series + " "))
		seriesIndex++
	}
	result.WriteString("\n")

	return result.String()
}

// Print 打印折线图
func (c *LineChart) Print() {
	fmt.Print(c.String())
}
