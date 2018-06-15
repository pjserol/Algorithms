package main

func selectionSort (listUnsorted []int) ([]int, int) {
	
	numberOfIteration := 0
	list := make([]int, len(listUnsorted))

	for i := range listUnsorted {
		list[i] = listUnsorted[i]
	}

	for i := 0; i<len(list) - 1; i++ {
		index := i
		for j := i + 1; j<len(list); j++ {
			if (list[j] < list[index]) {
				index = j
			}
			numberOfIteration++
		}
		list[i], list[index] = list[index], list[i]
	}

	return list, numberOfIteration
}