package main

import "fmt"

func main(){
	var a = []int{4,5,6,7,1,0,3}
	fmt.Println(a)

	aa := []string{"a","b","e","f"}

	a2 := append(aa,"e","q")
	fmt.Println(aa)
	fmt.Println(a2)

	a3 := make([]float64,3)
	a3[0] = 45.25
	a3[1] = 87.36
	a3[2] = 69.32
	a4 := append(a3,34.80,34.90 )
	fmt.Println(a3)
	fmt.Println(a4)
}