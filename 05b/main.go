/**
 * Day 5: Alchemical Reduction
 * https://adventofcode.com/2018/day/5
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("The input file could not be read")
	}
	scanner := bufio.NewScanner(f)
	scanner.Scan()

	// We can do an initial reduction before removing individual units (result will be the same)
	initialPoly := reducePolymer(scanner.Text())

	// Remove each unit in parallel
	results := make([]int, 'z'-'a'+1)
	resChan := make(chan int)
	for i := 'a'; i <= 'z'; i++ {
		go shortestReductionLenWithoutUnit(initialPoly, i, resChan)
	}
	for i := 0; i <= 'z'-'a'; i++ {
		results[i] = <-resChan
	}
	lowestRes := len(initialPoly)
	for _, i := range results {
		if i < lowestRes {
			lowestRes = i
		}
	}
	fmt.Printf("Lowest reduction length: %d\n", lowestRes)
}

func shortestReductionLenWithoutUnit(poly string, exLower rune, c chan int) {
	filteredPoly := strings.Replace(poly, string(exLower), "", -1)
	filteredPoly = strings.Replace(filteredPoly, string(unicode.ToUpper(exLower)), "", -1)
	reductionLen := len(reducePolymer(filteredPoly))
	fmt.Printf("Found reduction of length: %d by excluding %s\n", reductionLen, string(exLower))
	c <- reductionLen
}

func reducePolymer(poly string) string {
	res := ""
	for idx := 0; idx < len(poly); {
		if idx == len(poly)-1 {
			// Special case for last character -- even if we reach it, it can't react
			res += string(poly[idx])
			break
		}
		c0, c1 := rune(poly[idx]), rune(poly[idx+1])
		if unicode.ToLower(c0) == unicode.ToLower(c1) && unicode.IsLower(c0) != unicode.IsLower(c1) {
			idx += 2
		} else {
			res += string(c0)
			idx++
		}
	}
	fmt.Printf("Polymer length: %d\n", len(res))
	if len(res) == len(poly) {
		return res
	}
	return reducePolymer(res)
}
