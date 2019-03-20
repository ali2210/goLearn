package main

import (
	"fmt"
)

type Kelvin float64
type Celsius float64

const (
	AbsoluteZeroKelvin Celsius = 273.15
	AbsoluteZeroCelsius Kelvin = 273.15
	)

func main() {
	var k Kelvin = 1
	var c Celsius = 1
	kelv := KtoC(k)
	cels := CtoK(c)
	fmt.Println("value of kelvin", kelv , "value of celsius", cels )
}

func KtoC(k Kelvin) Celsius{
	return Celsius(k - AbsoluteZeroCelsius)
}

func CtoK(c Celsius) Kelvin{
    	return Kelvin(AbsoluteZeroKelvin + c)
}
