package main 

import "fmt"


func main() {

	fmt.Println(CountChar("hellllo", 'l'))

}

func CountChar(s string, c rune) int{

	var count int

	if s == "" {
		return 0
	}

	for _, char := range s {
		if char == c {
			count++
		}
	}
	return count 
}