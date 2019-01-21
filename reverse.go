package main

import (
	"fmt"

)

/*reverse the number using pointer*/

func reverse(a *[]int) *[]int{
	for i :=0; i< len(*a); i++{
		for j:= i+1; j < len(*a); j++{
			(*a)[i],(*a)[j] = (*a)[j],(*a)[i]
		}
	}
	return a
}


func main(){	
	arr := []int{1,2,3,4,5}
	fmt.Println("\nreverse:",reverse(&arr))
}
