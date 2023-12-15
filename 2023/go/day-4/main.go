package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Card struct {
	number         int
	winningNumbers []string
	MyNumbers      []string
	worth          int
	copies         int
}

func getSum(fileName string) int {
	file, err := os.Open(fileName)

	if err != nil {
		panic("Error while opening the txt file!")
	}
	scanner := bufio.NewScanner(file)
	cards := []Card{}
	num := 1
	for scanner.Scan() {
		line := scanner.Text()
		l := strings.Split(line, ":")

		l2 := strings.Split(l[1], "|")
		regex2 := `(\d+)`
		re2 := regexp.MustCompile(regex2)

		matches1 := re2.FindAllString(strings.TrimSpace(l2[0]), -1)
		matches2 := re2.FindAllString(strings.TrimSpace(l2[1]), -1)

		cards = append(cards, Card{
			number:         num,
			winningNumbers: matches1,
			MyNumbers:      matches2,
		})
		num++
	}
	// Calculate worth
	for idx := 0; idx < len(cards); idx++ {
		for _, num := range cards[idx].MyNumbers {
			for _, wNum := range cards[idx].winningNumbers {
				if num == wNum {
					cards[idx].worth += 1
					break
				}
			}
		}
		for j := 0; j < cards[idx].copies+1; j++ {
			for k := 1; k <= cards[idx].worth; k++ {
				cards[idx+k].copies += 1
			}
		}
	}
	sum := 0
	for i := 0; i < len(cards); i++ {
		for j := 0; j < cards[i].copies+1; j++ {
			for k := 1; k <= cards[i].worth; k++ {
				cards[i+k].copies += 1
			}
		}
	}
	for _, c := range cards {
		sum += c.copies + 1
	}
	return sum
}

func main() {
	fmt.Println(getSum("input.txt"))
}
