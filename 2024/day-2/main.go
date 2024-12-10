package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solutionPart1() int {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("Error opening file")
	}

	scanner := bufio.NewScanner(file)

	safe := 0
	for scanner.Scan() {
		strLevels := strings.Split(scanner.Text(), " ")

		l1 := castInt(strLevels[0])
		l2 := castInt(strLevels[1])

		increasing := false
		if l2 > l1 && l2-l1 <= 3 {
			increasing = true
		} else if l1 > l2 && l1-l2 <= 3 {
			increasing = false
		} else {
			continue
		}

		prevLevel := l2
		isSafe := true
		for _, s := range strLevels[2:] {
			level := castInt(s)
			if diff := level - prevLevel; level > prevLevel && diff <= 3 && increasing {
				prevLevel = level
			} else if diff := prevLevel - level; prevLevel > level && diff <= 3 && !increasing {
				prevLevel = level
			} else {
				isSafe = false
				break
			}
		}

		if isSafe {
			safe++
		}
	}

	return safe
}

func castInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		panic("Error casting string to int")
	}
	return num
}

func isSafe(i int, levels []int, pos, increasing int, canTolerate bool) bool {
	current := levels[pos]
	next := levels[pos+1]

	if (next > current) && (next-current <= 3) && (increasing+1 <= 2) {
		if pos == len(levels)-2 {
			return true
		}
		return isSafe(i, levels, pos+1, 1, canTolerate)
	} else if (current > next && current-next <= 3) && (increasing%2 == 0) {
		if pos == len(levels)-2 {
			return true
		}
		return isSafe(i, levels, pos+1, 2, canTolerate)
	} else if canTolerate {
		safeWOL1 := isSafe(i, removeAt(levels, pos), 0, 0, false)
		safeWOL2 := isSafe(i, removeAt(levels, pos+1), 0, 0, false)
		safe := safeWOL1 || safeWOL2
		return safe
	}
	return false
}

func solutionPart2() int {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("Error opening file")
	}

	scanner := bufio.NewScanner(file)
	safe := 0
	i := 1
	for scanner.Scan() {
		str := strings.Fields(scanner.Text())

		levels := []int{}
		for _, s := range str {
			levels = append(levels, castInt(s))
		}

		if isSafe(i, levels, 0, 0, true) {
			safe++
		} else {
			l3 := levels[:len(levels)-1]
			s1 := isSafe(i, removeAt(levels, 0), 0, 0, false)
			s2 := isSafe(i, removeAt(levels, 1), 0, 0, false)
			s3 := isSafe(i, l3, 0, 0, false)
			if s1 || s2 || s3 {
				safe++
			}
		}
		i++
	}
	return safe
}

func removeAt[T any](s []T, index int) []T {
	ret := make([]T, 0, len(s)-1)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

// 1 2 7 8 9
func main() {
	fmt.Println("sol1:", solutionPart1())
	fmt.Println("sol2:", solutionPart2())
}
