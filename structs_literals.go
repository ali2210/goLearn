package main


import ("fmt")


type Point struct{
	X,Y int
}

type Circle struct{
	Center Point   // embed structs-1
	Radius int
}

type Wheel struct{
	Circle Circle  //embed structs-2
	Spokes int
}

func main(){
	var w Wheel
	w.Circle.Center.X =2
	w.Circle.Center.Y =2
	w.Circle.Radius =3
	w.Spokes =3
	fmt.Println(w)
}
	
