package main
              /*different hash*/
import (
	"fmt"
	"crypto/sha256"
)


func isSame(a, b [32]byte) {
   var count int = 0
   var j int =0

	for i, _ := range a{
		if a[i] != b[j]{
			fmt.Printf("[a]:%x\t[b:]%x",a[i],b[j])
			j+=1
			count +=1
		 }
		fmt.Println(count)
	}
	
}


func main(){	
	var aliceText string = "key"
	var bobText string = "123" 	
	alice:= sha256.Sum256([]byte(aliceText))
	bob := sha256.Sum256([]byte(bobText))
	fmt.Printf("\t alice send message:%x",alice)
	fmt.Printf("\t bob receive message:%x",bob)
	isSame(alice,bob)
}
