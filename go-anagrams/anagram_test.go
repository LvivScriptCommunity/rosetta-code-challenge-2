package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestSortAnagram_lviv_vivl(t *testing.T) {
	if !IsAnagramBySort("lviv", "vivl") {
		t.Error("Should be not anagram")
	}
	if !IsAnagram("lviv", "vivl") {
		t.Error("Should be not anagram")
	}
}

func TestSortAnagram_lviV_vvvl(t *testing.T) {
	if IsAnagramBySort("lviV", "vvvl") {
		t.Error("Should be not anagram")
	}
	if IsAnagram("lviV", "vvvl") {
		t.Error("Should be not anagram")
	}
}

//go test -v -bench=. -benchmem -benchtime=5s
func BenchmarkAnagram(b *testing.B) {
	b.StopTimer()
	randomStrings := GetRandomStrings()
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		for i := 0; i < len(randomStrings)-1; i++ {
			IsAnagram(randomStrings[i], randomStrings[i+1])
		}
	}
}

func BenchmarkAnagramSort(b *testing.B) {
	b.StopTimer()
	randomStrings := GetRandomStrings()
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		for i := 0; i < len(randomStrings)-1; i++ {
			IsAnagramBySort(randomStrings[i], randomStrings[i+1])
		}
	}
}

func GetRandomStrings() []string {
	randomStrings := make([]string, 0)
	for i := 1; i < 1000; i++ {
		randomStrings = append(randomStrings, fmt.Sprintf("%08d", rand.Intn(100000000)))
	}
	return randomStrings
}
