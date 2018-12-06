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
	scanner := bufio.NewScanner(f)
	var rows []string
	for scanner.Scan() {
		chars := scanner.Text()
		rows = append(rows, chars)
	}
	// Omit one character position at a time, check for dupes
	for i := range rows[0] {
		checked := make(map[string]bool)
		for _, row := range rows {
			strToCheck := row[:i] + row[i+1:]
			if checked[strToCheck] {
				fmt.Println(strToCheck)
				os.Exit(0)
			}
			checked[strToCheck] = true
		}
	}
}
