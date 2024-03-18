package main

import "fmt"

func isPalindrome(x int) bool {
    if x < 0 {
        return false;
    }
	var  rem, rev int
    temp := x

	for {
		rem = x % 10
		rev = rev*10 + rem
		x /= 10
		if x == 0 {
			break // Break Statement used to exit from loop
		}
	}

	if rev == temp {
		return true
	} else {
		return false
	}

}

func main() {
	fmt.Println(isPalindrome(-101))
}
