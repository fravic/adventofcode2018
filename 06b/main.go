/**
 * Day 6: Chronal Coordinates
 * https://adventofcode.com/2018/day/6
 */

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type poi struct {
	x, y int
}

type mapPoint struct {
	distSum int
}

const distSumThreshold = 10000

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("The input file could not be read")
	}
	scanner := bufio.NewScanner(f)
	pois := make([]poi, 0)
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), ", ")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		pois = append(pois, poi{x, y})
	}
	x0, y0, x1, y1 := getBounds(pois)
	fmt.Println("Found bounds:", x0, y0, x1, y1)
	mapPoints := initMap(x0, y0, x1, y1)
	populateMapPointsFromPOIs(mapPoints, pois)

	regionSize := 0
	for x := range mapPoints {
		for y := range mapPoints[x] {
			if mapPoints[x][y].distSum < distSumThreshold {
				regionSize++
			}
		}
	}
	fmt.Println("Region size:", regionSize)
}

func getBounds(pois []poi) (x0, y0, x1, y1 int) {
	x0, y0 = math.MaxInt32, math.MaxInt32
	for _, c := range pois {
		if c.x < x0 {
			x0 = c.x
		}
		if c.y < y0 {
			y0 = c.y
		}
		if c.x > x1 {
			x1 = c.x
		}
		if c.y > y1 {
			y1 = c.y
		}
	}
	return
}

func initMap(x0, y0, x1, y1 int) (mapPoints map[int]map[int]*mapPoint) {
	mapPoints = make(map[int]map[int]*mapPoint)
	for x := x0; x <= x1; x++ {
		mapPoints[x] = make(map[int]*mapPoint)
		for y := y0; y <= y1; y++ {
			mapPoints[x][y] = &mapPoint{0}
		}
	}
	return
}

func populateMapPointsFromPOIs(mapPoints map[int]map[int]*mapPoint, pois []poi) {
	for _, p := range pois {
		for x := range mapPoints {
			for y := range mapPoints[x] {
				mapP := mapPoints[x][y]
				dist := abs(p.x-x) + abs(p.y-y)
				mapP.distSum += dist
			}
		}
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
