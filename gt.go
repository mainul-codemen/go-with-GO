package main

import (
	"fmt"
	"strconv"
)

func main()  {
	result := addBinary("101","101")
	fmt.Println(result)

}
func addBinary(a string, b string) string {
	if len(a) == 0 || a == "" {
		return b
	}
	if len(b) == 0 || b == ""{
		return a
	}

	if len(b) > len(a) {
		temp := a
		a = b
		b = temp
	}
	newB := appendZero(b,len(a))

	var result string
	var carry int = 0
	var sum int = 0
	var str1  =[]rune(a)
	var str2  = []rune(newB)
	for i := len(str1)-1; i>=0; i-- {
		num1,_ := strconv.Atoi(string(str1[i]))
		num2,_ := strconv.Atoi(string(str2[i]))
		currentSum := num1 + num2 + carry

		if currentSum == 3 {
			carry = 1
			sum =1
		}else if currentSum == 2 {
			carry = 1
			sum = 0
		}else{
			carry = 0
			sum = currentSum
		}
		result +="0"
		fmt.Println(result)

		result += string(sum)
	}

	if carry !=0 {
		result += string(carry)
	}
	fmt.Println(result)
	return reverse(result)

}


func reverse(result string) string{
	r := []rune(result)
	var res []rune

	for i := len(r) -1; i >= 0; i--{
		res = append(res,r[i])
	}
	return string(res)

}


func appendZero(b string, length int) string{
	var str string
	for i :=0; i < length - len(b); i++{
		str+= "0"
	}
	str=str+b
	return str
}