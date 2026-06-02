package main 

import "fmt"

func main() {

	// understanding nested loop
outerloop : // just learnt labels 
	for i := 1 ; i <= 3 ; i++ {
		for j := 1 ; j < 4 ; j++ {
			if i == j {
				continue outerloop // specify the outer loop to move to using label
			}
			fmt.Println(i, j)
		}
	}

}