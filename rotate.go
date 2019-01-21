package main

import (
	"fmt"

)

/*rotate the number using pointer*/

func rotate(a *[]int) *[]int{
	for i :=0; i< len(*a); i++{
		for j:=len(*a)-1; j!=i;j--{
			(*a)[i],(*a)[j] = (*a)[j],(*a)[i]
		}
	}
	return a
}


func main(){	
	arr := []int{1,2,3,4,5}
	fmt.Println("\nrotate:",rotate(&arr))
}
