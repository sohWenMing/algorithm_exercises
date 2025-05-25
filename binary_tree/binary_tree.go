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

func (n *BSTNode) InsertNode(inputVal int) (isInserted bool) {

	isInserted = false
	if inputVal == n.GetVal() {
		return
	}
	// no duplicates of values in a binary tree
	switch inputVal < n.GetVal() {
	case true:
		if n.GetLeft() == nil {
			n.SetLeft(InitBSTNode(inputVal))
			isInserted = true
			return
		}
		isInserted = n.GetLeft().InsertNode(inputVal)
		return
	case false:
		if n.GetRight() == nil {
			n.SetRight(InitBSTNode(inputVal))
			isInserted = true
			return
		}
		isInserted = n.GetRight().InsertNode(inputVal)
		return
	}
	return
}

func getValsFromNodes(nodes []*BSTNode) []int {
	returnedVals := make([]int, 0, len(nodes))
	for _, node := range nodes {
		returnedVals = append(returnedVals, node.val)
	}
	return returnedVals
}

func (n *BSTNode) DfSearch() []*BSTNode {
	visitedNodes := []*BSTNode{}
	visitedNodes = append(visitedNodes, n)
	if n.GetLeft() != nil {
		returnedFromLeft := n.GetLeft().DfSearch()
		visitedNodes = append(visitedNodes, returnedFromLeft...)
	}
	if n.GetRight() != nil {
		returnedFromRight := n.GetRight().DfSearch()
		visitedNodes = append(visitedNodes, returnedFromRight...)
	}
	return visitedNodes
}

func (n *BSTNode) GetVal() int {
	return n.val
}

func (n *BSTNode) GetLeft() *BSTNode {
	return n.left
}

func (n *BSTNode) GetRight() *BSTNode {
	return n.right
}

func (n *BSTNode) SetLeft(inputNode *BSTNode) {
	n.left = inputNode
	return
}
func (n *BSTNode) SetRight(inputNode *BSTNode) {
	n.right = inputNode
	return
}
