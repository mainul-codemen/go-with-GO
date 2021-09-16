package main

import "fmt"

func main() {

	res := add(1, 2, 3, 4, 6, 12)
	fmt.Println(res)
}

func add(val ...int) int {
	total := 0

	for _, v := range val {
		total += v
	}
	return total
}
