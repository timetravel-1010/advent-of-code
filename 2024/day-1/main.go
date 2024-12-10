package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func solutionPart1() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("error opening the input file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	totalDistance := 0
	leftL := []int{}
	rightL := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Fields(line)
		leftL = append(leftL, castInt(values[0]))
		rightL = append(rightL, castInt(values[1]))

	}
	sort.SliceStable(leftL, func(i, j int) bool { return leftL[i] < leftL[j] })
	sort.SliceStable(rightL, func(i, j int) bool { return rightL[i] < rightL[j] })

	for i := 0; i < len(leftL); i++ {
		totalDistance += int(math.Abs(float64(leftL[i]) - float64(rightL[i])))
	}

	fmt.Println(totalDistance)
}

func castInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		panic("error casting string to int")
	}

	return num
}

func countOccurrences[T comparable](slice []T, element T) int {
	count := 0
	for _, v := range slice {
		if v == element {
			count++
		}
	}
	return count
}

func SolutionPart2() int {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("error opening the input file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	leftL := []int{}
	rightL := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Fields(line)
		leftL = append(leftL, castInt(values[0]))
		rightL = append(rightL, castInt(values[1]))
	}

	simScore := 0
	for _, v := range leftL {
		s := countOccurrences(rightL, v)
		simScore += (v * s)
	}

	return simScore
}

func main() {
	fmt.Println(SolutionPart2())
}
