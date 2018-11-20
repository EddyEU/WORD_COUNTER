package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	dataBin, err := ioutil.ReadFile("words.txt")
	if err != nil {
		fmt.Print(err)
	}
	textFile := string(dataBin)
	wc := getWordsCountMapOfFile(textFile)
	w := getStringSortWordsOfWordsCountMap(wc)
	textFile = ""
	for i, word := range w {
		fmt.Println(word, wc[word])
		result := strconv.Itoa(wc[word])
		ini := strconv.Itoa(i + 1)
		textFile += "[" + ini + "] " + word + " : " + result + "\n"
	}
	writeToFile(textFile)
}

// getWordsCountMapOfFile : here
func getWordsCountMapOfFile(textFile string) map[string]int {
	lowerText := strings.ToLower(textFile)

	lowerTextClean1 := strings.Replace(lowerText, "`", "", -1)
	lowerTextClean2 := strings.Replace(lowerTextClean1, "'", "", -1)
	lowerTextClean3 := strings.Replace(lowerTextClean2, ",", "", -1)
	lowerTextClean4 := strings.Replace(lowerTextClean3, ".", "", -1)
	lowerTextClean5 := strings.Replace(lowerTextClean4, " - ", "", -1)
	lowerTextClean6 := strings.Replace(lowerTextClean5, " -", "", -1)
	lowerTextCleanFull := strings.Replace(lowerTextClean6, "- ", "", -1)

	//reg, err := regexp.Compile("[^a-zA-Z ]+$") //516 words  456
	reg, err := regexp.Compile("[^a-zA-z0-9]+$") //516 words
	//reg, err := regexp.Compile("[^a-zA-Z]+")//440 words
	//reg, err := regexp.Compile("[^a-zA-Z0-9]+") //440 words, output-ee.txt'

	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	onlyAlphaText := reg.ReplaceAllString(lowerTextCleanFull, " ")
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

func writeToFile(textToFile string) {
	file, err := os.Create("result.txt")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	fmt.Fprintf(file, textToFile)
}
