package main

import (
	"fmt"

	"github.com/mgutz/ansi"
)

type Board struct {
	Spots [][]Spot
}

func NewBoard(filename string) Board {
	lines := ReadLines(filename)

	b := Board{}
	b.Spots = make([][]Spot, BoardSize)
	for x := 0; x < BoardSize; x++ {
		b.Spots[x] = make([]Spot, BoardSize)
	}

	for y := 0; y < BoardSize; y++ {
		for x := 0; x < BoardSize; x++ {
			c := lines[y][x]
			b.Spots[y][x].C = c
			b.Spots[y][x].X = x
			b.Spots[y][x].Y = y
		}
	}

	return b
}

func (b *Board) PrintPath(p Path) {
	out := ""
	for y := 0; y < BoardSize; y++ {
		for x := 0; x < BoardSize; x++ {
			if p[0].X == x && p[0].Y == y {
				out += ansi.Color(string(b.At(x, y).C), "green")
			} else if p[len(p)-1].X == x && p[len(p)-1].Y == y {
				out += ansi.Color(string(b.At(x, y).C), "red")
			} else if p.SeenSpot(Spot{x, y, 0}) {
				out += ansi.Color(string(b.At(x, y).C), "blue")
			} else {
				out += string(b.At(x, y).C)
			}
		}
		out += "\n"
	}
	fmt.Println(out)
}

type Spot struct {
	X, Y int
	C    byte
}

var NullSpot = Spot{}

func (b *Board) At(x, y int) Spot {
	if x < 0 || y < 0 || y >= BoardSize || x >= BoardSize {
		return NullSpot
	} else {
		return b.Spots[y][x]
	}
}

func (b *Board) GetAdjacent(x, y int) []Spot {
	out := []Spot{}
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			spot := b.At(dx+x, dy+y)
			if spot != NullSpot {
				out = append(out, spot)
			}
		}
	}
	return out
}

type Path []Spot

func (p Path) Word() string {
	out := []byte{}
	for _, s := range p {
		out = append(out, s.C)
	}

	return string(out)
}

func (p Path) SeenSpot(s Spot) bool {
	for _, spot := range p {
		if spot.X == s.X && spot.Y == s.Y {
			return true
		}
	}
	return false
}

func (p Path) Clone() Path {
	out := make([]Spot, len(p))
	copy(out, p)
	return out
}

func (p Path) Explore(b Board) []Path {
	out := []Path{}
	last := p[len(p)-1]

	possibleSpots := b.GetAdjacent(last.X, last.Y)
	fmt.Println("possibleSpots", possibleSpots)

	for _, possibleSpot := range possibleSpots {
		if !p.SeenSpot(possibleSpot) {
			newPath := p.Clone()
			newPath = append(newPath, possibleSpot)
			out = append(out, newPath)
		}
	}
	fmt.Println("possibleSpots out", out)
	return out
}
