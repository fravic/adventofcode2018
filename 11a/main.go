/**
 * Day 11: Chronal Charge
 * https://adventofcode.com/2018/day/11
 */

package main

import (
	"fmt"
	"os"
	"strconv"
)

const gridSize = 300

type powerGrid = map[int]map[int]int

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go serialNumber")
		os.Exit(1)
	}
	serial, _ := strconv.Atoi(os.Args[1])

	// Calculate power levels starting from the bottom right (so we only have to iterate once)
	maxPowX := gridSize
	maxPowY := gridSize
	maxPow := 0
	grid := make(powerGrid)
	for diagonal := gridSize; diagonal > 0; diagonal-- {
		for x := gridSize; x > diagonal; x-- {
			calculatePowerLevel(x, diagonal, serial, grid, &maxPowX, &maxPowY, &maxPow)
		}
		for y := gridSize; y > diagonal; y-- {
			calculatePowerLevel(diagonal, y, serial, grid, &maxPowX, &maxPowY, &maxPow)
		}
		calculatePowerLevel(diagonal, diagonal, serial, grid, &maxPowX, &maxPowY, &maxPow)
	}
	fmt.Printf("Max power level: %d at (%d,%d)\n", maxPow, maxPowX, maxPowY)
}

func calculatePowerLevel(x, y, serial int, grid powerGrid, maxPowX, maxPowY, maxPow *int) {
	// Calculate the power of this cell
	rackID := x + 10
	pow := (rackID * y) + serial
	pow *= rackID
	powStr := strconv.Itoa(pow)
	pow = int(powStr[len(powStr)-3]-'0') - 5

	// Record in the grid
	_, exists := grid[x]
	if !exists {
		grid[x] = make(map[int]int)
	}
	grid[x][y] = pow

	// Add on the power levels of the 8 other cells that determine this cell's total power (if they exist)
	totalPow := 0
	for x0 := x; x0 <= x+2 && x0 <= gridSize-2; x0++ {
		for y0 := y; y0 <= y+2 && y0 <= gridSize-2; y0++ {
			totalPow += grid[x0][y0]
		}
	}
	if totalPow > *maxPow {
		*maxPowX = x
		*maxPowY = y
		*maxPow = totalPow
	}
}
