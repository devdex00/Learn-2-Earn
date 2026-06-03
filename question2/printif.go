package main 

import "fmt"

func main() {

	fmt.Println(PrintIf("DEX"))

}

func PrintIf(s string) string{
	for range s {
	
		if len(s) == 0 || len(s) >= 3 {
			return "G\n"
		}
	}
	return "invalid input\n"
}