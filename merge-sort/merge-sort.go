package main

import "fmt"

func main() {
	sampleArray := [10]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	mergeSort(0, 9, sampleArray[:])
	fmt.Println(sampleArray)
}

func mergeSort(p int, r int, array []int) {
	fmt.Println("p = ", p, "r=", r, "array = ", array)
	if r >= p {
		q := (p + r) / 2
		mergeSort(p, q, array)
		mergeSort(q+1, r, array)
		merge(p, q, r, array)
	}
}

func merge(p int, q int, r int, array []int) {
	fmt.Println("p = ", p, "q=", q, "r=", r)
	var n1, n2, leftIndex, rightIndex int
	n1 = (p - q) / 2
	n2 = (q + 1 - r) / 2
	leftIndex = p
	rightIndex = q + 1
	L := make([]int, n1)
	R := make([]int, n2)
	for i := p; i <= q; i++ {
		L[i] = array[i]
	}
	for i := q + 1; q <= r; i++ {
		R[i] = array[i]
	}
	i := p
	for leftIndex <= q && rightIndex <= r {
		if L[i] <= R[i] {
			array[i] = L[i]
			leftIndex++
		} else {
			array[i] = R[i]
			rightIndex++
		}
		i++
	}
	for leftIndex <= q {
		array[leftIndex] = L[leftIndex]
		leftIndex++
	}
	for rightIndex <= r {
		array[rightIndex] = R[rightIndex]
		rightIndex++
	}
	fmt.Println("Intermediate Step:", array)
}
