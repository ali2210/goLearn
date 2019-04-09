package main

import (
	"fmt"
)

func main() {
	//slices 
	names := [...]string{"ali","rahma","johny"}
	name := make([]string, 0, len(names))
	for _ , v := range names{
		name = append(name, v)
	 	fmt.Println(name)
	}
}

