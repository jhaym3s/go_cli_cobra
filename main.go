package main

import (
	"fmt"
)

func main() {

	number := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	reversedNumber := reverseSlice(number)
	maxNumber := findMaxNumber(number)
	for i := 0; i < len(number)/2; i++ {
		number[i], number[len(number)-i-1] = number[len(number)-i-1], number[i]
	}

	fmt.Println(reversedNumber, number, maxNumber)
}

func reverseSlice(slice []int) []int {
	length := len(slice)
	reversed := make([]int, length)

	for i, v := range slice {
		reversed[length-i-1] = v
	}
	return reversed
}
 
func findMaxNumber(slice []int)int{
	max := slice[0]
	for _, v := range slice {
	 if v > max {
		max = v
	}
	}
	return max
}
