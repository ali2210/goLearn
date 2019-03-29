package main

import (
	"fmt"
	
)

func main() {
	r := []byte("13579")
	fmt.Println(string(reverse(r)))
	
}

func reverse(b []byte) []byte{
	var j int = len(b)-1
	var i int = 0
		if b[i] != b[j]{
			b[i],b[j]= b[j],b[i]
		}
		return b

}
