package main

import (
	"fmt"
)
type Weekend int 


const (
	Sunday Weekend = iota
	Monday 
	Tuesday
	Wednesday
	Thrusday
	Friday
	Saturday
)

func main() {
	fmt.Printf("%d",Saturday)
}
