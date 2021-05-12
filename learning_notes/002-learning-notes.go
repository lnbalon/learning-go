package main

import (
	"fmt"
	"math"
	"math/rand"
)

func add(x, y int) int {
	return x + y
}

func swap(x,y string) (string, string) {
	return y, x
}


func split(i int) (j, k int) {
	j = i + 10
	k = i - 10
	return
}


func main() {
	fmt.Println(rand.Intn(100))
	fmt.Println(math.Sqrt(7))
	fmt.Println(math.Pi)
	fmt.Println(add(23,21))
	fmt.Println(swap("string", "swap"))
	fmt.Println(split(100))
}
