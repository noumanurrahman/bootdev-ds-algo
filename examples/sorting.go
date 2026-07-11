package examples

import (
	"fmt"
	"sorting-algorithms/algorithms"
)

func Sorting() {
	nums := []int{5, 2, 9, 4, 6, 3}
	algorithms.BubbleSort(nums)
	fmt.Println("Bubble Sort:", nums)
	algorithms.InsertionSort(nums)
	fmt.Println("Insertion Sort:", nums)
	algorithms.SelectionSort(nums)
	fmt.Println("Selection Sort:", nums)
	sortedNums := algorithms.MergeSort(nums)
	fmt.Println("Merge Sort:", sortedNums)
	algorithms.QuickSort(nums, 0, len(nums)-1)
	fmt.Println("Quick Sort:", nums)
}
