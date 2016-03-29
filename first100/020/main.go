/*
Projecteuler problem 20 https://projecteuler.net/problem=20
n! means n × (n − 1) × ... × 3 × 2 × 1

For example, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800,
and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.

Find the sum of the digits in the number 100!
*/

package main

import (
	"log"
	"math/big"
	"strconv"
)

const factorialCare = 100

func main() {

	c := make(chan *big.Int)
	go factorial(factorialCare, c)

	for n := range c {
		log.Println(addDigits(n))
	}
}

func factorial(number int, c chan *big.Int) {
	returnInt := new(big.Int)
	returnInt.SetString(strconv.Itoa(number), 0)
	for i := number - 1; i > 0; i-- {
		iBigInt := new(big.Int)
		iBigInt.SetString(strconv.Itoa(i), 0)
		returnInt = returnInt.Mul(returnInt, iBigInt)
	}
	c <- returnInt
	close(c)
	return
}

func addDigits(n *big.Int) int {
	asString := n.String()
	var digitTotal int
	for i := 0; i < len(asString); i++ {
		intOfDigit, _ := strconv.Atoi(asString[i : i+1])
		digitTotal += int(intOfDigit)
	}
	return digitTotal
}
