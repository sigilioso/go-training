package main

import (
	"fmt"
	"strings"
	//"golang.org/x/tour/wc"
)

// WordCount function
func WordCount(s string) map[string]int {
	counter := make(map[string]int)
	words := strings.Split(s, " ")
	for _, word := range words {
		_, already := counter[word]
		if already {
			counter[word]++
		} else {
			counter[word] = 1
		}
	}
	return counter
}

func main() {
	fmt.Println(WordCount("I ate a donut. Then I ate another donut."))
	//wc.Test(WordCount)
}
