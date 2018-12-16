/**
 * Day 12: Subterranean Sustainability
 * https://adventofcode.com/2018/day/12
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const numGenerations = 20

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading input file")
		os.Exit(0)
	}
	scanner := bufio.NewScanner(f)

	initialStateStr := getNextLine(scanner)[15:]
	states := make(map[int]bool)
	minIdx := 0
	maxIdx := 0
	for i, c := range initialStateStr {
		states[i] = c == '#'
		minIdx = min(i, minIdx)
		maxIdx = max(i, maxIdx)
	}

	getNextLine(scanner) // skip empty line

	rules := make(map[int]bool)
	for scanner.Scan() {
		ruleStrs := strings.Split(scanner.Text(), " => ")
		rule := 0
		for _, c := range ruleStrs[0] {
			if c == '#' {
				rule = (rule << 1) | 1
			} else {
				rule = rule << 1
			}
		}
		rules[rule] = ruleStrs[1][0] == '#'
	}

	lastTotal := 0
	for i := 0; i < numGenerations; i++ {
		minIdx, maxIdx = spawnNextGeneration(states, rules, minIdx, maxIdx)
		total := calculateTotal(states, minIdx, maxIdx)
		fmt.Printf("Gen %d, total %d (delta %d)\n", i+1, total, total-lastTotal)
		lastTotal = total
	}
}

func calculateTotal(states map[int]bool, minIdx, maxIdx int) (total int) {
	total = 0
	for i := minIdx; i <= maxIdx; i++ {
		if states[i] {
			total += i
		}
	}
	return
}

func getNextLine(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}

func spawnNextGeneration(
	states map[int]bool,
	rules map[int]bool,
	minIdx int,
	maxIdx int,
) (newMinIdx int, newMaxIdx int) {
	nextPlants := make(map[int]bool)
	for i := minIdx - 2; i < maxIdx+2; i++ {
		pattern := getPatternForPlant(states, i)
		newState, exists := rules[pattern]
		if exists {
			nextPlants[i] = newState
			if newState {
				minIdx = min(minIdx, i)
				maxIdx = max(maxIdx, i)
			}
		} else {
			nextPlants[i] = false
		}
	}
	for k, v := range nextPlants {
		states[k] = v
	}
	return minIdx, maxIdx
}

func getPatternForPlant(states map[int]bool, plantIdx int) (res int) {
	for i := plantIdx - 2; i <= plantIdx+2; i++ {
		state, exists := states[i]
		if exists && state {
			res = (res << 1) | 1
		} else {
			res = res << 1
		}
	}
	return
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
