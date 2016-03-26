/*
Projecteuler problem 18 https://projecteuler.net/problem=18
By starting at the top of the triangle below and moving to adjacent numbers on the row below, the maximum total from top to bottom is 23.

3
7 4
2 4 6
8 5 9 3

That is, 3 + 7 + 4 + 9 = 23.

Find the maximum total from top to bottom of the triangle below:

75
95 64
17 47 82
18 35 87 10
20 04 82 47 65
19 01 23 75 03 34
88 02 77 73 07 63 67
99 65 04 28 06 16 70 92
41 41 26 56 83 40 80 70 33
41 48 72 33 47 32 37 16 94 29
53 71 44 65 25 43 91 52 97 51 14
70 11 33 28 77 73 17 78 39 68 17 57
91 71 52 38 17 14 91 43 58 50 27 29 48
63 66 04 68 89 53 67 30 73 16 69 87 40 31
04 62 98 27 23 09 70 98 73 93 38 53 60 04 23

NOTE: As there are only 16384 routes, it is possible to solve this problem by trying every route. However, Problem 67, is the same challenge with a triangle containing one-hundred rows; it cannot be solved by brute force, and requires a clever method! ;o)
*/

package main

import (
	"log"
	"strconv"
	"strings"
)

func convertTheTriangle(theTriangle string) [][]string {
	firstSlice := strings.Split(theTriangle, "\n")
	var returnSlice [][]string
	for _, sliceLine := range firstSlice {
		sliceLine := strings.Replace(sliceLine, "  ", "", -1)
		returnSlice = append(returnSlice, strings.Split(sliceLine, " "))
	}
	return returnSlice
}

func main() {
	theTriangle := `75
  95 64
  17 47 82
  18 35 87 10
  20 04 82 47 65
  19 01 23 75 03 34
  88 02 77 73 07 63 67
  99 65 04 28 06 16 70 92
  41 41 26 56 83 40 80 70 33
  41 48 72 33 47 32 37 16 94 29
  53 71 44 65 25 43 91 52 97 51 14
  70 11 33 28 77 73 17 78 39 68 17 57
  91 71 52 38 17 14 91 43 58 50 27 29 48
  63 66 04 68 89 53 67 30 73 16 69 87 40 31
  04 62 98 27 23 09 70 98 73 93 38 53 60 04 23`
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
