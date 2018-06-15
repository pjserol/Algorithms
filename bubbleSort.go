package main

func BubbleSort(listUnsorted []int) ([]int, int) {
	
	numberOfIteration := 0
	swapped := true
	list := make([]int, len(listUnsorted))

	for i := range listUnsorted {
		list[i] = listUnsorted[i]
	}

	for swapped {
		swapped = false
		for i := 1; i < len(list); i++ {
			if (list[i-1] > list[i]) {
				list[i-1], list[i] = list[i], list[i-1]
				swapped = true
			}
			numberOfIteration++
		}
	}

	return list, numberOfIteration
}