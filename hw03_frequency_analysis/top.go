package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

type uniqueword struct {
	word  string
	count int
}

type uniquewords struct {
	uniquewords   []uniqueword
	mapuniqueword map[string]int
}

func newuniquewords() uniquewords {
	return uniquewords{make([]uniqueword, 0, 0), make(map[string]int)}
}

func separatebyfields(text string) []string {
	words := strings.Fields(text)
	reprep := regexp.MustCompile("[.,!?]$")
	for i := range words {
		words[i] = strings.ToLower(words[i])
		if reprep.MatchString(words[i]) {
			words[i] = reprep.ReplaceAllString(words[i], "")
		}
	}
	return words
}

func (uw *uniquewords) fillfromslice(words []string) {
	for _, word := range words {
		if word == "-" {
			continue
		}
		i, ok := uw.mapuniqueword[word]
		if ok {
			uw.uniquewords[i].count++
		} else {
			uw.mapuniqueword[word] = len(uw.uniquewords)
			uw.uniquewords = append(uw.uniquewords, uniqueword{word, 1})
		}
	}
}

func (uw *uniquewords) sortbycount() {
	sort.Slice(uw.uniquewords, func(i, j int) bool {
		if uw.uniquewords[i].count == uw.uniquewords[j].count {
			return uw.uniquewords[i].word < uw.uniquewords[j].word
		}
		return uw.uniquewords[i].count > uw.uniquewords[j].count
	})
}

func (uw *uniquewords) getwords(start int, end int) []string {
	if start >= len(uw.uniquewords) {
		return []string{}
	}
	if end >= len(uw.uniquewords) {
		return []string{}
	}
	words := make([]string, 0, end-start+1)
	for i := start; i <= end; i++ {
		words = append(words, uw.uniquewords[i].word)
	}
	return words
}

func Top10(text string) []string {
	uw := newuniquewords()
	words := separatebyfields(text)
	uw.fillfromslice(words)
	uw.sortbycount()
	top := uw.getwords(0, 9)
	return top
}
