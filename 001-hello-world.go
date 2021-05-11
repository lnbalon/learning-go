package main

import (
	"fmt"
)

var temp int16 =42


func main() {
	fmt.Printf("%v %T\n", temp, temp)
	var temp int = 12
	var bong float32 
	bong = 11
	fmt.Printf("%v %T\n", temp, temp)
	fmt.Printf("%v %T\n", bong, bong)
}