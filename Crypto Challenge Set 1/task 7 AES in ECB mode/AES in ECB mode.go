package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	//"crypto/aes"
)

func main() {
	// key := "YELLOW SUBMARINE"
	_file, err := ioutil.ReadFile("text.txt")
	if err != nil {
		fmt.Println("Error ", err)
	}
	file := strings.Split(string(_file), "\n")

	fmt.Println(file)
}
