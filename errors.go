package main

import (
	"fmt"
	"errors"
)

func main() {
	div, err := divide(0,3)
	if err != nil{
		fmt.Println(div)
	}
	fmt.Println(div,err)
}

func divide(x , y int )(int, error){
	if x == 0 {
		return x/y , nil
	}
	return x/y, errors.New("PASS")
}
