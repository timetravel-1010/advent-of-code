package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var values = map[string]int{
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

func getCalibrationValue(line string) int {
	chars := strings.Split(line, "")
	firstDigit := -1
	secondDigit := -1

	for _, c := range chars {
		num, err := strconv.Atoi(c)
		if err == nil {
			if firstDigit == -1 {
				firstDigit = num * 10
			} else {
				secondDigit = num
			}
		}
	}

	if secondDigit == -1 {
		return firstDigit + firstDigit/10
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
	for scanner.Scan() {
		sum += getCalibrationValue(scanner.Text())
	}
	return sum
}

func main() {
	file, err := readFile("input.txt")
	if err != nil {
		panic("an error has occurred!")
	}
	fmt.Println(getSum(file))

	cadena := "Hola,¿comoestas?"
	palabra := "estas"

	posicion := strings.Index(cadena, palabra)
	if posicion != -1 {
		fmt.Println("La palabra", palabra, "se encuentra en la posición", posicion, "en la cadena.")
	} else {
		fmt.Println("La palabra", palabra, "no se encuentra en la cadena.")
	}
}
