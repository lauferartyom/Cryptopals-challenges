package main

import (
    "fmt"
    "encoding/hex"
)


func encrypt(str string, key string) string{
    var r []byte
    
    for i := 0; i < len(str); i++{
        r = append(r, str[i] ^ key[i % 3])
    }
    
    return hex.EncodeToString(r)
}



func main() {
    str_1 := "Burning 'em, if you ain't quick and nimble I go crazy when I hear a cymbal"
    key := "ICE"
    a := encrypt(str_1, key)
    fmt.Println(a)
}