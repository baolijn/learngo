package main

import (
	"fmt"
)

type treeNode struct {
	value int
	left, right *treeNode
}

func (node treeNode) print() {
	fmt.Println(node.value, " ")
}

func (node *treeNode) setValue(value int)  {
	if node == nil {
		fmt.Println("Invalid Node!")
		return
	}
	node.value = value
}

func (node *treeNode)traverse()  {
	if node == nil{
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
}

func createTreeNode(value int) *treeNode {
	return &treeNode{value:value}

}
func main() {
	var root treeNode
	root = treeNode{value : 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)

	nodes := []treeNode{
		{value:3},
		{},
		{6, nil, &root},
	}
	root.print()
	root.left.setValue(4)
	root.left.print()
	fmt.Println(nodes)

	//pRoot := &root
	//pRoot.setValue(200)
	//pRoot.print()

	var pRoot *treeNode
	pRoot.setValue(200)
	pRoot = &root
	pRoot.setValue(300)
	pRoot.print()
	root.traverse()
}
