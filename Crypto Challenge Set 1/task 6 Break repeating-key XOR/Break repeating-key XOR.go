package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"math/bits"
)

func decode_base64(str string) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		fmt.Println("Error decode base64: ", err)
	}

	return data, err
}

func hamming_distance() int {
	str_1 := "this is a test"
	str_2 := "wokka wokka!!!"

	if len(str_1) != len(str_2) {
		fmt.Println("Different length!")
	}

	var count int

	for i := 0; i < len(str_1); i++ {
		count = count + bits.OnesCount8(str_1[i]^str_2[i])
	}

	return count
}

func main() {
	file, err := ioutil.ReadFile("text.txt")

	if err != nil {
		fmt.Println("Error file open!")
	}

	dec, _ := decode_base64(string(file))
	fmt.Println(string(dec))
	//fmt.Println(file)

}
