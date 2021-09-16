package main

import (
	"fmt"
	"sort"
)

type Developer struct {
	Name  string
	Phone string
	Age   int
}

type SortByName []Developer

func (sbn SortByName) Len() int {
	return len(sbn)
}

func (sbn SortByName) Less(i, j int) bool {
	return sbn[i].Name < sbn[j].Name
}

func (sbn SortByName) Swap(i, j int) {
	sbn[i], sbn[j] = sbn[j], sbn[i]
}

type SortByAge []Developer

func (sba SortByAge) Len() int {
	return len(sba)
}

func (sba SortByAge) Less(i, j int) bool {
	return sba[i].Age > sba[j].Age
}

func (sba SortByAge) Swap(i, j int) {
	sba[i], sba[j] = sba[j], sba[i]
}

type SortByPhone []Developer

func (sbp SortByPhone) Len() int {
	return len(sbp)
}

func (sbp SortByPhone) Less(i, j int) bool {
	return sbp[i].Phone < sbp[j].Phone
}

func (sbp SortByPhone) Swap(i, j int) {
	sbp[i], sbp[j] = sbp[j], sbp[i]
}

func main() {

	emp := []Developer{
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
		{
			Name:  "Sajib",
			Phone: "01575641236",
			Age:   50,
		},
		{
			Name:  "Halim",
			Phone: "01675641236",
			Age:   110,
		},
		{
			Name:  "Bob",
			Phone: "01375641236",
			Age:   100,
		},
	}

	emp2 := emp

	fmt.Println(emp)
	fmt.Println("  === Sorted Data By Name for struct ====   ")
	sort.Sort(SortByName(emp))
	fmt.Println(emp)

	fmt.Println("<-------------------Unsorted Sort By Age----------------------->")

	fmt.Println(emp2)
	fmt.Println("<-------------------Sort By Age----------------------->")
	sort.Sort(SortByAge(emp2))
	fmt.Println(emp2)

	fmt.Println("<-------------------Unsorted Sort By Phone----------------------->")
	fmt.Println(emp)
	fmt.Println("<-------------------Sort By Age----------------------->")
	sort.Sort(SortByPhone(emp2))
	fmt.Println(emp)

}
