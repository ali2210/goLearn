package main

import ("fmt")


type Point struct{
	X,Y int
}

type Circle struct{
	Point	//anonymous case#1
	Radius int
}

type Wheel struct{
	Circle 	//anonymous case#2
	Spokes int
}


func main(){

	var wheel Wheel
	wheel.Circle.Point.X =4
	wheel.Circle.Point.Y =4
	wheel.Circle.Radius =4
	wheel.Spokes =4
		fmt.Println(wheel)
	w:= Wheel{Circle{Point{4,4},3},2}
		fmt.Printf("%v",w)
}
