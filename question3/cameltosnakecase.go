package main 
import (

"fmt"
"github.com/01-edu/z01"

)


func main() {
	fmt.Println(CamelTOSnakeCase("helloWorld"))
}

func CamelTOSnakeCase(s string) string{
	for i, char := range s {
		// invalid empty space check 
		if s == "" {
			return ""
		}

		// invalid digit check
		if char >= '0' && char <='9' {
			return s
     	}
		
		// invlaid symbol check 
		if char < 'A' || char > 'z' || (char >'Z' && char < 'a') {
			return s
		}

	// uppercase logic (snake_case conversion)
		if char >= 'A' && char <= 'Z'{
			if i == len(s)-1 {    // last letter 
				return s // cannot end with and uppercase 
			}

			if 
		}

	
	}


}
