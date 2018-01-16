package main

import "fmt"

func main() {
	fmt.Println("hi")
	sampleArr := [10]int{10, 2, 4, 5, 67, 1, 4, 0, 3, 1}
	fmt.Println("Sample array original sample", sampleArr)
	sampleArrCall(sampleArr[:])
	fmt.Println("Sample array edited sample", sampleArr)

}

func sampleArrCall(array []int) {

	for i := 0; i < len(array)/2; i++ {
		temp := array[i]
		array[i] = array[len(array)-1-i]
		array[len(array)-1-i] = temp
	}
	fmt.Println("Sample array edited sample from within the sample:\n", array)
}
