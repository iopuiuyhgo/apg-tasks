package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	urls := []string{"https://www.wildberries.ru/", "https://www.google.com/"}
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			ans, err := http.Get(url)
			if err != nil {
				fmt.Println(url, err)
			} else {
				fmt.Println(url, ans.StatusCode)
			}
		}(url)
	}

	wg.Wait()
}
