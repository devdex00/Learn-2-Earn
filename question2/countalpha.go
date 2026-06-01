package main 

import (

"fmt"
//"github.com/01-edu/z01"

)


func main() {
	fmt.Println(CheckNumber(""))
}

func CheckNumber(s string) int {

	var count int

	if s == "" {
		return 0
	}

	for _, char := range s {

		if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' || char == ' ' {
			count++
		}
	}
		return count 
}