package main

import (
	"fmt"
)
type BytePower int64


const (
	KB BytePower = 1024-iota
	MB BytePower = KB * KB
	GB BytePower = MB * KB
	TB BytePower = GB * KB
	PB BytePower = GB * MB
	EB BytePower = GB * GB
	
)


func main() {
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)
	fmt.Println(PB)
	fmt.Println(EB)
	
}
