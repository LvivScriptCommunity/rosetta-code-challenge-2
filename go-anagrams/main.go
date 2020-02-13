package main

import (
	"fmt"
	"sort"
)

const maxAnagramCount = 5

func main() {
	words := readFile()
	m := make(map[string][]string)
	for _, word := range words {
		runes := runeSlice(word)
		sort.Sort(runes)
		key := string(runes)
		a := append(m[key], word)
		m[key] = a
	}
	for _, a := range m {
		if len(a) == maxAnagramCount {
			fmt.Printf("%s\n", a)
		}
	}
}
