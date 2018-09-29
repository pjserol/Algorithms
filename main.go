package main

import "fmt"

func main() {
	fmt.Println("--------------------")
	fmt.Println("-----Algorithms-----")
	fmt.Println("--------------------")
	fmt.Println()

	unsorted := []int{4, 5, 1, 2, 3}

	fmt.Println("-----Bubble Sort-----")
	sortedBubbleSort, nbIterationBubbleSort := BubbleSort(unsorted)
	displayResultSort(unsorted, sortedBubbleSort, nbIterationBubbleSort)

	fmt.Println("-----Selection Sort-----")
	sortedSelectionSort, nbIterationSelectionSort := selectionSort(unsorted)
	displayResultSort(unsorted, sortedSelectionSort, nbIterationSelectionSort)

	fmt.Println("-----Insertion Sort-----")
	sortedInsertionSort, nbIterationInsertionSort := selectionSort(unsorted)
	displayResultSort(unsorted, sortedInsertionSort, nbIterationInsertionSort)

	listMerge := make([]int, len(unsorted))

	for i := range unsorted {
		listMerge[i] = unsorted[i]
	}

	fmt.Println("-----Merge Sort-----")
	mergeSort(listMerge, 0, len(listMerge)-1)
	displayResultSort(unsorted, listMerge, 0)
}

func displayResultSort(unsorted []int, sorted []int, numberOfIteration int) {
	fmt.Println("Before : ", unsorted)
	fmt.Println("After  : ", sorted)
	fmt.Println("Number of iteration  : ", numberOfIteration)
	fmt.Println()
}
