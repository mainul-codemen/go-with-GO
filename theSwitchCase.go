package main

import (
	"fmt"
)

func main()  {
	var i = 10
	switch i {
	case 1:fmt.Println("one")
	case 10:fmt.Println("Ten")
	b := 3
	length := 5
	appendZero(b, length)
	 ss := reverse("anik")
	 fmt.Println(ss)
	}
}

//func appendZero(b,length int)  {
//	var str string
//	for i := 0; i < length - b; i++{
//		str += "0"
//	}
//	str = str+"10012"
//	fmt.Println(str)
//}
