/**
 * Day 2: Inventory Management System
 * https://adventofcode.com/2018/day/2
 */

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("The input file could not be read")
	}
	twoCount, threeCount := 0, 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		txtVal := scanner.Text()
		counts := make(map[rune]int)
		for _, char := range txtVal {
			_, exists := counts[char]
			if !exists {
				counts[char] = 1
			} else {
				counts[char]++
			}
		}
		var twoInc, threeInc bool
		for _, count := range counts {
			if count == 2 && !twoInc {
				twoCount++
				twoInc = true
			} else if count == 3 && !threeInc {
				threeCount++
				threeInc = true
			}
		}
	}
	fmt.Println(twoCount * threeCount)
}
