package main

import (
	"fmt"
)

func main(){
	var a [5]int

	a[4] = 15
	a[3] = 15
	a[2] = 15
	a[1] = 15
	fmt.Println(a)
	var total int
	for i := 0; i < len(a); i++{
		total += a[i]
	}
	fmt.Println(total)

	for i,value := range a{
		total += value
		fmt.Println(i)
	}
	fmt.Println(total)


	//x := [5]int{7,8,9,1,2}
	//var x =  [5]int{7,8,9,1,2}
	var x [5]int=  [5]int{7,8,9,1,2}


	fmt.Println(x)


	xy := make([]float32,5)

	xyz := make([]float32,7,10)
	fmt.Println(xy)
	fmt.Println(xyz)

	slice1 := []int{1,2,3}
	slice2 := append(slice1,4,5)


	fmt.Println(slice1, slice2)

	d := make(map[string]int)

	d["Key"] = 100
	d["Keyd"] = 1005

	fmt.Println(d)

	fmt.Println(x)

	

}