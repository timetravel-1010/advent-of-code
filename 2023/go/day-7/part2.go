package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type HandType string

type Hand struct {
	cards []string
	Type  string
	bid   int
}

var (
	FiveOfAKind  string = "FiveOfAKind"
	FourOfAKind  string = "FourOfAKind"
	FullHouse    string = "FullHouse"
	ThreeOfAKind string = "ThreeOfAKind"
	TwoPair      string = "TwoPair"
	OnePair      string = "OnePair"
	HighCard     string = "HighCard"

	labels []string = []string{
		"A", "K", "Q", "T", "9", "8", "7", "6", "5", "4", "3", "2", "J",
	}
	types []string = []string{
		FiveOfAKind, FourOfAKind, FullHouse, ThreeOfAKind, TwoPair, OnePair, HighCard,
	}
)

func main() {
	fmt.Println(getTotalWinnings("input.txt"))
}

func getTotalWinnings(fileName string) int {
	file, err := os.Open(fileName)
	if err != nil {
		panic("Error while opening the file!")
	}
	scanner := bufio.NewScanner(file)
	hands := []Hand{}
	for scanner.Scan() {
		line := scanner.Text()
		ls := strings.Split(line, " ")
		b, err := strconv.Atoi(ls[1])
		if err != nil {
			panic("Error while converting string to integer!")
		}
		c := strings.Split(ls[0], "")
		hands = append(hands, Hand{
			cards: c,
			bid:   b,
			Type:  getBestHandType(c),
		})
	}
	//fmt.Println(hands)
	sort.Slice(hands, func(i, j int) bool {
		return !compareHands(hands[i], hands[j])
	})
	//fmt.Println(hands)
	count := 0
	for i, h := range hands {
		count += (h.bid * (i + 1))
	}
	return count
}

func getStrongest(cards []string) string {
	copy := cards
	sort.Slice(copy, func(i, j int) bool {
		return compare(cards[i], cards[j])
	})
	return cards[0]
}

func replaceAll(arr []string, v string, t string) []string {
	newArr := make([]string, len(arr))
	copy(newArr, arr)
	for i, val := range arr {
		if val == v {
			newArr[i] = t
		}
	}
	return newArr
}

func getBestHandType(c []string) string {
	hands := []Hand{}
	for _, label := range c {
		//strongest := getStrongest(h.cards)
		// Brute force
		if label == "J" {
			continue
		}
		hands = append(hands, Hand{
			cards: c,
			Type:  getHandType(replaceAll(c, "J", label)), // Pending prove only have 1,
		})
	}
	sort.Slice(hands, func(i, j int) bool {
		return compareHands(hands[i], hands[j])
	})
	if len(hands) == 0 {
		return getHandType(c)
	}
	return hands[0].Type
}

func getHandType(c []string) string {
	count := countValues(c)
	same3 := false
	same2 := 0
	diff := 0
	for _, v := range count {
		switch v {
		case 5:
			return FiveOfAKind
		case 4:
			return FourOfAKind
		case 3:
			if same2 == 1 {
				return FullHouse
			}
			same3 = true
		case 2:
			if same3 {
				return FullHouse
			} else if same2 == 1 {
				return TwoPair
			}
			same2++
		case 1:
			diff++
		}
	}
	switch diff {
	case 5:
		return HighCard
	case 3:
		return OnePair
	case 2:
		return ThreeOfAKind
	}
	return ""
}

func countValues(arr []string) map[string]int {
	// Crea un mapa de valores para cada elemento
	dict := make(map[string]int)
	for _, str := range arr {
		dict[str] = dict[str] + 1
	}
	//fmt.Println(dict) // Output: map[apple:2 banana:3 orange:1]
	return dict
}

// Return true if c1 is stronger than c2
func compareCards(c1 string, c2 string) int {
	i1 := indexOf(c1, labels)
	i2 := indexOf(c2, labels)

	if i1 == i2 {
		return 0
	} else if i1 < i2 {
		return 1
	} else {
		return -1
	}
}

func compare(c1 string, c2 string) bool {
	return indexOf(c1, labels) < indexOf(c2, labels)
}

// Returns true if h1 is stronger than h2
func compareHands(h1 Hand, h2 Hand) bool {
	t1 := indexOf(h1.Type, types)
	t2 := indexOf(h2.Type, types)
	if t1 == t2 { // THe same hand type
		for i := 0; i < 5; i++ {
			result := compareCards(h1.cards[i], h2.cards[i])
			if result == 0 { // Are the same cards
				continue
			}
			return result == 1
		}
	}
	return t1 < t2
}

func indexOf(val string, arr []string) int {
	for i, v := range arr {
		if v == val {
			return i
		}
	}
	return -1
}
