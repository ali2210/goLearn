package main

import (
	"fmt"
)

func main() {
	arr := [4][4]string{}
	board := arrInit(arr)
	fmt.Printf("%v",board)
}

func arrInit(arr [4][4]string)*[4][4]string{
  for i := 0; i < 4; i++{
	for j := 0; j< 4;j++{
		arr[i][j] = "x"
	}
	
  }
   return &arr
}
	
