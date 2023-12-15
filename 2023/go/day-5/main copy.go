package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func getInput2(fileName string) ([]int, [][]MapStruct) {
	file, err := os.Open(fileName)
	if err != nil {
		panic("Error while opening the file!")
	}
	scanner := bufio.NewScanner(file)
	var seeds []int
	scanner.Scan()
	seeds = MapToInt(strings.Fields(scanner.Text())[1:], func(x string) int {
		n, _ := strconv.Atoi(x)
		return n
	})

	maps := make([][]MapStruct, 7)
	isMap := false
	i := 0
	mp := MapStruct{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if isMap {
				i++
			}
			isMap = false
			continue
		}
		if !isMap {
			isMap = regexp.MustCompile(`\bmap\b`).MatchString(line)
			continue
		}
		values := MapToInt(strings.Fields(line), func(x string) int {
			n, _ := strconv.Atoi(x)
			return n
		})
		mp = MapStruct{
			destinationRangeStart: values[0],
			sourceRangeStart:      values[1],
			rangeLength:           values[2],
		}
		maps[i] = append(maps[i], mp)
	}
	defer file.Close()
	return seeds, maps
}

func getLowestLocation2(fileName string) int {
	seeds, maps := getInput(fileName)
	lowestLocation := math.MaxInt

	aux := seeds
	next := aux
	for _, mp := range maps {
		for i := 0; i < len(aux); i++ {
			next[i] = Convert(aux[i], mp)
		}
		aux = next
	}
	for _, v := range aux {
		if v < lowestLocation {
			lowestLocation = v
		}
	}
	return lowestLocation
}

func Convert2(v1 int, mps []MapStruct) int {
	for _, mp := range mps {
		if v1 < (mp.sourceRangeStart+mp.rangeLength) && v1 >= mp.sourceRangeStart {
			acc := mp.destinationRangeStart
			for i := mp.sourceRangeStart; i < mp.sourceRangeStart+mp.rangeLength; i++ {
				if v1 == i {
					return acc
				}
				acc++
			}
		}
	}
	return v1
}

func MapToInt2(arr []string, f func(string) int) []int {
	newArr := make([]int, len(arr))
	for i, v := range arr {
		newArr[i] = f(v)
	}
	return newArr
}

func main2() {
	fmt.Println(getLowestLocation("input.txt"))
}
