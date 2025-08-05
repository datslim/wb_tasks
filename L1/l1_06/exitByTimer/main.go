package main

import (
	"fmt"
	"sync"
	"time"
)

// пример выхода из горутины используя таймер
func exitByTimer(wg *sync.WaitGroup) {
	defer wg.Done()
	timer := time.NewTimer(5 * time.Second) // создаем таймер на 5 секунд
	for {
		select {
		case <-timer.C: // когда таймер истекает
			fmt.Println("Время истекло. Выходим из горутины.")
			return // выходим из горутины (вызываем wg.Done())
		default:
			fmt.Println("Прошла секунда.") // печатаем сообщение
			time.Sleep(1 * time.Second)    // ожидаем секунду
		}
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go exitByTimer(&wg) // запускаем функцию в горутине
	wg.Wait()
}
