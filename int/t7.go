package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	urls := []string{"https://www.google.com/", "https://example.org", "https://example.net"}

	fmt.Println(callRequestsForURLs(urls, 3)[0].StatusCode)
}

// дернуть N урлов с лимитом K (то есть не больше K активных запросов одновременных), сигнатура:
func callRequestsForURLs(urls []string, K int) []*http.Response {
	ch := make(chan int)
	var wg sync.WaitGroup

	result := make([]*http.Response, len(urls))

	for i := 0; i < K; i += 1 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				url, ok := <-ch
				if !ok {
					break
				}
				ans, err := http.Get(urls[url])
				if err != nil {
					continue
				}
				result[url] = ans
			}
		}()
	}

	for i, _ := range urls {
		ch <- i
	}
	close(ch)
	wg.Wait()

	return result
}
