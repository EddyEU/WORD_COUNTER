package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strings"
)

// WordsCount : describe what this function does
func WordsCount(text string) map[string]int {
	lowerText := strings.ToLower(text)

	reg, err := regexp.Compile("[^a-zA-Z]+")
	if err != nil {
		fmt.Print(err)
	}
	onlyAlphaText := reg.ReplaceAllString(lowerText, " ")
	words := strings.Fields(onlyAlphaText)

	wordsCountMap := make(map[string]int)
	for _, word := range words {
		wordsCountMap[word]++
	}
	return wordsCountMap
}

// SortWords : here you tell us what SortWords is
func SortWords(wordsCountMap map[string]int) []string {
	var words []string
	for word := range wordsCountMap {
		words = append(words, word)
	}
	sort.Strings(words)
	return words
}

func main() {
	dataBin, err := ioutil.ReadFile("words.txt")
	if err != nil {
		fmt.Print(err)
	}
	textFile := string(dataBin)
	wc := WordsCount(textFile)
	w := SortWords(wc)
	for _, word := range w {
		fmt.Println(word, wc[word])
	}
}
