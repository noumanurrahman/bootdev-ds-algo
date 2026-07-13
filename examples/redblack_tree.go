package examples

import (
	"sorting-algorithms/structures"
)

func RedBlackTree() {
	tree := structures.RBTree{}
	tree.New()
	for i := range 10 {
		tree.Insert(i)
	}
	tree.Display()
}
