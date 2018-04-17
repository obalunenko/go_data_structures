package binary_tree

import "fmt"

// A TreeNode is a binary tree with integer values.

type TreeNode struct {
	LeftNode  *TreeNode
	Value     int
	RightNode *TreeNode
}

func (t *TreeNode) Insert(data int) *TreeNode {

	if t == nil {
		t.Value = data
		t.LeftNode = nil
		t.RightNode = nil
		return t

	}

	switch {
	case data > t.Value:
		t.RightNode.Insert(data)
	case data < t.Value:
		t.LeftNode.Insert(data)
	default:
		fmt.Printf("Already exist")
		return t

	}

	return t

}

func (t *TreeNode) Delete(data int) *TreeNode {
	return t
}

func (t *TreeNode) Search(data int) *TreeNode {

	return t
}

func (t *TreeNode) Print() {
	if t != nil {

	}

}
