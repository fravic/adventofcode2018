/**
 * Day 5: Alchemical Reduction
 * https://adventofcode.com/2018/day/5
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("The input file could not be read")
	}
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	fmt.Println(reducePolymer(scanner.Text()))
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
