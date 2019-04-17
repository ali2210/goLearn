package main

import (
	"fmt"
	"encoding/json"
)

type Books struct{
	Title string 
	Status bool
	Writer string 
	Comments string
}

		


func main() {
	var Library =[]Books{
		{Title:"Lord of Rings the fellow ship", Status:false, Writer:"J-J-R-Tolkein",Comments:""},
		{Title:"Quantum Aspects of Life",Status: false, Writer:"Paul-C",Comments:""},
		{Title:"You were born rich", Status:true, Writer:"Bob Proctor",Comments:"good book"},
		}
	if lib , err := json.MarshalIndent(&Library,""," "); err== nil{
			fmt.Printf("%s",lib)
		
	}
	
	
}
