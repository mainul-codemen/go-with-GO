package main

import "fmt"

type Employee struct {
	Name  string
	Phone string
	Age   int
}

func main() {

	emp := []Employee{
		{
			Name:  "Mainul",
			Phone: "01742503080",
			Age:   200,
		},
		{
			Name:  "Alamin",
			Phone: "01987520315",
			Age:   450,
		},
		{
			Name:  "Tree",
			Phone: "01875641236",
			Age:   100,
		},
	}

	fmt.Println(emp)

	for i, v := range emp {
		fmt.Println("Index Number = ", i, v)
	}

}
