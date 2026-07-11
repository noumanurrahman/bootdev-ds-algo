package examples

import (
	"fmt"
	"sorting-algorithms/structures"
)

func Stacks() {
	parenthesis := "()()"
	isBalanced := structures.IsBalanced(parenthesis)
	if isBalanced {
		fmt.Printf("'%s' is balanced\n", parenthesis)
	} else {
		fmt.Printf("'%s' is not balanced\n", parenthesis)
	}
}
