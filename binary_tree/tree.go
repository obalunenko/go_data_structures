package binaryTree

import (
	"fmt"
	"io"

	log "github.com/sirupsen/logrus"
)

// TreeNode is a binary tree with integer values.
type TreeNode struct {
	LeftNode  *TreeNode
	Value     int
	RightNode *TreeNode
}

// Insert fundtion adds value to tree
func (t *TreeNode) Insert(data int) *TreeNode {

	if t == nil {

		return &TreeNode{
			LeftNode:  nil,
			Value:     data,
			RightNode: nil,
		}
	} else if data > t.Value {

		t.RightNode = t.RightNode.Insert(data)

	} else if data < t.Value {

		t.LeftNode = t.LeftNode.Insert(data)
	}

	return t
}

// Delete function removes value from tree
func (t *TreeNode) Delete(data int) *TreeNode {
	return t
}

// Search function searches value in tree
func (t *TreeNode) Search(data int) *TreeNode {

	return t
}

// Print prints a visual representation of the tree
func (t *TreeNode) Print(out io.Writer) {

	_, err := fmt.Fprintln(out, "------------------------------------------------")
	if err != nil {
		log.Fatalf("could not leave with errors in print: %v", err)
	}
	stringify(t, out, 0, "")
	_, err = fmt.Fprintln(out, "------------------------------------------------")
	if err != nil {
		log.Fatalf("could not leave with errors in print: %v", err)
	}
}

// internal recursive function to print a tree
func stringify(t *TreeNode, out io.Writer, level int, symbol string) {

	if t != nil {
		format := ""
		for i := 0; i < level; i++ {

			format += "       "

		}
		if level == 0 {
			format += symbol + "───" + "───[ "

		} else {
			format += symbol + "───[ "
		}
		level++
		stringify(t.RightNode, out, level, "┌───")

		_, err := fmt.Fprintf(out, format+"%d"+"\n", t.Value)
		if err != nil {
			log.Fatalf("could not leave with errors in print: %v", err)
		}

		stringify(t.LeftNode, out, level, "└───")

	}

}

// InitTree function creates tree from array of int values
func InitTree(values ...int) *TreeNode {

	if len(values) > 0 {
		var rootNode = &TreeNode{
			LeftNode:  nil,
			Value:     values[0],
			RightNode: nil,
		}
		for i, value := range values {
			if i > 0 {
				rootNode = rootNode.Insert(value)
				//rootNode.Print(os.Stdout)
			}
		}
		return rootNode

	}
	return nil

}
