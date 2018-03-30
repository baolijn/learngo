package main

import (
	"learngo/tree"
	"fmt"
)

type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode)printPostOrder() {
	if myNode == nil || myNode.node == nil{
		return
	}
	left := myTreeNode{myNode.node.Left }
	left.printPostOrder()
	right := myTreeNode{myNode.node.Right}
	right.printPostOrder()
	myNode.node.Print()
}

func main() {
	var root tree.Node
	root = tree.Node{Value : 3}
	root.Left = &tree.Node{}
	root.Left.Right = tree.CreateTreeNode(2)
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Right.Left.SetValue(4)
	root.Traverse()

	nodeCount := 0
	root.TraverseFunc(func(node *tree.Node) {
		nodeCount++
	})

	fmt.Println("Node count:", nodeCount)
    //myRoot := myTreeNode{&root}
	//myRoot.printPostOrder()

	//nodes := []treeNode{
	//	{value:3},
	//	{},
	//	{6, nil, &root},
	//}
	//root.print()
	//root.Left.SetValue(4)
	//root.left.print()
	//fmt.Println(nodes)

	//pRoot := &root
	//pRoot.setValue(200)
	//pRoot.print()

	//var pRoot *tree.Node
	//pRoot.SetValue(200)
	//pRoot = &root
	//pRoot.SetValue(300)
	//pRoot.Print()
	c := root.TraverseWithChannel()
	maxValue := 0
	for node := range c {
		if node.Value > maxValue {
			maxValue = node.Value
		}
	}
	fmt.Println("Max node value is ", maxValue)
}
