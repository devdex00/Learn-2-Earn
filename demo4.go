package main 

import "fmt"

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
}


func IsCamelCase(s string) bool{
	// word must not end in capital letter 
	lastchar := s[len(s)-1] 
	if lastchar >= 'A' && lastchar <= 'Z'{
		return false
	}
	// loops through the string except the last word 
	for i := 0; i < len(s) - 1; i++{
			// checks if there are 2 consecutive uppercase letters 
		if (s[i] >= 'A' && s[i] < 'Z') && (s[i+1] >= 'A' && s[i+1] <= 'Z'){
			return false 
		}
		// loops through the string again for symbols 
		for _, r := range s {
			if r < 'A' && r > 'Z' || r < 'a' && r > 'z' {
				return false 
			}
		}
	}
	return true
}

func CamelToSnakeCase(s string) string {
	if s == "" || !IsCamelCase(s) {
		return s
	  }
	  var output []rune
	  for i, char := range s {
		if char > 'A' && char < 'Z' && i > 0{
			output = append(output, '_')
		}
		output = append(output, char)
	}
	return string(output)
}