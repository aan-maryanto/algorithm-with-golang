package main

import (
	"fmt"
	"strings"
)

func isPalindromeStringRecursive(str string) bool {
	if str == "" || len(str) <= 1 {
		return true
	}

	if str[0] != str[len(str)-1] {
		return false
	}

	return isPalindromeStringRecursive(str[1 : len(str)-1])

}

func isPalindromeTwoPointer(str string) bool {
	if str == "" || len(str) <= 1 {
		return true
	}

	i := 0
	j := len(str) - 1
	for i < j {
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func isPalindromInt(n int) bool {
	var temp = n
	var reverse int
	for n != 0 {
		remain := n % 10
		reverse = reverse*10 + remain
		n = n / 10
	}
	if temp == reverse {
		fmt.Printf("%d with %d is palindrome", temp, reverse)
		return true
	}
	fmt.Printf("%d with %d is not palindrome", temp, reverse)
	return false
}

func isPalindromString(str string) string {
	var rev strings.Builder

	if str == "" || len(str) <= 1 {
		rev.WriteString(str)
		return str + " with " + rev.String() + " is palindrom "
	}
	j := len(str)
	for j > 0 {
		rev.WriteString(str[j-1 : j])
		j--
	}
	if reverse := rev.String(); str == reverse {
		return str + " with " + reverse + " is palindrom "
	} else {
		return str + " with " + reverse + " is not palindrom "
	}

}

func main() {
	println(isPalindromeStringRecursive("FINKA"))
	println(isPalindromeTwoPointer("MADAM"))
	println(isPalindromInt(323))
	println(isPalindromString("MADAM"))
}
