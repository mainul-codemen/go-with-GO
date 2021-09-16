package main

import (
	"os"
)

func main() {
	// file, err := os.Open("tt.txt")
	// if err != nil {
	// 	return
	// }
	// defer file.Close()
	// stat, err := file.Stat()
	// fmt.Println("--------", stat)
	// if err != nil {
	// 	return
	// }
	// bs := make([]byte, stat.Size())
	// _, err = file.Read(bs)
	// if err != nil {
	// 	return
	// }
	// str := string(bs)
	// fmt.Println(str)

	filec, err := os.Create("mainul.txt")
	if err != nil {
		return
	}
	defer filec.Close()
	filec.WriteString("This is for test purspose")
}
