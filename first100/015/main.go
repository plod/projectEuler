/*
Projecteuler problem 15 https://projecteuler.net/problem=15

Starting in the top left corner of a 2×2 grid, and only being able to move to the right and down, there are exactly 6 routes to the bottom right corner.

How many such routes are there through a 20×20 grid?
*/
//used http://www.mathblog.dk/project-euler-15/ to understand problem mathematically
package main

import "log"

const gridSize = 20

func main() {
	grid := [gridSize + 1][gridSize + 1]int{}

	//initalise the grid with boundary conditions
	for i := 0; i < gridSize; i++ {
		grid[i][gridSize] = 1
		grid[gridSize][i] = 1
	}
	//calculate the paths
	for i1 := gridSize - 1; i1 >= 0; i1-- {
		for i2 := gridSize - 1; i2 >= 0; i2-- {
			grid[i1][i2] = grid[i1+1][i2] + grid[i1][i2+1]
		}
	}
	log.Println(grid[0][0])
}
