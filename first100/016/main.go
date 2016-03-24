/*
Projecteuler problem 16 https://projecteuler.net/problem=16

2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.

What is the sum of the digits of the number 2^1000?
*/
package main

import (
	"log"
	"math/big"
	"strconv"
)

const powerOfTwo = 1000

func findPower(i int64) *big.Int {
	returnVal := new(big.Int).Exp(big.NewInt(2), big.NewInt(i), nil)
	log.Printf("Power Return Val = %d\n", returnVal)
	return returnVal
}

func sumOfDigits(number *big.Int) int {
	stringOfNumber := number.String()
	var sum int
	for i := 0; i < len(stringOfNumber); i++ {
		var digit int
		digit, _ = strconv.Atoi(string(stringOfNumber[i]))
		log.Println(digit)
		sum += digit
	}
	return sum
}

func main() {
	log.Println(sumOfDigits(findPower(powerOfTwo)))
}
