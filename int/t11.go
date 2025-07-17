package main

import (
	"sync"
)

//1. Merge n channels
//2. Если один из входных каналов закрывается,
//то нужно закрыть все остальные каналы

func joinChannels(channels ...chan int) chan int {
	r := make(chan int)

	var wg sync.WaitGroup

	for _, v := range channels {
		wg.Add(1)
		go func(c chan int) {
			defer wg.Done()
			for {
				select {
				case v, ok := <-c:
					if !ok {
						return
					}
					r <- v
				}
			}
		}(v)
	}

	go func() {
		wg.Wait()
		close(r)
	}()

	return r
}

func main() {
	case3()
}
