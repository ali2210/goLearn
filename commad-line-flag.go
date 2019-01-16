package main
 // command line flag
import (
	"fmt"
	"bufio"
	"os"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
)


func input() string{
var input string
	read := bufio.NewScanner(os.Stdin)
	for read.Scan(){
		input = read.Text()
		break;
	}
	return input
}


func main(){	
	var choice string
	wordptr := flag.String("sha256", "sha384","sha512")
	flag.Parse()
	fmt.Println("you select:",*wordptr)
	switch *wordptr{
		case "sha256":
			choice = input()
			hashTx := sha256.Sum256([]byte(choice))
			fmt.Printf("%x",hashTx)
		case "sha384":
			choice = input()
			hashTx := sha512.Sum384([]byte(choice))
			fmt.Printf("%x",hashTx)
		case "sha512":
			choice = input()
			hashTx := sha512.Sum512([]byte(choice))
			fmt.Printf("%x",hashTx)
	}

	
	
}
