/**
 * Day 3: No Matter How You Slice It
 * https://adventofcode.com/2018/day/3
 */

package main

import (
	"bufio"
	"fmt"
	"os"
)

const fabricSize = 1000

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("The input file could not be opened")
	}
	scanner := bufio.NewScanner(f)
	fabric := make(map[int]map[int]int)
	for i := 0; i < fabricSize; i++ {
		fabric[i] = make(map[int]int)
	}
	for scanner.Scan() {
		line := scanner.Text()
		var id string
		var x, y, w, h int
		fmt.Sscanf(line, "#%s @ %d,%d: %dx%d", &id, &x, &y, &w, &h)
		for i := 0; i < w; i++ {
			for j := 0; j < h; j++ {
				fabric[x+i][y+j]++
			}
		}
	}
	overlapping := 0
	for i := 0; i < fabricSize; i++ {
		for j := 0; j < fabricSize; j++ {
			if fabric[i][j] > 1 {
				overlapping++
			}
		}
	}
	fmt.Println(overlapping)
}
