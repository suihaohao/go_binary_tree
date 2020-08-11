package binary_tree

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type BST struct {
	Root *Node
}

func (bst *BST) Insert(value int) {
	newNode := &Node{
		Value: value,
	}
	if bst.Root == nil {
		bst.Root = newNode
	} else {
		insertNode(bst.Root, newNode)
	}
}

func (bst *BST) Remove(value int) bool {
	_, existed := remove(bst.Root, value)
	return existed
}

func (bst *BST) Min() (int, bool) {
	return min(bst.Root)
}

func (bst *BST) Max() (int, bool) {
	return max(bst.Root)
}

//判断是否从存在某个节点
func (bst *BST) Search(value int) bool {
	_, existed := search(bst.Root, value)
	return existed
}

//前序遍历
func (bst *BST) PreOrderTraverse(f func(int)) {
	preOrderTraverse(bst.Root, f)
}

//后序遍历
func (bst *BST) PostOrderTraverse(f func(int)) {
	postOrderTraverse(bst.Root, f)
}

//中序遍历
func (bst *BST) InOrderTraverse(f func(int)) {
	inOrderTraverse(bst.Root, f)
}

//获取节点个数
func (bst *BST) GetNodeNum() int {
	return getNodeNum(bst.Root)
}

//获取层数
func (bst *BST) GetDegree() int {
	return getDegree(bst.Root)
}

//获取k层节点个数
func (bst *BST) GetNthNodeNum(nth int) int {
	return getNthNodeNum(bst.Root, nth)
}

//求叶子节点个数
func (bst *BST) GetLeavesNum() int {
	return getLeavesNum(bst.Root)
}

//判断二叉树是否平衡
func (bst *BST) IsBalance() bool {
	return isBalance(bst.Root)
}
