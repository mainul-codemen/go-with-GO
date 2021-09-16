package main

import (
	"fmt"
)

type Circle struct {
	x, y, z float64
}

func main() {
	c := Circle{
		x: 10,
		y: 50,
		z: 30,
	}
	r := circleArea(&c)
	fmt.Println(r)

	rs := c.area()

	fmt.Println(rs)

}

func circleArea(c *Circle) float64 {
	return c.x * c.y * c.z
}

func (c *Circle) area() float64 {
	return c.x * c.y * c.z
}
