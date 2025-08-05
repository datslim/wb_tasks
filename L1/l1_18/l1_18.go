package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// структура безопасного счетчика
type SafeCounter struct {
	counter int
	mu      sync.Mutex
}

// конструктор для безопасного счетчика, инициализируем поле "счетчик" нулем
func NewSafeCounter() *SafeCounter {
	return &SafeCounter{
		counter: 0,
	}
}

// метод для безопасного инкремента
func (c *SafeCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counter++
}

// метод для получения итого счетчика
func (c *SafeCounter) GetCounter() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.counter
}

func main() {
	var wg sync.WaitGroup
	safeCounter := NewSafeCounter()

	numberOfGoroutines := rand.Intn(100) // выбираем случайное количество горутин для теста
	wg.Add(numberOfGoroutines)
	for range numberOfGoroutines {
		go func() {
			defer wg.Done()
			safeCounter.Increment() // инкрементируем наш счетчик внутри горутины
		}()
	}

	wg.Wait()
	counterAfterConcurrentIncrement := safeCounter.GetCounter() // получаем итоговое значение счетчика

	// сравниваем полученный результат с количеством запущенных горутин
	if counterAfterConcurrentIncrement == numberOfGoroutines {
		fmt.Printf("После инкремента c помощью %d горутин, значение равно: %d\nРезультат корректен.\n", numberOfGoroutines, counterAfterConcurrentIncrement)
	} else {
		fmt.Printf("Ошибка: ожидалось %d, получено %d\n", numberOfGoroutines, counterAfterConcurrentIncrement)
	}
}
