package main

import "fmt"

func main()  {
	var str string = "mainul"
	res  :=reverse(str)
	fmt.Println(res)
}

//func reverse(result string) string{
//	r := []rune(result)
//	var res []rune
//
//	for i := len(r) -1; i >= 0; i--{
//		res = append(res,r[i])
//	}
//	return string(res)
//
//}
