package main 

import (
	
    //"fmt"
	"github.com/01-edu/z01"
	)

func main() {

  //w := "Hello"
  var w string
  w = "Hello123"
	for _, char := range w {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			//fmt.Println(i, string(char)) test
			z01.PrintRune(char)
			
		}
	}
	z01.PrintRune('\n')
   return
}

