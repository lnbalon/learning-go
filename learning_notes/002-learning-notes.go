package main

import "fmt"

func main() {
	sum := 0
	for i := 2; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
