package main

import "fmt"

func main() {

	var score = make(map[string]int)

	score["Hitman"] = 1
	score["Pandya"] = 10
	score["sakib"] = 20

	fmt.Print(score)

	for k, v := range score {
		fmt.Println(k, v)
	}
}
