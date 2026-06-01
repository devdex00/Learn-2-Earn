package main

import "fmt"

func main() {


fmt.Println(PintRectPerimeter())

}

func PintRectPerimeter(x int, y int) int{

	if x <= 0 || y <= 0 {
		return -1
	}

	perimeter := 2 * x + 2 * y

	return perimeter

}