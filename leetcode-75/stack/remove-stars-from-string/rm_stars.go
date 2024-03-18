package removestarsfromstring

import "fmt"

type Stack []rune

func (s *Stack) push(c rune) {
	*s = append((*s), c)
}

func (s *Stack) pop() {
	if len(*s) <= 0 {
		fmt.Println("empty stack")
	} else {
		*s = (*s)[:len(*s)-1]
	}
}

func removeStars(s string) string {
	var resultString string
	var strStack Stack

	for _, c := range s {
		if c == '*' {
			strStack.pop()
		} else {
			strStack.push(c)
		}
	}

	resultString = string(strStack)

	return resultString
}
