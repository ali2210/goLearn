package main

import (
	"fmt"
)

	type Op struct{
		x int
	}
	
	func(o *Op)Add(y int)int{
		return o.x +y
	}

func main() {
	op := Op{4}  // method value 
	res := op.Add
	fmt.Println(res(4))
	
	oprend := &Op{5}   // method value unsual way
	add := (*oprend).Add
	fmt.Println(add(4))
	
}
