package binarytree_solution

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
	isInserted = false
	if n == nil {
		return InitBSTNode(inputVal), true
	}
	if inputVal == n.val {
		return n, false
	}
	// no duplicates of values in a binary tree
	if inputVal < n.val {
		if n.left == nil {
			newNode := InitBSTNode(inputVal)
			n.SetLeft(newNode)
			return n, true
		}
		_, isInserted := InsertNode(n.left, inputVal)
		return n, isInserted
	} else {
		if n.right == nil {
			newNode := InitBSTNode(inputVal)
			n.SetRight(newNode)
			return n, true
		}
		_, isInserted := InsertNode(n.right, inputVal)
		return n, isInserted
	}
}

func DeleteNode(n *BSTNode, inputVal int) (returnedNode *BSTNode) {
	if n == nil {
		return n
	}
	if inputVal < n.val {
		n.SetLeft(DeleteNode(n.left, inputVal))
		return n
	}
	if inputVal > n.val {
		n.SetRight(DeleteNode(n.right, inputVal))
		return n
	}

	// at this point, we should be at the node where the val matches
	if n.left == nil && n.right == nil {
		return nil
	}
	if n.left != nil && n.right == nil {
		return n.GetLeft()
	}
	if n.right != nil && n.left == nil {
		return n.GetRight()
	}
	successor := n.right.GetMinNode()
	n.SetVal(successor.GetVal())
	n.SetRight(DeleteNode(n.right, successor.val))
	return n
}

func (n *BSTNode) GetMaxNode() *BSTNode {
	if n == nil {
		return nil
	}
	currNode := n
	for currNode.right != nil {
		currNode = currNode.right
	}
	return currNode
}

func (n *BSTNode) GetMinNode() *BSTNode {
	if n == nil {
		return nil
	}
	currNode := n
	for currNode.left != nil {
		currNode = currNode.left
	}
	return currNode
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
	if n.left != nil {
		returnedFromLeft := n.left.DfSearch()
		visitedNodes = append(visitedNodes, returnedFromLeft...)
	}
	if n.right != nil {
		returnedFromRight := n.right.DfSearch()
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
func (n *BSTNode) SetVal(inputVal int) {
	n.val = inputVal
	return
}
