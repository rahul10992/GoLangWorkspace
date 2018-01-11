package main

import "fmt"

func main() {
	b := [12]int{23, 54, 2, 58, 19, -21, -4, 0, 7, 6, 10, -0}
	fmt.Println("Sorted Array=", bubbleSort(b[:]))
}

func bubbleSort(array []int) []int {

	for i := 0; i < len(array); i++ {
		for j := len(array) - 1; j >= i; j-- {
			if array[j] > array[j+1] {
				temp := array[j+1]
				array[j+1] = array[j]
				array[j] = temp
			}
		}
		fmt.Println("Intermediate stage: ", array)
	}
	return array
}
