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
	closestPoiIdx  int
	closestPoiDist int
}

const invalidPOI = -1

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
	areaCounts := make(map[int]int)
	populateMapPointsFromPOIs(mapPoints, pois, areaCounts)
	borderPois := findBorderPOIs(mapPoints, x0, y0, x1, y1)

	maxAreaPOI := 0
	for i := range areaCounts {
		if i != invalidPOI && areaCounts[i] > areaCounts[maxAreaPOI] && !borderPois[i] {
			maxAreaPOI = i
		}
	}
	fmt.Println("Largest area:", areaCounts[maxAreaPOI])
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
			mapPoints[x][y] = &mapPoint{invalidPOI, math.MaxInt32}
		}
	}
	return
}

func populateMapPointsFromPOIs(mapPoints map[int]map[int]*mapPoint, pois []poi, areaCounts map[int]int) {
	for i, p := range pois {
		for x := range mapPoints {
			for y := range mapPoints[x] {
				mapP := mapPoints[x][y]
				dist := abs(p.x-x) + abs(p.y-y)
				if dist < mapP.closestPoiDist {
					// Keep track of how much area each POI covers
					areaCounts[mapP.closestPoiIdx]--
					areaCounts[i]++
					mapP.closestPoiIdx = i
					mapP.closestPoiDist = dist
				} else if dist == mapP.closestPoiDist {
					// Points equally distant to two or more POIs don't count for any of them
					areaCounts[mapP.closestPoiIdx]--
					mapP.closestPoiIdx = invalidPOI
				}
			}
		}
	}
}

func findBorderPOIs(mapPoints map[int]map[int]*mapPoint, x0, y0, x1, y1 int) (borderPois map[int]bool) {
	borderPois = make(map[int]bool)
	for x := x0; x <= x1; x++ {
		borderPois[mapPoints[x][y0].closestPoiIdx] = true
		borderPois[mapPoints[x][y1].closestPoiIdx] = true
	}
	for y := y0; y <= y1; y++ {
		borderPois[mapPoints[x0][y].closestPoiIdx] = true
		borderPois[mapPoints[x1][y].closestPoiIdx] = true
	}
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
