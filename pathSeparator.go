package main

import (
	"fmt"
	"path"
)

func  main()  {

	var dir,file string
	dir,file = path.Split("go/src/css/main.css")
	fmt.Println(dir)
	fmt.Println(file)

	speed := 100
	force := 20.4
	speed = speed + int(force)
	fmt.Println(float32(speed))
}