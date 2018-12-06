/**
 * Day 4: Repose Record
 * https://adventofcode.com/2018/day/4
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
)

type guard struct {
	id                  int
	totalMinutes        int
	lastMinuteAsleep    int
	asleepCountByMinute map[int]int
}

func main() {
	lines := readLines()
	sort.Strings(lines)

	guards := make(map[int]*guard)
	var currentGuard *guard
	for _, line := range lines {
		r, _ := regexp.Compile(`^\[\d{4}-\d{2}-\d{2} \d{2}:(\d{2})\] (.+)$`)
		matches := r.FindStringSubmatch(line)
		minuteStr, event := matches[1], matches[2]
		minute, _ := strconv.Atoi(minuteStr)
		if event == "wakes up" {
			currentGuard.wakeUp(minute)
		} else if event == "falls asleep" {
			currentGuard.fallAsleep(minute)
		} else {
			r, _ := regexp.Compile(`^Guard #(\d+) begins shift$`)
			matches := r.FindStringSubmatch(event)
			guardID, _ := strconv.Atoi(matches[1])
			var found bool
			currentGuard, found = guards[guardID]
			if !found {
				currentGuard = &guard{
					id:                  guardID,
					asleepCountByMinute: make(map[int]int),
				}
				guards[guardID] = currentGuard
			}
		}
	}
	mostAsleepGuard := findMostAsleepGuard(guards)
	mostAsleepMinute := keyWithMaxValue(mostAsleepGuard.asleepCountByMinute)
	fmt.Println(mostAsleepGuard.id, mostAsleepMinute, mostAsleepGuard.id*mostAsleepMinute)
}

func (g *guard) wakeUp(minute int) {
	g.totalMinutes += minute - g.lastMinuteAsleep
	for i := g.lastMinuteAsleep; i < minute; i++ {
		g.asleepCountByMinute[i]++
	}
}

func (g *guard) fallAsleep(minute int) {
	g.lastMinuteAsleep = minute
}

func findMostAsleepGuard(guards map[int]*guard) *guard {
	var mostAsleep *guard
	for _, guard := range guards {
		if mostAsleep == nil || guard.totalMinutes > mostAsleep.totalMinutes {
			mostAsleep = guard
		}
	}
	return mostAsleep
}

func keyWithMaxValue(m map[int]int) int {
	res := 0
	for key, val := range m {
		if res == 0 || val > m[res] {
			res = key
		}
	}
	return res
}

func readLines() (lines []string) {
	lines = make([]string, 0)
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("The input file could not be read")
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return
}
