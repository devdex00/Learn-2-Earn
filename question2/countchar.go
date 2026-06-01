package main 

import "fmt"


func main() {

	fmt.Println(CountChar("hellllo", 'l'))

}

func CountChar(s string, c rune) int{

	var count int

	for _, char := range s {
		if char == c {
			count++
		}
	}
	return count 
}