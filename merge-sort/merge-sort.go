package main

import "fmt"

func main() {
	//sampleArray := [4]int{4, 3, 2, 1}
	//sampleArray := [5]int{5, 4, 3, 2, 1}
	sampleArray := [10]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	//sampleArray := [2]int{2, 1}
	mergeSort(0, len(sampleArray)-1, sampleArray[:])
	fmt.Println(sampleArray)
}

func mergeSort(p int, r int, array []int) {
	fmt.Println("p = ", p, "r=", r, "array = ", array)
	if p < r {
		q := (p + r) / 2
		mergeSort(p, q, array)
		mergeSort(q+1, r, array)
		merge(p, q, r, array)
	}
}

func merge(p int, q int, r int, array []int) {
	fmt.Println("p = ", p, "q=", q, "r=", r)
	var n1, n2, leftIndex, rightIndex int
	n1 = q - p + 1
	n2 = r - q
	leftIndex = 0
	rightIndex = 0
	L := make([]int, n1)
	R := make([]int, n2)
	for i := 0; i < n1; i++ {
		L[i] = array[p+i]
	}
	for i := 0; i < n2; i++ {
		R[i] = array[q+1+i]
	}
	i := p
	for leftIndex < n1 && rightIndex < n2 {
		if L[leftIndex] <= R[rightIndex] {
			array[i] = L[leftIndex]
			leftIndex++
		} else {
			array[i] = R[rightIndex]
			rightIndex++
		}
		i++
	}
	for leftIndex < n1 {
		array[i] = L[leftIndex]
		leftIndex++
		i++
	}
	for rightIndex < n2 {
		array[i] = R[rightIndex]
		rightIndex++
		i++
	}
	fmt.Println("Intermediate Step:", array)
}
