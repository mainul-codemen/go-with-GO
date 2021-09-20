package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	s, _ := readFile("READMEmd")
	fmt.Println(string(s))

}

func readFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	return ioutil.ReadAll(f)
}
