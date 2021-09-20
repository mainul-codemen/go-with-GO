package main

import "fmt"

func main() {
	fmt.Println("Starting to panic")
	panicRecover()
	fmt.Println("Program regains control after the panic revovery...")
}

func panicRecover() {
	defer fmt.Println("defer call -1")
	defer func() {
		fmt.Println("Defer call -2")
		if e := recover(); e != nil {
			fmt.Println("Recover with:", e)
		}
	}()
	defer fmt.Println("Defer -3")

	panic("Just panicking for the sake of example")
	fmt.Println("This will never be calledq")
}
