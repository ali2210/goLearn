package main

import (
	"fmt"

)

type List struct{
	Data int
	Next *List
}
var Node List

func main(){

	add(1)
	add(2)
	add(3)
	add(4)
}

func createNode(i int)*List{
	Node.Data = i
	Node.Next = nil
	return &Node
}

func add(i int){
	var head = &Node
	var temp = head
	var newNode *List
	if (*head).Next == nil{
		 newNode = createNode(i)
		 (*head).Next = newNode
		  //fmt.Println("data", (*head).Data, "next", (*head).Next)
		printList(head)	
	}else{
		for (*temp).Next != head {
				temp = (*temp).Next
		}
		  newNode := createNode(i)
		  (*temp).Next = newNode
		    //fmt.Println("data", (*temp).Data, "next", (*temp).Next)	 
			printList(temp) 
	}
	
	
}

func printList(head *List){
	var temp = head
	
	if temp == nil{
		fmt.Println("List empty")
	}else{
		if (*temp).Next != nil{
			fmt.Println("Data->",(*temp).Data)
			fmt.Println("Next->",(*temp).Next)
		}
	}
}


