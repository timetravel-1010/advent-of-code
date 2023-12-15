package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func getWays2(time int, distance int) int {
	ways := []int{}
	speed := 1
	for i := 1; i < time; i++ {
		if speed*(time-i) > distance {
			ways = append(ways, i)
		}
		speed++
	}
	return len(ways)
}

func getMarginError2(fileName string) int {
	file, err := os.Open(fileName)
	if err != nil {
		panic("Error while opening the file!")
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	re := regexp.MustCompile(`[0-9]+`)
	numbers := re.FindAllString(scanner.Text(), -1)
	time := StringToInt(numbers)

	fmt.Println(time)
	scanner.Scan()
	re = regexp.MustCompile(`[0-9]+`)
	numbers = re.FindAllString(scanner.Text(), -1)
	distance := StringToInt(numbers)
	fmt.Println(distance)

	res := 1
	for i := 0; i < len(time); i++ {
		ways := getWays(time[i], distance[i])
		if ways > 0 {
			res *= ways
		}
	}
	return res
}

func StringToInt2(ss []string) []int {
	nums := make([]int, len(ss))
	for i := 0; i < len(ss); i++ {
		n, err := strconv.Atoi(ss[i])
		if err != nil {
			panic("Error while converting the integer.")
		}
		nums[i] = n
	}
	return nums
}

func main2() {
	fmt.Println(getMarginError("input.txt"))
}
