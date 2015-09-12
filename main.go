package main

import (
	"fmt"
	"sort"
	"time"
)

const (
	BoardSize = 4
)

func main() {
	LoadDictionary("dict.dict")

	board := NewBoard("board.board")

	horizon := []Path{}

	// seed with each starting path
	for x := 0; x < BoardSize; x++ {
		for y := 0; y < BoardSize; y++ {
			currentPath := Path{board.At(x, y)}
			horizon = append(horizon, currentPath)
		}
	}
	words := []string{}

	for len(horizon) != 0 {
		fmt.Println("Size of horizon", len(horizon))

		currentPath := horizon[0]
		horizon = horizon[1:]

		possiblePaths := currentPath.Explore(board)
		fmt.Println(currentPath.Word(), len(possiblePaths))

		for _, possiblePath := range possiblePaths {
			fmt.Println("\t", possiblePath.Word(), []byte(possiblePath.Word()))

			if IsWord(possiblePath.Word()) {
				fmt.Println(words)

				if len(possiblePath) >= 5 {
					fmt.Println()
					fmt.Println(possiblePath.Word())
					fmt.Println()

					board.PrintPath(possiblePath)
					time.Sleep(time.Second * 3)
				}
				words = AppendIfUnique(words, possiblePath.Word())
			}

			if !HasChildren(possiblePath.Word()) {
				continue
			}

			horizon = append(horizon, possiblePath)
		}
	}

	sort.Sort(ByLength(words))
	fmt.Println(words)
}

func AppendIfUnique(words []string, word string) []string {
	for _, w := range words {
		if w == word {
			return words
		}
	}
	return append(words, word)
}

type ByLength []string

func (s ByLength) Len() int {
	return len(s)
}
func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}
