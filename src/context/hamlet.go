package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	url     = "http://www.textfiles.com/etext/AUTHORS/SHAKESPEARE/shakespeare-hamlet-25.txt"
	timeout = 5
)

func hamletLines() ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Received a non-success message %v", resp.StatusCode)
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(content), "\n"), nil
}

func search(ctx context.Context, cancel context.CancelFunc, matches chan string, word string, lines []string) {
	for i, line := range lines {
		select {
		case <-ctx.Done():
			return
		default:
			if strings.Index(line, word) != -1 {
				matches <- fmt.Sprintf("%v: %v", i, line)
			}
		}
	}
	fmt.Println("All lines checked, but not found enough matches...")
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Word parameter is required")
	}
	word := os.Args[1]
	lines, err := hamletLines()
	if err != nil {
		log.Fatal(err)
	}
	matches := make(chan string)
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	fmt.Println("Searching...")
	go search(ctx, cancel, matches, word, lines)
Loop:
	for i := 0; i < 50; i++ {
		select {
		case match := <-matches:
			fmt.Println(match)
		case <-ctx.Done():
			cancel()
			break Loop
		}
	}
	cancel()
}
