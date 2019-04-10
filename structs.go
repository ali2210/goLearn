package main

import (
	"fmt"
)

type Students struct{
	ID  int
	Name  string
}

var record Students

func Id(i int) int{
	record.ID = i
	return record.ID
}	

func Name(name string)string{
	record.Name = name
	return record.Name
}

	
func main() {
	var id int = 0
	var name string = "ali"
	fmt.Println("id:", Id(id), "name:", Name(name))
}
