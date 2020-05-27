package main

import "fmt"

// 20. Valid Parentheses (Easy)
func isValid(s string) bool {
	if s == "" {
		return true
	}
	maps := map[rune]rune{
		'{': '}',
		'[': ']',
		'(': ')',
	}
	stack := []rune{}
	for _, char := range s {
		if char == '{' || char == '[' || char == '(' {
			stack = append(stack, char)
		} else {
			if len(stack) > 0 && maps[stack[len(stack)-1]] == char {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("{]"))
}
