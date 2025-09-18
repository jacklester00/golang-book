package main

import (
	"fmt"
	"golang-book/chapter8/math"
)

func main() {
	xs := []float64{1, 2, 3, 4}
	fmt.Printf("Average of %v is %v\n", xs, math.Average(xs))

	fmt.Printf("Max of %v is %v\n", xs, math.Max(xs))
	fmt.Printf("Min of %v is %v\n", xs, math.Min(xs))
}
