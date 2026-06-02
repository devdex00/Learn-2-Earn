package main 

import "fmt"

func main() {

	/*// classic for loop pattern 
	for i := 0 ; i <= 5 ; i++ {
		fmt.Println(i)
	}


fmt.Println("\n")

	// conditions only for loop pattern
	/*var x int 
	x = 0
	for x <= 5 {
		fmt.Println(x)
		x++
	}

	// infinite for loop 
	/*for {
		fmt.Println("HI")
	}

	// return value test 

	var age int
	age = 25 
	if age < 20 {
		fmt.Println("Too young")
		return 
	} else {
		fmt.Println("Allowed")
	}*/

	for i := 0 ; i <= 10 ; i++ {
		if i%2 != 0 {
			continue  // skips the execution cycle for 3 and then move on with the other iteration
		}
		fmt.Println(i)
	}
	
}