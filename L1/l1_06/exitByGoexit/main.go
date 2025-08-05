package main

import (
	"fmt"
	"runtime"
	"sync"
)

// пример выхода из горутины используя runtime.Goexit()

func exitByGoexit(wg *sync.WaitGroup) {
	defer wg.Done()
	// printf выполнится даже при Goexit потому что-то использует defer.
	defer fmt.Printf("Это сообщение напечатается даже при Goexit потому что-то использует defer.\n")
	runtime.Goexit()                                       // выходим из горутины
	fmt.Printf("Это сообщение никогда не напечатается.\n") // сообщение не напечатается, потому что горутина завершилась
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go exitByGoexit(&wg) // запускаем функцию в горутине
	wg.Wait()
}
