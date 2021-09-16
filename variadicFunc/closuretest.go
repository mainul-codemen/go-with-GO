package main

import "fmt"

func main() {

	erew := func(val ...int) int {
		sum := 0
		for _, v := range val {
			sum += v
		}
		return sum
	}
	fmt.Println(erew(1, 2, 3, 4, 5))

	fmt.Println("Hello and hi")
}
