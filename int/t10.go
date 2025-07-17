package main

import (
	"context"
	"sync"
)

//1. Merge n channels
//2. Если один из входных каналов закрывается,
//то нужно закрыть все остальные каналы

func case3(channels ...chan int) chan int {
	r := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup

	for _, v := range channels {
		wg.Add(1)
		go func(c chan int) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					//close(c) нужно закрыть все остальные каналы??
					return
				case v, ok := <-c:
					if !ok {
						cancel()
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
