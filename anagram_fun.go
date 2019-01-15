package main

/*ANAGRAM */

import (
	"fmt"
	
)


func intToString(str_1, str_2 []byte){
 var count int =0		
		 for i,_ := range str_1{
			for j,_ :=range str_2{
				if str_1[i]== str_2[j]{
					count+=1
					if count == len(str_1){ 
						fmt.Printf("\tfinal output **** anagram*** ")
						fmt.Println("code","doce")
					}
				 }
				  
			 }
		 }
	
}

func main(){	
	str_a := []byte{'c','o','d','e'}
	str_b := []byte {'d','o','c','e'}
	intToString(str_a,str_b)
}
