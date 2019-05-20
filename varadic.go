package main

import (
	"fmt"
)

func main() {
	values := []int{1,2,3}
	fmt.Println("Value[0]",seriesAdd(values[0]))
	fmt.Println("Total", seriesAdd(values...))
	fmt.Println("Total", seriesAdd(values))
}

func seriesAdd(x ...int)int{
	var count int = 0
	for _, v := range x{
		count+= v
	}
	return count
}
