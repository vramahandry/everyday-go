package main

import "fmt"

func sum(x int, y int) int {
	// This code has a bug that adds 1
	return x + y
}
func main() {
	x := 5
	y := 5
	fmt.Printf("%d + %d = %d\n", x, y, sum(x, y))
}
