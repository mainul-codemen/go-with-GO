package main

import "fmt"

func main() {
	z := new(int)

	three(z)

	fmt.Println(z)
}

func three(z *int) {
	*z = 3
}
