package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

const (
	NoDiagonal = iota
	DiagonalLeft
	DiagonalRight
	DiagonalUndefined
)

type Word struct {
	pos      int
	cursor   int
	next     rune
	word     string
	diagonal int
	done     bool
	level    int
}

func (w *Word) String() string {
	return fmt.Sprintf("&{pos:%d, cursor:%d, next:%s, word:%s, diag: %d, done:%t, level: %d}", w.pos, w.cursor, string(w.next), w.word, w.diagonal, w.done, w.level)
}

func NewWord(pos int, w string, d int) *Word {
	return &Word{
		pos:      pos,
		cursor:   1,
		next:     rune(w[1]),
		word:     w,
		diagonal: d,
	}
}

func (w *Word) isNext(c rune, newPos int) bool {
	if c == w.next {
		if c == rune(w.word[len(w.word)-1]) {
			w.done = true
			return true
		}
		w.level++
		w.cursor++
		w.next = rune(w.word[w.cursor])
		if w.diagonal != NoDiagonal {
			if w.pos > newPos {
				w.diagonal = DiagonalLeft
			} else {
				w.diagonal = DiagonalRight
			}
			w.pos = newPos
		}
		return true
	}
	return false
}

func solutionPart1(fileName string) int {
	file, err := os.Open(fileName)
	if err != nil {
		panic("error opening the file")
	}

	scanner := bufio.NewScanner(file)
	occs := 0
	vWords := map[int]*Word{}
	dWords := []*Word{}
	level := 0
	for scanner.Scan() {
		l := scanner.Text()

		// Same line words
		occs += strings.Count(l, "XMAS")
		occs += strings.Count(l, "SAMX")

		for _, c := range "XMAS" {
			idx := strings.IndexRune(l, c)
			pad := idx

			for idx != -1 {
				// Vertical words
				if word := vWords[pad]; word == nil {
					if c == 'X' {
						vWords[pad] = NewWord(pad, "XMAS", NoDiagonal)
					} else if c == 'S' {
						vWords[pad] = NewWord(pad, "SAMX", NoDiagonal)
					}
				} else if vWords[pad].isNext(c, pad) {
					if vWords[pad].done {
						occs++
						if c == 'X' {
							vWords[pad] = NewWord(pad, "XMAS", NoDiagonal)
						} else if c == 'S' {
							vWords[pad] = NewWord(pad, "SAMX", NoDiagonal)
						}
					}
				} else {
					delete(vWords, pad)
					if c == 'X' {
						vWords[pad] = NewWord(pad, "XMAS", NoDiagonal)
					} else if c == 'S' {
						vWords[pad] = NewWord(pad, "SAMX", NoDiagonal)
					}
				}
				// Diagonal
				toRemove := []int{}
				for i := 0; i < len(dWords); i++ {
					dw := dWords[i]
					if dw.level+1 != level {
						continue
					}
					if dw.diagonal == DiagonalUndefined {
						if !dw.isNext(c, pad) {
							toRemove = append(toRemove, i)
						}
					} else {
						left := dw.diagonal == DiagonalLeft && dw.pos-1 == pad
						right := dw.diagonal == DiagonalRight && dw.pos+1 == pad
						if left || right {
							if dw.isNext(c, pad) {
								if dw.done {
									occs++
									toRemove = append(toRemove, i)
								}
							} else {
								toRemove = append(toRemove, i)
							}
						}
					}
				}
				dWords = deleteElements(dWords, toRemove)
				if c == 'X' || c == 'S' {
					for i := 0; i < 2; i++ {
						name := "XMAS"
						if c == 'S' {
							name = "SAMX"
						}
						diagonal := DiagonalLeft
						if i == 1 {
							diagonal = DiagonalRight
						}
						newW := NewWord(pad, name, diagonal)
						newW.level = level
						dWords = append(dWords, newW)
					}
				}
				idx = strings.IndexRune(l[pad+1:], c)
				pad += idx + 1
			}
		}
		level++
	}
	return occs
}

func solutionPart2(fileName string) int {
	file, err := os.Open(fileName)
	if err != nil {
		panic("error opening the file")
	}

	scanner := bufio.NewScanner(file)
	occs := 0
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var groups [][]string
	for i := 0; i < len(lines)-2; i++ {
		groups = append(groups, lines[i:i+3])
	}

	for _, lines := range groups {
		middle := lines[1]
		aIdx := strings.Index(middle, "A")
		pad := aIdx
		for aIdx != -1 && pad != len(middle)-1 {
			if aIdx == 0 && pad == 0 {
				aIdx = strings.Index(middle[pad+1:], "A")
				pad += aIdx + 1
				continue
			}
			up := lines[0]
			down := lines[2]
			upLeft := string(up[pad-1])
			upRight := string(up[pad+1])
			bottomLeft := string(down[pad-1])
			bottomRight := string(down[pad+1])
			if (upLeft == "M" && upRight == "S" && bottomLeft == "M" && bottomRight == "S") ||
				(upLeft == "S" && upRight == "M" && bottomLeft == "S" && bottomRight == "M") ||
				(upLeft == "M" && upRight == "M" && bottomLeft == "S" && bottomRight == "S") ||
				(upLeft == "S" && upRight == "S" && bottomLeft == "M" && bottomRight == "M") {
				occs++
			}
			aIdx = strings.Index(middle[pad+1:], "A")
			pad += aIdx + 1
		}
	}
	return occs
}
func deleteElements[T any](arr []T, indexes []int) []T {
	sort.Sort(sort.Reverse(sort.IntSlice(indexes)))

	indexMap := make(map[int]struct{}, len(indexes))
	for _, idx := range indexes {
		indexMap[idx] = struct{}{}
	}

	result := []T{}
	for i, val := range arr {
		if _, found := indexMap[i]; !found {
			result = append(result, val)
		}
	}

	return result
}

func main() {
	fmt.Println(solutionPart2("input.txt"))
}
