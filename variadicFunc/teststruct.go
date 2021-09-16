package main

import "fmt"

func main() {

	type User struct {
		Name  string
		Age   int
		Email string
	}

	user := User{
		Name:  "mainul",
		Age:   550,
		Email: "mainul@codemen.net",
	}
	fmt.Println(user)
	u := new(User)

	u.Name = "test"
	u.Age = 58
	u.Email = "test@gmail.com"

	fmt.Printf("%v", u)

	uu := &User{
		Name:  "tr",
		Age:   50,
		Email: "re",
	}
	fmt.Printf("%v", uu)
	greeting := "Hello, world!"
	for i := 0; i < len(greeting); i++ {
		fmt.Println(greeting[i])
	}
}
