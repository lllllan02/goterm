package goterm

import (
	"fmt"
	"strings"
)

// TreeNode 表示树形结构的一个节点
type TreeNode struct {
	Name     string      // 节点名称
	Value    any         // 节点值
	Children []*TreeNode // 子节点
}

// Tree 表示一个树形结构
type Tree struct {
	Root *TreeNode // 根节点
}

// NewTree 创建一个新的树形结构
func NewTree(name string, value any) *Tree {
	return &Tree{
		Root: &TreeNode{
			Name:  name,
			Value: value,
		},
	}
}

// AddChild 添加子节点
func (t *TreeNode) AddChild(name string, value any) *TreeNode {
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
