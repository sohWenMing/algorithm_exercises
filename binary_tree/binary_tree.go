package binarytree

type BSTNode struct {
	val   int
	left  *BSTNode
	right *BSTNode
}

func InitBSTNode(inputVal int) *BSTNode {
	return &BSTNode{
		inputVal,
		nil,
		nil,
	}
}

func InsertNode(n *BSTNode, inputVal int) (newNode *BSTNode, isInserted bool) {
	return nil, false
}

func DeleteNode(n *BSTNode, inputVal int) (returnedNode *BSTNode) {
	return nil
}

func (n *BSTNode) GetMaxNode() *BSTNode {
	return nil
}

func (n *BSTNode) GetMinNode() *BSTNode {
	return nil
}
func getValsFromNodes(nodes []*BSTNode) []int {

	return nil
}

func (n *BSTNode) DfSearch() []*BSTNode {
	visitedNodes := []*BSTNode{}
	visitedNodes = append(visitedNodes, n)
	if n.left != nil {
		returnedFromLeft := n.left.DfSearch()
		visitedNodes = append(visitedNodes, returnedFromLeft...)
	}
	if n.right != nil {
		returnedFromRight := n.right.DfSearch()
		visitedNodes = append(visitedNodes, returnedFromRight...)
	}
	return visitedNodes
} // keep this, is used for testing

func (n *BSTNode) GetVal() int {
	return 0
}

func (n *BSTNode) GetLeft() *BSTNode {
	return nil
}

func (n *BSTNode) GetRight() *BSTNode {
	return nil
}

func (n *BSTNode) SetLeft(inputNode *BSTNode) {
	return
}
func (n *BSTNode) SetRight(inputNode *BSTNode) {
	return
}
func (n *BSTNode) SetVal(inputVal int) {
	return
}
