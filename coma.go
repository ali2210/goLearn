package main

//non-recursive version of comma using bytes.Buffer
import (
	"fmt"
	"bytes"
)

func intToString(str []int) string{
	 var buf bytes.Buffer
	 for i,v := range str {		
		if i > 0{
			buf.WriteString(",")
		 }
		fmt.Fprintf(&buf,"%d",v)
	 }
	return buf.String()
}

func main(){
	str := []int{1,2,3,4}
	fmt.Printf("%s",intToString(str))
}
