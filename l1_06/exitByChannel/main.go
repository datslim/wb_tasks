package main

import (
	"fmt"
	"sync"
	"time"
)

// пример остановки выполнения горутины используя канал уведомления
func exitByChannel(wg *sync.WaitGroup, stopChannel <-chan struct{}) {
	defer wg.Done()
	i := 0 // создаем переменную, которую будем инкреминтировать
	for {
		select {
		case <-stopChannel: // когда поступает значение (или канал закрывается)
			return // выходим -> вызываем wg.Done()
		default:
			fmt.Printf("i равен %d\n", i) // печатаем значение i
			time.Sleep(1 * time.Second)   // ожидаем секунду
			i++                           // инкрементируем i
		}
	}
}

func main() {
	var wg sync.WaitGroup
	stopChannel := make(chan struct{}) // создаем канал типа пустой структуры для передачи сигнала о завершении

	wg.Add(1)
	go exitByChannel(&wg, stopChannel) // запускаем функцию в горутине

	time.Sleep(5 * time.Second) // ждем 5 секунд
	close(stopChannel)          // закрываем канал

	fmt.Printf("Контекст завершился спустя 5 секунд\n")
	wg.Wait()

}
