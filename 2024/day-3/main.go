package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func solutionPart1(fileName string) int {
	file, err := os.Open(fileName)
	if err != nil {
		panic("error opening the file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0

	for scanner.Scan() {
		re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
		matches := re.FindAllString(scanner.Text(), -1)
		for _, m := range matches {
			p := strings.TrimRight(strings.TrimLeft(strings.TrimLeft(m, "mul"), "("), ")")
			values := strings.Split(p, ",")
			sum += castInt(values[0]) * castInt(values[1])
		}
	}
	return sum
}

func castInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		panic("error casting string to int")
	}
	return num
}

type MulResult struct {
	DoOrDont string
	Muls     []string
}

func solutionPart2(fileName string) int {
	file, err := os.Open(fileName)
	if err != nil {
		panic("error opening the file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	var input string
	for scanner.Scan() {
		input += scanner.Text()
	}

	pattern := `mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`
	re := regexp.MustCompile(pattern)

	matches := re.FindAllString(input, -1)
	isAfterDo := true

	for _, m := range matches {
		if m == "do()" {
			isAfterDo = true
		} else if m == "don't()" {
			isAfterDo = false
		} else if isAfterDo {
			values := strings.Split(m[4:len(m)-1], ",")
			sum += castInt(values[0]) * castInt(values[1])
		}
	}
	return sum
}

func main() {
	fmt.Println(solutionPart2("test2.txt"))
}
