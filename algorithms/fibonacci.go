package algorithms

// GetFibonnaci returns a string of Fibonnaci
func GetFibonnaci(x int) []int {
	
	var b int
	var a []int
	a = make([]int, x)
	a[0] = 1
	a[1] = 1
	
	for y:= 2; y<x; y++ {
		b = a[y-1] + a[y-2]
		a[y] = b
	}
	return a
}