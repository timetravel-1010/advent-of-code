package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)


func getSum2(fileName string) int {
	file, err := os.Open(fileName)

	if err != nil {
		panic("Error while opening the txt file!")
	}

	scanner := bufio.NewScanner(file)
	cards := []Card{}
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		l := strings.Split(line, ":")
		regex := `Card\s+(\d+)`
		re := regexp.MustCompile(regex)

		// Encuentra el número en el string
		match := re.FindStringSubmatch(l[0])

		num, err := strconv.Atoi(match[1])
		if err != nil {
			panic("Error while converting string to int:" + match[0])
		}
		l2 := strings.Split(l[1], "|")
		regex2 := `(\d+)`
		re2 := regexp.MustCompile(regex2)

		// Encuentra todos los números en la primera parte
		matches1 := re2.FindAllString(strings.TrimSpace(l2[0]), -1)

		// Encuentra todos los números en la segunda parte
		matches2 := re2.FindAllString(strings.TrimSpace(l2[1]), -1)

		cards = append(cards, Card{
			number:         num,
			winningNumbers: matches1, //strings.Split(l2[0], " "),
			MyNumbers:      matches2, //strings.Split(l2[1], " "),
		})
		i++
	}
	fmt.Println(cards)

	sum := 0
	for i, c := range cards {
		for _, num := range c.MyNumbers {
			for _, wNum := range c.winningNumbers {
				if num == wNum {
					if cards[i].worth == 0 {
						cards[i].worth = 1
					} else {
						cards[i].worth *= 2
					}
					fmt.Println("entra con:", num, " winnum:", wNum)
					break
				}
			}
		}
		sum += cards[i].worth
	}
	fmt.Println(cards)
	return sum
}

func main2() {
	fmt.Println(getSum("input.txt"))
}
