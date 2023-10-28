package main

import "fmt"

func reversedNum(x int) int {
	var reversedNum int
	for x > 0 {
		// add the last digit, next loop multiply it by 10
		reversedNum = reversedNum*10 + x%10
		// keep the number without the last digit for the next loop
		x = x / 10
		//fmt.Println(reversedNum)
		//fmt.Println(tmp)
	}
	return reversedNum
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}

	if x%10 == 0 {
		return false
	}
	return reversedNum(x) == x
}

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(122))
}
