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
}

func displayResultSort(unsorted []int, sorted []int, numberOfIteration int) {
	fmt.Println("Before : ", unsorted)
	fmt.Println("After  : ", sorted) 
	fmt.Println("Number of iteration  : ", numberOfIteration) 
	fmt.Println()
}