package main

import (
	"sync"
)

//1. Иногда приходят нули. В чем проблема? Исправь ее
//2. Если функция bank_network_call выполняется 5 секунд,
//то за сколько выполнится balance()? Как исправить проблему?
//3. Представим, что bank_network_call возвращает ошибку дополнительно.
//Если хотя бы один вызов завершился с ошибкой, то balance должен вернуть ошибку.

// 2. За 5 / (количество физических потоков) с округлением вверх

func balance() (int, error) {
	x := make(map[int]int, 1)
	var m sync.Mutex

	var wg sync.WaitGroup

	var balanceErr error

	// call bank
	for i := 0; i < 5; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			m.Lock()
			if balanceErr != nil {
				m.Unlock()
				return
			}
			b, err := bank_network_call(i)
			if err != nil {
				balanceErr = err
			}
			x[i] = b
			m.Unlock()
		}()
	}
	wg.Wait()

	// Как-то считается сумма значений в мапе и возвращается
	return sumOfMap, balanceErr
}
