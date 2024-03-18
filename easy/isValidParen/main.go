/* Given a string s containing just the characters '(', ')', '{', '}', '[' and ']',
determine if the input string is valid.

An input string is valid if:

    Open brackets must be closed by the same type of brackets.
    Open brackets must be closed in the correct order.
    Every close bracket has a corresponding open bracket of the same type. */

package main

import "fmt"

type Stack []rune

func (s *Stack) push(c rune) {
	*s = append(*s, c)
}

func (s *Stack) pop() {
	*s = (*s)[:len(*s)-1]
}

func (s Stack) isTop(c rune) bool {
	if len(s) <= 0 {
		return false
	}
	if s[len(s)-1] == c {
		return true
	} else {
		return false
	}
}


func isValid(s string) bool {
	if len(s) < 2 {
		return false
	}
	var stack Stack

	parens := map[rune]rune {
		'}':'{',
		']':'[',
		')':'(',
	}

	for _,char  := range s {
		c := rune(char)
		for k,v := range parens {
			if v == c {
				stack.push(v)
			} 
			if k == c {
				if stack.isTop(parens[k]) {
					stack.pop()
				} else {
					return false
				}
			}
		}
	}
	if len(stack) > 0 {
		return false
	}
	return true
}

func main() {
	fmt.Println(isValid("){"))
}







