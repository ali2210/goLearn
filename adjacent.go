package main

import (
	"fmt"

)

//adjacent elimination

func count(strings []string , s string)int{
	 var count int =0	
	for _,v := range strings{
		if v == s {
			count+=1
		}
	}
	return count

}

func eliminate(s []string, loc int)[]string{
	for i, _ := range s{
		if s[i] == s[loc]{
			s[loc] = ""
		}
	}
	return s
}

func locations(s []string , str string)int{
	var index int =0  
	for i, _ := range s{
		if s[i] == str{
			index = i
		}
	}
	return index
}

func main(){	
	arr := []string{"hello","go","hello"}
	if count(arr,"hello") > 1{
		index := locations(arr,"hello")
		fmt.Println("eliminate:",eliminate(arr,index))
	}		
}
	
