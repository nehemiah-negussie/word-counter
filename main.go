package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var topNWords int

type WordCount struct {
	Word string
	Count int
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func cleanString(text string) string {
	reg := regexp.MustCompile(`[^a-zA-Z\s]`)
	cleaned := reg.ReplaceAllString(text, "")
	return strings.ToLower(cleaned)
}

func wordCount(words []string) map[string]int {
	res := make(map[string]int)
	for _, word := range words {
		count, ok := res[word]
		if ok {
			res[word] = count + 1
		} else {
			res[word] = 1
		}
	}
	return res
}

func displayCount(countMap map[string]int) {
	var counts []WordCount

	for word, count := range countMap {
		counts = append(counts, WordCount{word, count})
	}

	sort.Slice(counts, func(i, j int) bool {
		return counts[i].Count > counts[j].Count
	})

	if topNWords > len(counts) {
		topNWords = len(counts)
	}
	
	fmt.Printf("The top %d words are:\n", topNWords)
	for _, wordCountPair := range counts[:topNWords] {
		fmt.Printf("%v: %d\n", wordCountPair.Word, wordCountPair.Count)
	}
}


func main() {
	args := os.Args[1:]

	if (len(args) < 2) {
		fmt.Println("USAGE: program <filepath> <N>, where N is the top N words")
		return
	}

	dat, err := os.ReadFile(args[0])
	check(err)

	topNWords, err = strconv.Atoi(args[1])
	check(err)

	if topNWords < 0 {
		topNWords = 0
	}

	txt_file := (string(dat))

	txt_file = cleanString(txt_file)

	words := strings.Fields(txt_file)

	displayCount(wordCount(words))

}