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

func main() {
	fmt.Println(rand.Intn(100))
	fmt.Println(math.Sqrt(7))
	fmt.Println(math.Pi)
	fmt.Println(add(23,21))
	fmt.Println(swap("string", "swap"))
}
