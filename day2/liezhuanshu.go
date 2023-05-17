package main

import "fmt"

type Node struct {
	ID       int
	ParentID int
	Name     string
	Children []*Node
}

func main() {
	nodes := []*Node{
		{ID: 1, ParentID: 0, Name: "Root Node"},
		{ID: 2, ParentID: 1, Name: "Child Node 1"},
		{ID: 3, ParentID: 1, Name: "Child Node 2"},
		{ID: 4, ParentID: 2, Name: "Grandchild Node 1"},
		{ID: 5, ParentID: 2, Name: "Grandchild Node 2"},
		{ID: 6, ParentID: 3, Name: "Grandchild Node 3"},
	}

	rootNodes := BuildTree(nodes)

	// 遍历树形结构并打印节点
	Traverse(rootNodes, 0)
}

func BuildTree(nodes []*Node) []*Node {
	// 构建节点映射
	m := make(map[int]*Node)
	for _, node := range nodes {
		m[node.ID] = node
	}

	// 构建树形结构
	var rootNode []*Node
	for _, node := range nodes {
		if node.ParentID == 0 {
			rootNode = append(rootNode, node)
			continue
		}

		parentNode, ok := m[node.ParentID]
		if !ok {
			continue
		}
		parentNode.Children = append(parentNode.Children, node)

	}

	return rootNode
}

// Traverse 遍历树形结构并打印节点
func Traverse(nodes []*Node, level int) {
	for _, node := range nodes {
		fmt.Printf("%s- %s\n", getIndent(level), node.Name)
		Traverse(node.Children, level+1)
	}
}

func getIndent(level int) string {
	var indent string
	for i := 0; i < level; i++ {
		indent += "  "
	}
	return indent
}
