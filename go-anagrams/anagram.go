package main

import (
	"bufio"
	"log"
	"os"
	"sort"
)

type runeSlice []rune

func (b runeSlice) Len() int           { return len(b) }
func (b runeSlice) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b runeSlice) Less(i, j int) bool { return b[i] < b[j] }

func IsAnagram(str1, str2 string) bool {
	mapChar := make(map[rune]int)
	for _, char := range str1 {
		if _, isPresent := mapChar[char]; isPresent {
			mapChar[char] += 1
		} else {
			mapChar[char] = 1
		}
	}
	for _, char := range str2 {
		if _, isPresent := mapChar[char]; isPresent {
			mapChar[char] -= 1
			if mapChar[char] == 0 {
				delete(mapChar, char)
			}
		} else {
			return false
		}
	}
	return len(mapChar) == 0
}

func IsAnagramBySort(str1, str2 string) bool {
	b1 := runeSlice(str1)
	b2 := runeSlice(str2)
	sort.Sort(b1)
	sort.Sort(b2)
	for i, char := range b1 {
		if char != b2[i] {
			return false
		}
	}
	return true
}

func readFile() []string {
	result := make([]string, 0)
	file, err := os.Open("unixdict.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}
