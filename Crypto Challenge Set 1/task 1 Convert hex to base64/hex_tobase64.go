package main

import (
    "fmt"
    "encoding/hex"
    "encoding/base64"
)

func hex_to_string(str string) string{
    st, err := hex.DecodeString(str)
    if err != nil{
        panic(err)
    }
    return string(st)
}

func string_to_base64(str string) string{
    st := base64.StdEncoding.EncodeToString([]byte(str))
    return st
}


func main() {
    str := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
    a := hex_to_string(str)
    fmt.Println(string_to_base64(a))
}