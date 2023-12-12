package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	values = map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	numbers = []string{
		"zero",
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}
)

func getCalibrationValue(line string) int {

	// search by word
	firstDigitIdx := -1
	secondDigitIdx := -1
	firstDigit := -1
	secondDigit := -1

	for _, num := range numbers {
		if i := strings.Index(line, num); i != -1 {
			if firstDigitIdx != -1 {
				if i < firstDigitIdx {
					if firstDigitIdx > secondDigitIdx {
						secondDigitIdx = firstDigitIdx 
						secondDigit = (firstDigit / 10)
					}
					firstDigitIdx = i
					firstDigit = (values[num] * 10)
				} else if i > secondDigitIdx {
					secondDigitIdx = i 
					secondDigit = values[num]
				}
			} else {
				firstDigitIdx = i
				firstDigit = (values[num] * 10)
			}
		}

		if i := strings.LastIndex(line, num); i != -1 {
			if firstDigitIdx != -1 {
				if i < firstDigitIdx {
					if firstDigitIdx > secondDigitIdx {
						secondDigitIdx = firstDigitIdx 
						secondDigit = (firstDigit / 10)
					}
					firstDigitIdx = i
					firstDigit = (values[num] * 10)
				} else if i > secondDigitIdx {
					secondDigitIdx = i 
					secondDigit = values[num]
				}
			} else {
				firstDigitIdx = i
				firstDigit = (values[num] * 10)
			}
		}
	}

	// search by character
	chars := strings.Split(line, "")

	for i, c := range chars {
		num, err := strconv.Atoi(c)
		if err == nil {
			if firstDigit == -1 {
				firstDigit = num * 10
			} else {
				if i < firstDigitIdx {
					if firstDigitIdx > secondDigitIdx {
						secondDigitIdx = firstDigitIdx 
						secondDigit = (firstDigit / 10)
					}
					firstDigitIdx = i
					firstDigit = (num * 10)
				} else if i > secondDigitIdx {
					secondDigitIdx = i
					secondDigit = num
				}
			}
		}
	}

	if secondDigit == -1 {
		return firstDigit + (firstDigit/10)
	}
	return firstDigit + secondDigit
}

func readFile(fileName string) (*os.File, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, errors.New("error by reading the file.")
	}
	return file, nil
}

func getSum(file *os.File) int {
	scanner := bufio.NewScanner(file)
	sum := 0
	lines := 0
	for scanner.Scan() {
		sum += getCalibrationValue(scanner.Text())
		lines++
	}
	fmt.Println(lines, "lines read.")
	return sum
}

func main() {
	
	file, err := readFile("input.txt")
	if err != nil {
		panic("an error has occurred!")
	}
	fmt.Println(getSum(file))
	defer file.Close()
	//fmt.Println(getCalibrationValue("vmzcrhtdvnm6fivepkbhcxj"))
	//fmt.Println(getCalibrationValue("6t"))
	// fmt.Println(getCalibrationValue("otwone3one3one"))
}
