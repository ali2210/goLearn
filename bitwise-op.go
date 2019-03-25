package main

import (
	"fmt"
)

func main() {
	var x int8 = 1 << 2 | 1 << 4
	var y int8 = 1 << 3 | 1 << 6 
	fmt.Printf("x=%08b\n" , x)           
	fmt.Printf("y=%08b\n" , y)
	
	fmt.Printf("bit-wise AND=%08b\n" , x&y)
	fmt.Printf("bit-wise OR=%08b\n" , x|y)
	fmt.Printf("bit-wise XOR=%08b\n" , x^y)
	fmt.Printf("bit-wise bitClear=%08b\n" , x&^y)
	fmt.Printf("left shift=%08b\n" , x<<1)
	fmt.Printf("right shift=%08b\n" ,x>>1)
}
