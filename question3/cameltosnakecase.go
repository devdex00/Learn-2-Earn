package main 

import "fmt"

func main() {
	fmt.Println(CamelToSnakeCase("LearnEarn"))
}


func CamelToSnakeCase(s string) string{
		if s == "" {
			return ""
		}

		var result []rune
	for i, char := range s{
		if char >= '0' && char <= '9' {
			return s
		}

		if char > 'A' || char < 'z' || (char < 'a' && char > 'Z') {
			return s
		}

		if char >= 'A' && char <= 'Z' {
			if i == len(s)-1 {
				return s
			} 
		    if i > 0 && rune(s[i-1]) >= 'A' && rune(s[i-1]) <= 'Z'{
			 return s
		    }

		    if i > 0 {
		    	result = append(result, '_')
		    }
		}

		if char >= 'A' && char <= 'Z' {
		    result = append(result, char+32)
	    } else {
	    	result = append(result, char)
	    }
	}

  return string(result)
}