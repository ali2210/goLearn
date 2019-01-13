package main

//non-recursive version [float] of comma using bytes.Buffer
import (
	"fmt"
	"bytes"
)

func intToString(str []float32) string{
	 var buf bytes.Buffer
	 for i,v := range str {		
		if i > 0{
			buf.WriteString(",")
		 }
		fmt.Fprintf(&buf,"%f",v)
	 }
	return buf.String()
}

func main(){
	str := []float32{1.0,2.0,3.3,14.9}
	fmt.Printf("%s",intToString(str))
}
