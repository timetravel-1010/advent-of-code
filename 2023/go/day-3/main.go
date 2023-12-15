package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Number struct {
	value  int
	posx   int
	posy   int
	length int
}

type Symbol struct {
	posx           int
	posy           int
	partNumbers    []int
	partNumbersLen int
}

func getSum(fileName string) int {
	file, err := os.Open(fileName)
	numbers := []Number{}
	symbols := []Symbol{}

	if err != nil {
		panic("an error has occurred!")
	}

	scanner := bufio.NewScanner(file)
	sum := 0
	lines := 0
	y := 0

	for scanner.Scan() {
		str := scanner.Text()

		reNumbers := regexp.MustCompile(`\d+`)
		reSymbols := regexp.MustCompile(`[^0-9.]`)

		submatchNumbers := reNumbers.FindAllString(str, -1)
		submatchNumbersIndex := reNumbers.FindAllStringIndex(str, -1)

		submatchSymbolsIndex := reSymbols.FindAllStringIndex(str, -1)

		for i, num := range submatchNumbers {
			val, err := strconv.Atoi(num)
			pos := submatchNumbersIndex[i]
			if err != nil {
				panic("an error has occurred!")
			}
			numbers = append(numbers, Number{
				value:  val,
				posx:   pos[0],
				posy:   y,
				length: pos[1] - pos[0],
			})
		}
		for _, pos := range submatchSymbolsIndex {
			symbols = append(symbols, Symbol{
				posx:           pos[0],
				posy:           y,
				partNumbers:    []int{},
				partNumbersLen: 0,
			})
		}
		y++
		lines++
	}
	for _, num := range numbers {
		for i := range symbols {
			// verify if the number is on the same line, on the line above or on the line below the symbol
			// In the same line
			if num.posy == symbols[i].posy {
				if num.posx == (symbols[i].posx-1) || num.posx == (symbols[i].posx+1) {
					symbols[i].partNumbers = append(symbols[i].partNumbers, num.value)
					symbols[i].partNumbersLen += 1
				} else if (num.posx + num.length - 1) == (symbols[i].posx - 1) {
					symbols[i].partNumbers = append(symbols[i].partNumbers, num.value)
					symbols[i].partNumbersLen += 1
				}
			}
			// Above a symbols[i]bol
			if num.posy == (symbols[i].posy-1) || num.posy == (symbols[i].posy+1) {
				if num.posx == (symbols[i].posx-1) || num.posx == (symbols[i].posx+1) {
					symbols[i].partNumbers = append(symbols[i].partNumbers, num.value)
					symbols[i].partNumbersLen += 1
				} else if (num.posx + num.length - 1) == (symbols[i].posx - 1) {
					symbols[i].partNumbers = append(symbols[i].partNumbers, num.value)
					symbols[i].partNumbersLen += 1
				} else if (symbols[i].posx >= num.posx) && (symbols[i].posx <= (num.posx + num.length - 1)) {
					symbols[i].partNumbers = append(symbols[i].partNumbers, num.value)
					symbols[i].partNumbersLen += 1
				}
			}
		}
	}
	for _, s := range symbols {
		if s.partNumbersLen == 2 {
			sum += s.partNumbers[0] * s.partNumbers[1]
		}
	}
	return sum
}

func main() {
	fmt.Println(getSum("./input.txt"))
}
