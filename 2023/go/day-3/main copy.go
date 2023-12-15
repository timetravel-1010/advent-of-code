package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)


func getSum2(fileName string) int {
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

		//submatchSymbols := reSymbols.FindAllString(str, -1)
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
				posx: pos[0],
				posy: y,
			})
		}
		y++
		lines++
	}
	for _, num := range numbers {
		for _, sym := range symbols {
			// verify if the number is on the same line, on the line above or on the line below the symbol

			// In the same line
			if num.posy == sym.posy {
				if num.posx == (sym.posx-1) || num.posx == (sym.posx+1) {
					sum += num.value
					fmt.Println("val1:", num.value, "sym:", sym.posy)
					break
				} else if (num.posx + num.length - 1) == (sym.posx - 1) {
					sum += num.value
					fmt.Println("val2:", num.value, "sym:", sym.posy)
					break
				}
			}
			// Above a symbol
			if num.posy == (sym.posy - 1) {
				if num.posx == (sym.posx-1) || num.posx == (sym.posx+1) {
					sum += num.value
					fmt.Println("val1:", num.value, "sym:", sym.posy)
					break
				} else if (num.posx + num.length - 1) == (sym.posx - 1) {
					sum += num.value
					fmt.Println("val2:", num.value, "sym:", sym.posy)
					break
				} else if (sym.posx >= num.posx) && (sym.posx <= (num.posx + num.length - 1)) {
					sum += num.value
					fmt.Println("val3:", num.value, "sym:", sym.posy)
					break
				}
				// Upper right diagonal

				// Upper left diagonal

			}
			// Under a symbol
			if num.posy == (sym.posy + 1) {
				// Lower right diagonal

				// Lower left diagonal
				if num.posx == (sym.posx-1) || num.posx == (sym.posx+1) {
					sum += num.value
					fmt.Println("val1:", num.value, "sym:", sym.posy)
					break
				} else if (num.posx + num.length - 1) == (sym.posx - 1) {
					sum += num.value
					fmt.Println("val2:", num.value, "sym:", sym.posy)
					break
				} else if (sym.posx >= num.posx) && (sym.posx <= (num.posx + num.length - 1)) {
					sum += num.value
					fmt.Println("val3:", num.value, "sym:", sym.posy)
					break
				}
			}
		}
	}
	fmt.Println("----------")
	fmt.Println(numbers)
	fmt.Println("===========")
	fmt.Println(symbols)
	fmt.Println(lines, "lines read.")
	return sum
}

func main2() {
	fmt.Println(getSum("input.txt"))
}
