package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)




func main(){	
	const mix = "367 o"
	for _, v := range mix{
		if unicode.IsSpace(v) == true{
			encode := utf8.EncodeRune([]byte(mix),5)
			fmt.Println("encode:",encode)
			fmt.Printf("ascii :%q",byte (encode))
		}
		
	}
}
	
