package main

import "fmt"

func main() {
fmt.Println(RectPerimeter(4, 4))
}

func RectPerimeter(w, h int) int{
 var perimeter int 
	if w < 0 || h < 0 {
		return -1
	} 

	perimeter = 2 * (w + h)

	return perimeter 

}