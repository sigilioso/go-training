package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type Fetched struct {
	already map[string]bool
	mux     sync.Mutex
	wg      sync.WaitGroup
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, fetched *Fetched) {
	defer fetched.wg.Done()
	fetched.mux.Lock()
	_, alreadyFetched := fetched.already[url]
	fetched.mux.Unlock()
	if alreadyFetched || depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fetched.mux.Lock()
	fetched.already[url] = true
	fetched.mux.Unlock()
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		fetched.wg.Add(1)
		go Crawl(u, depth-1, fetcher, fetched)
	}
	return
}

func main() {
	fetched := Fetched{already: make(map[string]bool)}
	fetched.wg.Add(1)
	Crawl("https://golang.org/", 4, fetcher, &fetched)
	fetched.wg.Wait()
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
