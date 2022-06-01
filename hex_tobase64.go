package main 

import (
	"encoding/hex"
	"encoding/base64"
	"fmt"
)

func hex_to_base(str string) string{
	_str, err := hex.DecodeString(str)
	if err != nil{
		fmt.Printf("Error: %s", err)
	}
	str = string(base64.StdEncoding.EncodeToString(_str))
	return str
}


func main(){
	str := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	fmt.Println(hex_to_base(str))
}