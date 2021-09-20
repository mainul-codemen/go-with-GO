package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	s, _ := ReadFile("README.md")
	fmt.Println(string(s))
}

func ReadFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ioutil.ReadAll(f)
}
