package goterm

import (
	"fmt"
	"strings"
)

// TreeNode 表示树形结构的一个节点
type TreeNode struct {
	Name     string      // 节点名称
	Value    interface{} // 节点值
	Children []*TreeNode // 子节点
}

// Tree 表示一个树形结构
type Tree struct {
	Root *TreeNode // 根节点
}

// NewTree 创建一个新的树形结构
func NewTree(name string, value interface{}) *Tree {
	return &Tree{
		Root: &TreeNode{
			Name:  name,
			Value: value,
		},
	}
}

// AddChild 添加子节点
func (t *TreeNode) AddChild(name string, value interface{}) *TreeNode {
	child := &TreeNode{
		Name:  name,
		Value: value,
	}
	t.Children = append(t.Children, child)
	return child
}

// String 返回树形结构的字符串表示
func (t *Tree) String() string {
	if t.Root == nil {
		return ""
	}

	var sb strings.Builder
	t.Root.buildTreeString(&sb, "", "")
	return sb.String()
}

// buildTreeString 构建树形结构的字符串表示
func (n *TreeNode) buildTreeString(sb *strings.Builder, prefix string, childrenPrefix string) {
	sb.WriteString(prefix)

	// 添加节点名称和值
	sb.WriteString(n.Name)
	if n.Value != nil {
		sb.WriteString(": ")
		sb.WriteString(fmt.Sprint(n.Value))
	}
	sb.WriteString("\n")

	// 处理子节点
	for i, child := range n.Children {
		isLast := i == len(n.Children)-1

		if isLast {
			// 最后一个子节点
			newPrefix := childrenPrefix + "└── "
			newChildrenPrefix := childrenPrefix + "    "
			child.buildTreeString(sb, newPrefix, newChildrenPrefix)
		} else {
			// 非最后一个子节点
			newPrefix := childrenPrefix + "├── "
			newChildrenPrefix := childrenPrefix + "│   "
			child.buildTreeString(sb, newPrefix, newChildrenPrefix)
		}
	}
}

// 保留原来的String方法以兼容现有代码
func (n *TreeNode) String(depth int, isLast bool, isRoot bool) string {
	var sb strings.Builder

	if isRoot {
		// 根节点，无前缀
		n.buildTreeString(&sb, "", "")
	} else {
		// 非根节点，根据depth和isLast生成前缀
		prefix := ""
		childrenPrefix := ""

		// 添加缩进，根据depth参数
		for i := 1; i < depth; i++ {
			prefix += "    "
			childrenPrefix += "    "
		}

		// 添加当前节点的连接符
		if isLast {
			prefix += "└── "
			childrenPrefix += "    "
		} else {
			prefix += "├── "
			childrenPrefix += "│   "
		}

		n.buildTreeString(&sb, prefix, childrenPrefix)
	}

	return sb.String()
}

// Print 打印树形结构
func (t *Tree) Print() {
	fmt.Print(t.String())
}

// PrintWithStyle 使用样式打印树形结构
func (t *Tree) PrintWithStyle(style *Style) {
	fmt.Print(style.Sprint(t.String()))
}

// 示例：创建树形结构
func ExampleTree() {
	// 创建目录结构示例
	dirTree := NewTree("project", nil)
	src := dirTree.Root.AddChild("src", nil)
	src.AddChild("main.go", nil)
	utils := src.AddChild("utils", nil)
	utils.AddChild("logger.go", nil)
	utils.AddChild("style.go", nil)
	examples := src.AddChild("examples", nil)
	examples.AddChild("main.go", nil)
	dirTree.Root.AddChild("README.md", nil)
	dirTree.Root.AddChild("go.mod", nil)

	// 打印目录结构
	dirTree.Print()

	// 创建系统状态示例
	sysTree := NewTree("系统状态", nil)
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

	// 使用样式打印系统状态
	sysTree.PrintWithStyle(New().Bold().Blue())
}
