package main

func insertionSort(listUnsorted []int) ([]int, int) {

	numberOfIteration := 0
	list := make([]int, len(listUnsorted))

	for i := range listUnsorted {
		list[i] = listUnsorted[i]
	}

	for i := 1; i < len(list); i++ {
		j := i

		for j >= 0 && list[j-1] > list[j] {
			list[j], list[j-1] = list[j-1], list[j]
			j = j - 1

			numberOfIteration++
		}
	}

	return listUnsorted, numberOfIteration
}
