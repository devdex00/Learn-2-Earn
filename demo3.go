func CamelToSnakeCase(s string) string {
	if s == "" {
		return ""
	}

	var result []rune
	for i, char := range s {
		// 1. Invalid character checks
		if char >= '0' && char <= '9' {
			return s // Numbers are not allowed in camelCase
		}
		if char < 'A' || char > 'z' || (char > 'Z' && char < 'a') {
			return s // Only letters are allowed
		}

		// 2. Uppercase logic (snake_case conversion)
		if char >= 'A' && char <= 'Z' {
			if i == len(s)-1 {
				return s // Cannot end with an uppercase letter
			}
			if i > 0 && rune(s[i-1]) >= 'A' && rune(s[i-1]) <= 'Z' {
				return s // Cannot have two uppercase letters in a row
			}
			if i > 0 {
				result = append(result, '_')
			}
		}

		// 3. Add the character (lowercased for camelCase)
		if char >= 'A' && char <= 'Z' {
			result = append(result, char+32) // Convert to lowercase
		} else {
			result = append(result, char)
		}
	}

	return string(result)
}
