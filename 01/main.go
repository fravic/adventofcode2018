/**
 * Day 1: Chronal Calibration
 * https://adventofcode.com/2018/day/1
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const maxIterations = 9999

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("The input file could not be read")
	}
	var total int64
	freqs := make(map[int64]bool)
	for i := 0; i < maxIterations; i++ {
		f.Seek(0, 0)
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			textVal := scanner.Text()
			intVal, err := strconv.ParseInt(textVal, 10, 64)
			if err != nil {
				fmt.Printf("Invalid integer encountered: %s\n", textVal)
			} else {
				total += intVal
			}
			if freqs[total] {
				fmt.Println(total)
				os.Exit(0)
				break
			}
			freqs[total] = true
		}
	}
}
