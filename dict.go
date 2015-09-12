package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Trie struct {
	Value    byte
	Valid    bool
	Children []*Trie
}

var Head = &Trie{}

func (t *Trie) GetChild(c byte) *Trie {
	for _, child := range t.Children {
		if child.Value == c {
			return child
		}
	}
	return nil
}

func (t *Trie) AddOrGetChild(c byte) *Trie {
	child := t.GetChild(c)
	if child != nil {
		return child
	}

	newTrie := &Trie{Value: c, Valid: false}
	t.Children = append(t.Children, newTrie)
	return newTrie
}

func (t *Trie) Print(depth int) {
	indent := strings.Repeat(" ", depth)
	fmt.Println(indent, string(t.Value), t.Valid)
	for _, child := range t.Children {
		child.Print(depth + 1)
	}
}

func IsWord(word string) bool {
	current := Head

	for _, c := range word {
		child := current.GetChild(byte(c))
		if child == nil {
			return false
		}
		current = child
	}
	return current.Valid
}

func HasChildren(word string) bool {
	current := Head

	for _, c := range word {
		child := current.GetChild(byte(c))
		if child == nil {
			return false
		}
		current = child
	}
	return len(current.Children) >= 1

}

func Insert(word string) {
	current := Head
	for _, c := range word {
		child := current.GetChild(byte(c))
		if child == nil {
			current = current.AddOrGetChild(byte(c))
		} else {
			current = child
		}
	}
	current.Valid = true
}

func LoadDictionary(filename string) {
	words := ReadLines(filename)
	for _, word := range words {
		Insert(word)
	}
}

func ReadLines(filename string) []string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.ToLower(string(data)), "\n")

	// Remove last line if empty
	if len(lines[len(lines)-1]) == 0 {
		lines = lines[0 : len(lines)-1]
	}
	return lines
}
