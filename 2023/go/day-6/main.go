package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func getWays(time int, distance int) int {
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

func getMarginError(fileName string) int {
	file, err := os.Open(fileName)
	if err != nil {
		panic("Error while opening the file!")
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	re := regexp.MustCompile(`[0-9]+`)
	numbers := re.FindAllString(scanner.Text(), -1)
	s1 := ""
	for _, n := range numbers {
		s1 += n
	}
	time, err := strconv.Atoi(s1)
	fmt.Println(time)
	scanner.Scan()
	re = regexp.MustCompile(`[0-9]+`)
	numbers = re.FindAllString(scanner.Text(), -1)
	s2 := ""
	for _, n := range numbers {
		s2 += n
	}
	distance, err := strconv.Atoi(s2)
	fmt.Println(distance)

	return getWays(time, distance)
}

func StringToInt(ss []string) []int {
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

func main() {
	fmt.Println(getMarginError("input.txt"))
}
