package main

import "learngo/tree"

func main() {
	var root tree.Node
	root = tree.Node{Value : 3}
	root.Left = &tree.Node{}
	root.Left.Right = tree.CreateTreeNode(2)
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Right.Left.SetValue(4)

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
	root.Traverse()
}
