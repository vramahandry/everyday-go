package main

import (
	"fmt"

	"github.com/vramahandry/everyday-go/chapter6/sum-testing/internal/myMath"
)

func main() {
	x := 5
	y := 5
	fmt.Printf("%d + %d = %d\n", x, y, myMath.Sum(x, y))
}
