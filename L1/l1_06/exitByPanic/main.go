package main

import (
	"fmt"
	"sync"
)

// пример выхода из горутины при панике

func exitByPanic(wg *sync.WaitGroup) {
	defer wg.Done()

	defer func() { // выполнится даже при панике потому что использует defer
		if r := recover(); r != nil { // если паника возникла, то ловим ее
			fmt.Printf("Горутина завершилась с паникой: %v\n", r) // печатаем сообщение
		}
	}()

	fmt.Printf("Вызываем панику.\n")
	panic("паника вызвана синтетически")                   // вызываем панику
	fmt.Printf("Это сообщение никогда не напечатается.\n") // не выполнится из-за паники
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go exitByPanic(&wg) // запускаем функцию в горутине
	wg.Wait()
}
