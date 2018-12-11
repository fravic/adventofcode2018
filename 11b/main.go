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

type squareIdentifier struct {
	x, y, squareSize, pow int
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go serialNumber")
		os.Exit(1)
	}
	serial, _ := strconv.Atoi(os.Args[1])
	maxIdentifier := &squareIdentifier{}
	grid := calculatePowerGrid(serial)
	areaPowerCache := make(map[int]powerGrid)
	for i := 1; i <= gridSize; i++ {
		calculateMaxAreaPowerForSquareSize(i, grid, areaPowerCache, maxIdentifier)
	}
	fmt.Printf("\nMax power level: %d at (%d,%d,%d)\n", maxIdentifier.pow, maxIdentifier.x, maxIdentifier.y, maxIdentifier.squareSize)
}

// Calculate and cache power levels for each grid cell
func calculatePowerGrid(serial int) (grid powerGrid) {
	grid = make(powerGrid)
	for x := 1; x < gridSize; x++ {
		for y := 1; y < gridSize; y++ {
			calculatePowerLevel(x, y, serial, grid)
		}
	}
	return
}

func calculatePowerLevel(x, y, serial int, grid powerGrid) {
	// Calculate the power of this cell
	rackID := x + 10
	pow := (rackID * y) + serial
	pow *= rackID
	powStr := strconv.Itoa(pow)
	pow = int(powStr[len(powStr)-3]-'0') - 5

	// Record in the grid
	createNestedMapIfNecessary(grid, x)
	grid[x][y] = pow
}

func calculateMaxAreaPowerForSquareSize(squareSize int, grid powerGrid, areaPowerCache map[int]powerGrid, maxIdentifier *squareIdentifier) {
	fmt.Print(".")
	for diagonal := gridSize; diagonal > 0; diagonal-- {
		for x := gridSize; x > diagonal; x-- {
			calculateMaxAreaPowerAtPoint(x, diagonal, squareSize, grid, areaPowerCache, maxIdentifier)
		}
		for y := gridSize; y > diagonal; y-- {
			calculateMaxAreaPowerAtPoint(diagonal, y, squareSize, grid, areaPowerCache, maxIdentifier)
		}
		calculateMaxAreaPowerAtPoint(diagonal, diagonal, squareSize, grid, areaPowerCache, maxIdentifier)
	}
}

func calculateMaxAreaPowerAtPoint(x, y, s int, grid powerGrid, areaPowerCache map[int]powerGrid, maxIdentifier *squareIdentifier) {
	if x+s-1 > gridSize || y+s-1 > gridSize {
		return
	}

	// totalPow[s][x][y] = totalPow[s-1][x+1][y+1] + SUM(pow[x0][y]) + SUM(pow[x][y0]) + pow[x][y]
	areaPow := areaPowerCache[s-1][x+1][y+1]
	for x0 := x + 1; x0 < x+s; x0++ {
		areaPow += grid[x0][y]
	}
	for y0 := y + 1; y0 < y+s; y0++ {
		areaPow += grid[x][y0]
	}
	areaPow += grid[x][y]

	// Update the maximum if necessary
	if areaPow > maxIdentifier.pow {
		*maxIdentifier = squareIdentifier{
			x, y, s, areaPow,
		}
	}

	// Update the areaPowerCache for future calculations
	_, exists := areaPowerCache[s]
	if !exists {
		areaPowerCache[s] = make(powerGrid)
	}
	createNestedMapIfNecessary(areaPowerCache[s], x)
	areaPowerCache[s][x][y] = areaPow
}

func createNestedMapIfNecessary(m map[int]map[int]int, i int) {
	_, exists := m[i]
	if !exists {
		m[i] = make(map[int]int)
	}
}
