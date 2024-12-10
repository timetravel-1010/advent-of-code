package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func solutionPart1(fileName string) int {
	sum := 0
	file, err := os.Open(fileName)
	if err != nil {
		panic("error opening the file")
	}

	scanner := bufio.NewScanner(file)

	rules := make(map[string]map[string]bool)
	secondSection := false
	for scanner.Scan() {
		ordered := true
		middle := 0
		l := scanner.Text()
		if l == "" {
			secondSection = true
			continue
		}

		if !secondSection {
			pages := strings.Split(l, "|")
			if _, ok := rules[pages[0]]; !ok {
				rules[pages[0]] = map[string]bool{pages[1]: true}
			} else {
				rules[pages[0]][pages[1]] = true
			}
		} else {
			update := strings.Split(l, ",")
			middle, _ = strconv.Atoi(update[len(update)/2])
			ordered = sort.SliceIsSorted(update, func(i, j int) bool {
				return rules[update[i]][update[j]]
			})
			if ordered {
				sum += middle
			}
		}
	}
	return sum
}

func solutionPart2(fileName string) int {
	sum := 0
	file, err := os.Open(fileName)
	if err != nil {
		panic("error opening the file")
	}

	scanner := bufio.NewScanner(file)
	rules := make(map[string]map[string]bool)
	secondSection := false
	for scanner.Scan() {
		ordered := true
		l := scanner.Text()
		if l == "" {
			secondSection = true
			continue
		}

		if !secondSection {
			pages := strings.Split(l, "|")
			if _, ok := rules[pages[0]]; !ok {
				rules[pages[0]] = map[string]bool{pages[1]: true}
			} else {
				rules[pages[0]][pages[1]] = true
			}
			continue
		}
		update := strings.Split(l, ",")
		ordered = sort.SliceIsSorted(update, func(i, j int) bool {
			return rules[update[i]][update[j]]
		})
		if !ordered {
			sort.SliceStable(update, func(i, j int) bool {
				return rules[update[i]][update[j]]
			})
			middle, _ := strconv.Atoi(update[len(update)/2])
			sum += middle
			continue
		}
	}
	return sum
}

func main() {
	fmt.Printf("sol1: %d\n", solutionPart1("input.txt"))
	fmt.Printf("sol2: %d\n", solutionPart2("input.txt"))
}
