package main

import (
	"fmt"
)

func main() {
	//slices to maps 
	
	person := make(map[string]int)
	person["ali"] = 24
	person["rahma"] =23
	
	if details , ok := person["rahma"]; ok{
		fmt.Println(ok)
	}
	
}



