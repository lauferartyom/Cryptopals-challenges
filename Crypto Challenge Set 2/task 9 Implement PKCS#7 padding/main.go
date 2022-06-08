package main

import (
	"fmt"
)

func add_byte(buf []byte, size int) string {
	var buffer string

	for i := 0; i < len(buf); i++ {
		buffer += string(buf[i])
	}

	for i := len(buf) - 1; i < size-1; i++ {
		buffer += "\x04"
	}

	return buffer
}

func main() {
	str := "YELLOW SUBMARINE"
	fmt.Println(add_byte([]byte(str), 20))
	//str_1 := "YELLOW SUBMARINE♦♦♦♦"
}
