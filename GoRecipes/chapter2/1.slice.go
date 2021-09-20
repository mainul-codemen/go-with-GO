package main

import "fmt"

func main() {

	z := make([]int, 5)
	z[0] = 1
	z[1] = 20

	r := []int{}
	fmt.Println(r)
	res := retslice(z)
	r = res

	fmt.Println(r)
}

func retslice(sl []int) []int {
	return sl
}
