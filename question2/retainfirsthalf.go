package main

import (
"fmt" 
//"github.com/01-edu/z01"
)

func main() {
fmt.Println(RetainFirstHalf("Emmanuel"))
}
 
func RetainFirstHalf(s string) string{
	
	if s == "" {
		return ""
	}

	if len(s) == 1 {
		return s
	}
    
 
/*
	// 
	var count int
	for range s {
		count++
	}
	half := count / 2
	return s[:half]

*/

// method from my check point guide 
	var halflength int
	halflength = len(s) / 2
	return  s[:halflength]

}