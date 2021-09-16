package main

import (
	"fmt"
	"sort"
)

func main() {
	x := []int{1, 2, 3, 4, 5}
	sort.Ints(x)
	fmt.Println(x)

	s := []string{"Go", "Bravo", "Gopher", "Alpha", "Grin", "Delta"}
	sort.Strings(s)
	fmt.Println(s)

}
