package main

import (
	"fmt"
)

	type List struct{
		 data int
		 tail *List
	}
	var list List
	func(r *List)Insert()( *List) {
		if r == nil{
			panic ("Empty List")
		}
		return r
	}
		
	
	
func main() {
	var l *List = &list
	l.data = 3
	l.tail = nil
	p := l.Insert()
	fmt.Printf("%p", p)
	fmt.Println(p.data)
	h  := &list
	h.data = 1
	h.tail = p
	c := h.Insert()
	fmt.Println(c.data, c.tail)
}
