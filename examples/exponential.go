package examples

import (
	"fmt"
	"sorting-algorithms/algorithms"
)

func Exponential() {
	fibExponential := algorithms.FibonacciExponential(10)
	fmt.Println("Fibonacci Exponential:", fibExponential)
	fibPolynomial := algorithms.Fibonacci(100)
	fmt.Println("Fibonacci Polynomial:", fibPolynomial)

	digits := "6287"
	combinations := algorithms.LetterCombinations(digits)
	fmt.Println("Letters Combinations:", combinations)

	growth := algorithms.ExponentialGrowth(10, 2, 4)
	fmt.Println("Exponential Growth:", growth)
}
