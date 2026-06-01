package main 

import "fmt"

func main() {

	fmt.Println(PrintIfNot("DE"))

}

func PrintIfNot(s string) string{

	if len(s) < 3 {
		return "G\n"
	}
 return "Invalid input\n"
}