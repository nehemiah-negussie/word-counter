package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

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

func main() {
	args := os.Args[1:]

	if (len(args) < 1) {
		fmt.Println("Missing text filename")
		return
	}

	dat, err := os.ReadFile(args[0])
	check(err)

	txt_file := (string(dat))

	txt_file = cleanString(txt_file)

	words := strings.Fields(txt_file)

	fmt.Println(words[:5])

	fmt.Println(wordCount(words))

}