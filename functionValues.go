package main

import (
	"fmt"
)

func main() {
	f:= square
	fmt.Println(f(2))
	f = underRoot
	fmt.Println(f(4))
	f = add
	fmt.Println(f(2,2)) // error
	
}

func square(x int)int {
	return x * x
}

func underRoot(x int)int {
	return x*1/2
}

func add(x, y int)int {
	return x+y
}
