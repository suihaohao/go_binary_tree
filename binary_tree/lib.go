package binary_tree

import (
	"math"
)

func insertNode(root, newNode *Node) {
	if newNode.Value < root.Value {
		if root.Left == nil {
			root.Left = newNode
		} else {
			insertNode(root.Left, newNode)
		}
	} else if newNode.Value > root.Value {
		if root.Right == nil {
			root.Right = newNode
		} else {
			insertNode(root.Right, newNode)
		}
	}
}

func remove(root *Node, value int) (*Node, bool) {
	if root == nil {
		return nil, false
	}
	var existed bool
	if value < root.Value {
		root.Left, existed = remove(root.Left, value)
		return root, existed
	}
	if value > root.Value {
		root.Right, existed = remove(root.Right, value)
		return root, existed
	}
	existed = true
	if root.Left == nil && root.Right == nil {
		root = nil
		return root, existed
	}
	if root.Left == nil {
		root = root.Right
		return root, existed
	}
	if root.Right == nil {
		root = root.Left
		return root, existed
	}
	smallestInRight, _ := min(root.Right)
	root.Value = smallestInRight
	root.Right, _ = remove(root.Right, smallestInRight)
	return root, existed
}

func min(node *Node) (int, bool) {
	if node == nil {
		return 0, false
	}
	n := node
	for {
		if n.Left == nil {
			return n.Value, true
		}
		n = n.Left
	}
}

func max(node *Node) (int, bool) {
	if node == nil {
		return 0, false
	}
	n := node
	for {
		if n.Right == nil {
			return n.Value, true
		}
		n = n.Right
	}
}

func search(node *Node, value int) (*Node, bool) {
	if node == nil {
		return nil, false
	}
	if value < node.Value {
		return search(node.Left, value)
	}
	if value > node.Value {
		return search(node.Right, value)
	}
	return node, true
}
func preOrderTraverse(n *Node, f func(int)) {
	if n != nil {
		f(n.Value)
		preOrderTraverse(n.Left, f)
		preOrderTraverse(n.Right, f)
	}
}
func postOrderTraverse(n *Node, f func(int)) {
	if n != nil {
		postOrderTraverse(n.Left, f)
		postOrderTraverse(n.Right, f)
		f(n.Value)
	}
}
func inOrderTraverse(n *Node, f func(int)) {
	if n != nil {
		inOrderTraverse(n.Left, f)
		f(n.Value)
		inOrderTraverse(n.Right, f)
	}
}

func getNodeNum(node *Node) int {
	if node == nil {
		return 0
	} else {
		return getNodeNum(node.Left) + getNodeNum(node.Right) + 1
	}
}

func getDegree(node *Node) int {
	if node == nil {
		return 0
	}
	maxDegree := 0
	leftDegree := getDegree(node.Left)
	rightDegree := getDegree(node.Right)
	if leftDegree > rightDegree {
		maxDegree = leftDegree
	} else {
		maxDegree = rightDegree
	}
	return maxDegree + 1
}

func getNthNodeNum(node *Node, nth int) int {
	if node == nil {
		return 0
	}
	if nth == 1 {
		return 1
	}
	return getNthNodeNum(node.Left, nth-1) + getNthNodeNum(node.Right, nth-1)
}

func getLeavesNum(node *Node) int {
	if node == nil {
		return 0
	}
	if node.Left == nil && node.Right == nil {
		return 1
	}
	return getLeavesNum(node.Left) + getLeavesNum(node.Right)
}

func isBalance(node *Node) bool {
	if node == nil {
		return true
	}
	leftDegree := getDegree(node.Left)
	rightDegree := getDegree(node.Right)
	flag := false
	if (math.Abs(float64(leftDegree - rightDegree))) <= 1 {
		flag = true
	}
	return flag && isBalance(node.Left) && isBalance(node.Right)
}
