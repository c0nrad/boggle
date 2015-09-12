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
	validPaths := []Path{}

	for len(horizon) != 0 {

		currentPath := horizon[0]
		horizon = horizon[1:]

		possiblePaths := currentPath.Explore(board)

		for _, possiblePath := range possiblePaths {

			if IsWord(possiblePath.Word()) {
				validPaths = AppendIfUnique(validPaths, possiblePath)
			}

			if !HasChildren(possiblePath.Word()) {
				continue
			}

			horizon = append(horizon, possiblePath)
		}
	}

	sort.Sort(Paths(validPaths))

	for _, validPath := range validPaths {
		fmt.Println()
		fmt.Println(validPath.Word())
		fmt.Println()
		board.PrintPath(validPath)
		time.Sleep(time.Second * 3)

	}
}

func AppendIfUnique(paths []Path, path Path) []Path {
	for _, p := range paths {

		if p.Word() == path.Word() {
			return paths
		}
	}
	return append(paths, path)
}
