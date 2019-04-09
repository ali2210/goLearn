package main

import (
	"fmt"
)

func main() {
	//slices to maps 
	names := [...]string{"ali","rahma","johny"}
	person := make(map[string]int)
	
	for _, v := range names{
		person[v] = 24
		fmt.Println(person)
	}
}
