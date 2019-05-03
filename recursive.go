package main

import (
	"fmt"
)

func main() {
	i := fib(5)
	fmt.Println(i)
		
		
	
}

func fib(n int)(int) {
	if n == 1{
		return n 
	}
	return n* fib(n-1)
	
	
}
