package main

import "fmt"

func main() {
	//b := [12]int{23, 54, 2, 58, 19, -21, -4, 0, 7, 6, 10, -0}
	b := [5]int{5, 4, 3, 2, 1}
	fmt.Println("Sorted Array=", insertionSort(b[:]))
}

func insertionSort(array []int) []int {
	for i := 1; i < len(array); i++ {
		temp := array[i]
		j := i
		for j > 0 && array[j-1] > temp {
			array[j] = array[j-1]
			j--
		}
		array[j] = temp
		fmt.Println("Intermediate array:", array)
	}
	return array
}
