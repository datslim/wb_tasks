package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// пример остановки выполнения горутины используя контекст
func exitByContext(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	i := 0 // создаем переменную, которую будем инкреминтировать
	for {
		select {
		case <-ctx.Done(): // когда контекст завершается
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

	ctx, cancel := context.WithCancel(context.Background()) // создаем контекст с отменой

	wg.Add(1)

	go exitByContext(ctx, &wg) // запускаем функцию в горутине

	time.Sleep(5 * time.Second) // ждем 5 секунд
	cancel()                    // завершаем контекст

	fmt.Printf("Контекст завершился спустя 5 секунд\n")
	wg.Wait()
}
