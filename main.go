package main

import (
	"fmt"
	"learning-go/algorithms"
)

func main(){

	// fmt.Println(algorithms.GetFibonnaci((20)))
	array := []int{11,2,3,4,45}
	sorted := algorithms.BubbleSort(array)
	fmt.Println(sorted)
}