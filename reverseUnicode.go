package main

import (
	"fmt"
	"unicode/utf8"
)

/*reverse the number using unicode*/

func reverse(a []byte)[]byte {
    for i:= 0; i < len(a); i++{
	for j := i+1; j< len(a); j++{
		a[i],a[j] = a[j],a[i]
	}
     }
   return a
}

func main(){	
	s := 'ß'
	buf := make([]byte,3)	
	arr := utf8.EncodeRune(buf,s)
	fmt.Println("original:",buf)
	fmt.Println("reverse:",reverse(buf),"size:",arr)
}
