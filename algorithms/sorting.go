package algorithms

import "fmt"

// BubbleSort is the fucking best
func BubbleSort(x []int) []int {
	fmt.Println(len(x))
	end := len(x)

	for end > 0 {
		for i := 0; i<end-1; i++ {
			if x[i] > x[i+1] {
				x[i+1], x[i] = x[i], x[i+1]
			}
		}
		end--
	}
	return x
}