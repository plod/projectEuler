/*
Projecteuler problem 67 https://projecteuler.net/problem=67
By starting at the top of the triangle below and moving to adjacent numbers on the row below, the maximum total from top to bottom is 23.

3
7 4
2 4 6
8 5 9 3

That is, 3 + 7 + 4 + 9 = 23.

Find the maximum total from top to bottom in triangle.txt (right click and 'Save Link/Target As...'), a 15K text file containing a triangle with one-hundred rows.

NOTE: This is a much more difficult version of Problem 18. It is not possible to try every route to solve this problem, as there are 299 altogether! If you could check one trillion (1012) routes every second it would take over twenty billion years to check them all. There is an efficient algorithm to solve it. ;o)
*/

package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func convertTheTriangle(theTriangle string) [][]string {
	file, _ := os.Open(theTriangle)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var returnSlice [][]string
	for scanner.Scan() {
		sliceLine := strings.Replace(scanner.Text(), "  ", "", -1)
		returnSlice = append(returnSlice, strings.Split(sliceLine, " "))
	}
	return returnSlice
}

func main() {
	theTriangle := "./p067_triangle.txt"
	theTriangleSlice := convertTheTriangle(theTriangle)
	for i := len(theTriangleSlice) - 2; i >= 0; i-- {
		//log.Println("current line looks like ", theTriangleSlice[i])
		for i2 := 0; i2 < len(theTriangleSlice[i]); i2++ {
			currentNumber, _ := strconv.Atoi(theTriangleSlice[i][i2])
			for i3 := 0; i3 < 2; i3++ {
				inTriangleNumber, _ := strconv.Atoi(theTriangleSlice[i][i2])
				comparingNumber, _ := strconv.Atoi(theTriangleSlice[i+1][i2+i3])
				//log.Println("I am checking if", currentNumber, " + ", comparingNumber, "(", currentNumber+comparingNumber, ") is greater than", inTriangleNumber)
				if currentNumber+comparingNumber > inTriangleNumber {
					theTriangleSlice[i][i2] = strconv.Itoa(currentNumber + comparingNumber)
					//log.Println("it is")
				} /*else {
					log.Println("it is not")
				}*/
			}
		}
		//log.Println("line now looks like ", theTriangleSlice[i])
	}
	log.Println(theTriangleSlice[0][0])
}
