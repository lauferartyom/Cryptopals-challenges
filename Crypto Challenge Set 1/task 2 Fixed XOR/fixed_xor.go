package main

import (
	"encoding/hex"
	"fmt"
)

func Fixed_XOR(str1 string, str2 string) string {
	str_one, err_one := hex.DecodeString(str1)
	if err_one != nil {
		fmt.Printf("Error one: %s", err_one)
	}
	str_two, err_two := hex.DecodeString(str2)
	if err_two != nil {
		fmt.Printf("Error two: %s", err_two)
	}

	var r string

	for i := 0; i < len(str_one); i++ {
		r = r + string(str_one[i]^str_two[i])
	}
	return hex.EncodeToString([]byte(r))
}

func main() {
	str_1 := "1c0111001f010100061a024b53535009181c"
	str_2 := "686974207468652062756c6c277320657965"

	fmt.Println(Fixed_XOR(str_1, str_2))
}
