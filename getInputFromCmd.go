package main

import (
	"fmt"
	"os"
)

// A slice can store multiple values
func main()  {
	var name string
	name = os.Args[1]
	name,age := "fadfa",2019
	fmt.Println("Hello greet",name,"!")
	fmt.Println("My name is ",name)
	fmt.Println("My age is ",age)

}
