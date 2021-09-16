package main

import (
	"fmt"
	add "go-with-GO/addition"
)
var dog = "My dog name is varible"
func main(){
	fmt.Println("Hello World")
	
	f := add.AdditionsTwoValue()
	fmt.Println(f)
	name := "Mainul"
	name = "hasan"
	age := 65
	fmt.Println(name,age)
	var (
		x string
		y string
	)

	x = "mainul"
	y = "hasan"

	fmt.Println(x,y)


	fmt.Println(dog)

	fi()

}

func fi(){

	var (
		a = 10
		b = 100

	)
	fmt.Println(dog)
	fmt.Println(a)
	fmt.Println(b)

	var input float64

	fmt.Scanf("%f", &input)

	out := input * 2

	fmt.Println(out)

}
