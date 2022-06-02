package main

import (
    "encoding/hex"
    "fmt"
)


func freq(str string) float32{
    _freq := map[byte]float32{
        'a': 8.2389258,    'b': 1.5051398,    'c': 2.8065007,    'd': 4.2904556,
        'e': 12.813865,    'f': 2.2476217,    'g': 2.0327458,    'h': 6.1476691,
        'i': 6.1476691,    'j': 0.1543474,    'k': 0.7787989,    'l': 4.0604477,
        'm': 2.4271893,    'n': 6.8084376,    'o': 7.5731132,    'p': 1.9459884,
        'q': 0.0958366,    'r': 6.0397268,    's': 6.3827211,    't': 9.1357551,
        'u': 2.7822893,    'v': 0.9866131,    'w': 2.3807842,    'x': 0.1513210,
        'y': 1.9913847,    'z': 0.0746517, ' ' : 13,
    }
    
    var f float32
    
    for i := 0; i < len(str); i++{
        f += _freq[str[i]]
    }
    
    return f
}



func best_freq(buffer []float32) byte{
    var tmp float32 = 0
    var index byte = 0
    for i := 0; i < 256; i++{
        if buffer[i] >= tmp{
            tmp = buffer[i]
            index = byte(i)            
        }
    }
    return index
}


func crack(index byte, str []byte) string{
    var r string
    
    for i := 0; i < len(str); i++{
        r += string(str[i] ^ index)
    }
    return r
}



func decode(str string) string{
    s, _ := hex.DecodeString(str)
    var r string
    
    fmt.Println(len(s))
    var buffer []float32
    for i := 0; i < 256; i++{
        for j := 0; j < len(s); j++{
            r += string(s[j] ^ byte(i))
        }
        buffer = append(buffer, freq(r))
        r = ""
    }
    
    index := best_freq(buffer)
    
    return crack(index, s)
}

func main() {
    str := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
    fmt.Println(decode(str))
    //fmt.Println(crack(s, ))
}