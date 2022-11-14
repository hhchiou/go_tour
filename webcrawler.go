package main

import (
	"fmt"
	"sync"
)

type SafeCache struct {
	mu sync.Mutex
	c  map[string]bool
}

func (c *SafeCache) TestAndSet(key string) bool {
	c.mu.Lock()
	val := c.c[key]
	c.c[key] = true
	defer c.mu.Unlock()
	return val
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, c *SafeCache, ch chan int) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 || c.TestAndSet(url) {
		ch <- 0
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		ch <- 0
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	ch2 := make(chan int)
	for _, u := range urls {
		go Crawl(u, depth-1, fetcher, c, ch2)
	}
	for i := 0; i < len(urls); i++ {
		<-ch2
	}
	ch <- 0
	return
}

func main() {
	c := SafeCache{c: make(map[string]bool)}
	ch := make(chan int)
	go Crawl("https://golang.org/", 4, fetcher, &c, ch)
	<-ch
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
