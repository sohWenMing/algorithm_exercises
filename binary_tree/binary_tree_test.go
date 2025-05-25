package binarytree

import (
	"reflect"
	"testing"
)

func TestDfSearch(t *testing.T) {
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
		firstVal := test.inputs[0]
		otherVals := test.inputs[1:]

		firstNode := InitBSTNode(firstVal)
		for _, val := range otherVals {
			firstNode.InsertNode(val)
		}
		gotVals := getValsFromNodes(firstNode.DfSearch())
		wantVals := test.expected
		if !reflect.DeepEqual(gotVals, wantVals) {
			t.Errorf("\ngot: %v\nwant: %v", gotVals, wantVals)
		}

	}
}
