package main

import "fmt"

func main() {
	b := [12]int{23, 54, 2, 58, 19, -21, -4, 0, 7, 6, 10, -0}
	fmt.Println("Sorted Array=", selectionSort(b[:]))
}

func selectionSort(array []int) []int {

	for i := 0; i < len(array)-1; i++ {
		minIndex := i
		for j := i; j < len(array); j++ {
			if array[minIndex] > array[j] {
				minIndex = j
			}
		}
		temp := array[i]
		array[i] = array[minIndex]
		array[minIndex] = temp
		fmt.Println("Intermediate stage: ", array)
	}
	return array
}
