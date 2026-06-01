package main 

import "fmt"

func main() {
   fmt.Println(CheckNumber("hello"))
}

func CheckNumber(s string) bool{

	for _, char := range s {
		if (char >= '0' && char <= '9') {
			return true
		} 
    }
	return false 
}