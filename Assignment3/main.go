package main

import (
	"fmt"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	inputFile := os.Args[1]
	file, err := os.Open(inputFile)
	check(err)

	//bs := make([]byte, 9999)
	//file.Read(bs)
	//fmt.Println(string(bs))
	io.Copy(os.Stdout, file)
	fmt.Println()
}
