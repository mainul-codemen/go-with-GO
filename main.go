package main

import "fmt"

type Textview struct {
	Name    string
	id      int
	Address string
}

func main() {

	text := []Textview{
		Textview{
			Name:    "Fakhrul",
			id:      15103234,
			Address: "Manikgong",
		},
		Textview{
			Name:    "Mainul",
			id:      15103205,
			Address: "khulna",
		},
		Textview{
			Name:    "Abdullah",
			id:      15103296,
			Address: "barishal",
		},
	}

	for i, v := range text {
		fmt.Println(i)
		fmt.Println(v)
	}

}
