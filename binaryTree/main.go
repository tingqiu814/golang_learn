package main

import "fmt"

// Node 节点
type Node struct {
	Left  *Node
	Right *Node
	Value int
}

var arr []int

func main() {
	node1 := new(Node)
	node1.Value = 10
	node1.Insert(101)
	node1.Insert(8)
	node1.Insert(12)
	node1.Insert(9)
	printNode(node1)
	x := inOrder(node1)
	fmt.Println(x)
}
func printNode(n *Node) {
	if n.Left != nil {
		printNode(n.Left)
	}
	fmt.Printf(" %v", n.Value)
	if n.Right != nil {
		printNode(n.Right)
	}
}

// Insert 插入值
func (t *Node) Insert(v int) {
	node := new(Node)
	node.Value = v

	if t == nil {
		t = node
	} else {
		InsertNode(t, node)
	}

}

// InsertNode 插入节点
func InsertNode(n *Node, newNode *Node) {
	if newNode.Value < n.Value {
		if n.Left == nil {
			n.Left = newNode
		} else {
			InsertNode(n.Left, newNode)
		}
	} else {
		if n.Right == nil {
			n.Right = newNode
		} else {
			InsertNode(n.Right, newNode)
		}
	}
}

// 将排序
func inOrder(n *Node) []int {
	if n != nil {
		inOrder(n.Left)
		arr = append(arr, n.Value)
		inOrder(n.Right)
	}
	return arr
}
