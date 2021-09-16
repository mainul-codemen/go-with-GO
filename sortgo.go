package main

import (
	"fmt"
	"sort"
)

func main() {
	data := make([]int, 5, 10)
	data[0] = 15
	data[1] = 20
	data[2] = 80
	data[3] = 3
	data[4] = 49
	data = append(data, 100)
	fmt.Println(data)
	fmt.Println("Sorted Data")
	sort.Ints(data)
	fmt.Println(data)
}
