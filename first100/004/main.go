/*
Projecteuler problem 4 https://projecteuler.net/problem=4

A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.

*/
package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

const digitLimit = 3

func isPalindrome(compositeNumber string) bool {
	byteLimit := len(compositeNumber) / 2
	for i := 0; i < byteLimit; i++ {
		if string(compositeNumber[i]) != string(compositeNumber[len(compositeNumber)-(i+1)]) {
			return false
		}
	}
	return true
}

func main() {
	var largestPalindromic int
	var palindromics = []int{}
	var maxNumber int
	for i := 0; i < digitLimit; i++ {
		if i != 0 {
			maxNumber *= 10
		}
		maxNumber += 9
	}
	iterationLimit := int(math.Pow10(digitLimit - 1))
	for iLeft := maxNumber; iLeft > iterationLimit; iLeft-- {
		for iRight := maxNumber; iRight > 1*int(math.Pow10(digitLimit-1)); iRight-- {
			compositenumber := strconv.Itoa(iLeft * iRight)
			if isPalindrome(compositenumber) {
				compositenumber, _ := strconv.Atoi(compositenumber)
				palindromics = append(palindromics, compositenumber)
				break
			}
		}
	}
	//this seems innefficient as I am getting all palindromics but after getting the wrong answer it confirmed that I would get right answer
	sort.Ints(palindromics)

	largestPalindromic = palindromics[len(palindromics)-1]

	fmt.Println(largestPalindromic)
}
