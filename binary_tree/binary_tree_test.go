package binarytree

import (
	"reflect"
	"testing"
)

func TestDfSearchAndInsert(t *testing.T) {
	type test struct {
		inputs   []int
		expected []int
	}

	tests := []test{
		{
			[]int{2, 1, 3},
			[]int{2, 1, 3},
		},
		{
			[]int{2},
			[]int{2},
		},
		{
			[]int{5, 10, 1, 7, 3, 7, 6},
			[]int{5, 1, 3, 10, 7, 6},
		},
	}

	for _, test := range tests {
		firstNode := createBTreeAndReturnFirstNode(test.inputs)
		gotVals := getValsFromNodes(firstNode.DfSearch())
		wantVals := test.expected
		if !reflect.DeepEqual(gotVals, wantVals) {
			t.Errorf("\ngot: %v\nwant: %v", gotVals, wantVals)
		}
	}
}

func TestDeleteNode(t *testing.T) {
	input := []int{5, 10, 1, 7, 3, 7, 6}
	firstNode := createBTreeAndReturnFirstNode(input)

	firstNode = DeleteNode(firstNode, 1)
	gotVals := getValsFromNodes(firstNode.DfSearch())
	wantVals := []int{5, 3, 10, 7, 6}
	if !reflect.DeepEqual(gotVals, wantVals) {
		t.Errorf("\ngot: %v\nwant: %v", gotVals, wantVals)
	}
	// first test - remove a node from the list that exists

	firstNode = DeleteNode(firstNode, 7)
	gotVals = getValsFromNodes(firstNode.DfSearch())
	wantVals = []int{5, 3, 10, 6}
	if !reflect.DeepEqual(gotVals, wantVals) {
		t.Errorf("\ngot: %v\nwant: %v", gotVals, wantVals)
	}
	// second test - remove another mode from the list that exists

	firstNode = DeleteNode(firstNode, 7)
	gotVals = getValsFromNodes(firstNode.DfSearch())
	wantVals = []int{5, 3, 10, 6}
	if !reflect.DeepEqual(gotVals, wantVals) {
		t.Errorf("\ngot: %v\nwant: %v", gotVals, wantVals)
	}
	// last test - remove node from list that doesn't exist
	// resulting list should still be the same
}
func TestDeleteRootNode(t *testing.T) {
	input := []int{5, 10, 1, 7, 3, 7, 6}
	firstNode := createBTreeAndReturnFirstNode(input)

	firstNode = DeleteNode(firstNode, 5)
	gotVals := getValsFromNodes(firstNode.DfSearch())
	wantVals := []int{6, 1, 3, 10, 7}
	if !reflect.DeepEqual(gotVals, wantVals) {
		t.Errorf("\ngot: %v\nwant: %v", gotVals, wantVals)
	}
}

func TestGetMinAndMaxNode(t *testing.T) {

	input := []int{5, 10, 1, 7, 3, 7, 6}
	firstNode := createBTreeAndReturnFirstNode(input)

	gotMinNode := firstNode.GetMinNode()
	wantVal := 1
	if gotMinNode.GetVal() != wantVal {
		t.Errorf("got %d\n want %d\nnode: %v",
			gotMinNode.val,
			wantVal,
			gotMinNode)
	}
	gotMaxNode := firstNode.GetMaxNode()
	wantVal = 10
	if gotMaxNode.GetVal() != wantVal {
		t.Errorf("got %d\n want %d\nnode: %v",
			gotMaxNode.val,
			wantVal,
			gotMaxNode)
	}
}

func createBTreeAndReturnFirstNode(input []int) *BSTNode {
	firstVal := input[0]
	otherVals := input[1:]

	firstNode := InitBSTNode(firstVal)
	for _, val := range otherVals {
		InsertNode(firstNode, val)
	}
	return firstNode

}
