package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strings"
)

// getWordsCountMapOfFile : here
func getWordsCountMapOfFile(textFile string) map[string]int {
	lowerText := strings.ToLower(textFile)

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

// getStringSortWordsOfWordsCountMap : here
func getStringSortWordsOfWordsCountMap(wordsCountMap map[string]int) []string {
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
	wc := getWordsCountMapOfFile(textFile)
	w := getStringSortWordsOfWordsCountMap(wc)
	for _, word := range w {
		fmt.Println(word, wc[word])
	}
}
