package main

import (
	"fmt"
	"go-with-GO/chapter8/math"
)

func main() {
	xs := []float64{10, 2, 3, 4, 5}
	avg := math.Average(xs)
	fmt.Println(avg)

}
