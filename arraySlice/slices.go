package main

import "fmt"

func main(){
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"

	fmt.Println(elements)

	name,ok := elements["H"]

	fmt.Println(name,ok)

	if name,ok := elements["He"];ok{
		fmt.Println(name)
	}

	 nn,ok := elements["Li"]
	fmt.Println(nn,ok)

	ele := map[string]map[string]string{
		"H": map[string]string{
			"name":"Hydrogen",
			"state":"gas",
		},
		"He": map[string]string{
			"name":"Helium",
			"state":"gas",
		},
		"Li": map[string]string{
			"name":"Lithium",
			"state":"solid",
		},
	}

	fmt.Println(ele)

	if el,ok := ele["Li"]; ok {
		fmt.Println(el["name"],el["state"])
	}

}