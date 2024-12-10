package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
)

func solutionPart1(fileName string) int {
	file, err := os.ReadFile(fileName)
	if err != nil {
		panic("error reading the file")
	}

	diskMap := string(file)
	free := false
	expanded := []int{}
	id := 0
	for _, c := range diskMap {
		num, _ := strconv.Atoi(string(c))
		for times := 0; times < num; times++ {
			if free {
				expanded = append(expanded, -1)
			} else {
				expanded = append(expanded, id)
			}
		}
		free = !free
		if !free {
			id++
		}
	}

	checksum := 0
	lastId := len(expanded) - 1
	for i := 0; i <= lastId; i++ {
		b := expanded[i]
		if b == -1 {
			if last := expanded[lastId]; last != -1 {
				checksum += i * last
			} else {
				i--
			}
			lastId--
		} else {
			checksum += i * b
		}
	}
	return checksum
}

func solutionPart2(fileName string) int {
	file, err := os.ReadFile(fileName)
	if err != nil {
		panic("error reading the file")
	}

	diskMap := string(file)
	expanded := []int{}
	lens := make(map[int]int)
	id := 0
	pos := 0
	i := 0
	idx := []int{}
	freeSpace := []int{}
	for idd, c := range diskMap {
		num, _ := strconv.Atoi(string(c))
		if string(c) != "\n" {
			lens[i] = num
			if len(idx) > 0 && i == idx[len(idx)-1] {
			} else {
				idx = append(idx, i)
				if idd%2 != 0 && lens[i] > 0 {
					freeSpace = append(freeSpace, i)
				}
			}
		}
		i += num
		for times := 0; times < num; times++ {
			if idd%2 != 0 {
				expanded = append(expanded, -1)
			} else {
				expanded = append(expanded, id)
			}
			pos++
		}
		if idd%2 == 0 {
			id++
		}
	}

	reversed := make([]int, len(idx))
	copy(reversed, idx)

	slices.Reverse(reversed)
	for _, lastIdx := range reversed {
		lastPos := lastIdx
		last := expanded[lastIdx]
		if last == -1 {
			continue
		}
		for pos := 0; pos < len(freeSpace); pos++ {
			i := freeSpace[pos]
			if lens[lastIdx] < lens[i] {
				for j := i; j < i+lens[lastIdx]; j++ {
					expanded[j] = last
					expanded[lastPos] = -1
					lastPos++
				}
				freeSpace[pos] = freeSpace[pos] + lens[lastIdx]
				lens[freeSpace[pos]] = lens[i] - lens[lastIdx]
				lens[i] = lens[lastIdx]
				break
			} else if lens[lastIdx] == lens[i] {
				for j := i; j < i+lens[lastIdx]; j++ {
					expanded[j] = last
					expanded[lastPos] = -1
					lastPos++
				}
				freeSpace = append(freeSpace[:pos], freeSpace[pos+1:]...)
				lens[i] = 0
				break
			}
		}
	}
	checksum := 0

	for i, b := range expanded {
		if b == -1 {
			continue
		}
		checksum += i * b
	}

	return checksum
}

func main() {
	//fmt.Println("sol1:", solutionPart1("test.txt"))
	fmt.Println("sol2:", solutionPart2("input.txt"))
}
