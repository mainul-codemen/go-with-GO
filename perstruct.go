package main

import "fmt"

type Person struct {
	Name string
}

type Android struct {
	Person
	Model string
}

func main() {
	fmt.Print()
	p := Person{
		Name: "Mainul",
	}
	fmt.Println(p.talk())

	a := Android{
		Person: p,
		Model:  "the",
	}
	s := Server{
		an: a,
	}

	fmt.Println(s.model())
}

type Server struct {
	an Android
}

func (s Server) model() string {

	return s.model() + s.an.talk() + s.an.Name

}

func (p Person) talk() string {
	return p.Name
}
