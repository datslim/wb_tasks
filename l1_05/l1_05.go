package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// функция для постоянной записи в канал (пока не закроется контекст)
// аргументы: контекст, группа ожидания, канал
func writer(ctx context.Context, wg *sync.WaitGroup, data chan int) {
	wg.Add(1) // добавляем 1 горутину в вейтгруппу

	go func() { // запускаем горутину
		defer wg.Done() // уменьшаем счетчик вейтгруппы
		i := 0          // создаем переменную, которую будем отправлять в канал
		for {
			select {
			case data <- i: // записываем в канал
				i++ // увеличиваем переменную
			case <-ctx.Done(): // если контекст завершен, закрываем канал и выходим из функции
				close(data)
				return
			}
		}
	}()
}

// функция для чтения из канала (пока не закроется канал)
// аргументы: группа ожидания, канал
func reader(wg *sync.WaitGroup, data chan int) {
	wg.Add(1) // добавляем 1 горутину в вейтгруппу

	go func() { // запускаем горутину
		defer wg.Done()       // уменьшаем счетчик вейтгруппы
		for i := range data { // читаем из канала
			fmt.Printf("Значение: %d\n", i)
		}
		fmt.Println("Канал закрыт.")
	}()
}

func main() {
	fmt.Printf("Введите количество секунд: ")
	var N int
	fmt.Scan(&N) // считываем количество секунд

	data := make(chan int) // создаем канал для передачи данных

	// создаем контекст с таймаутом в N секунд и отменяем его по завершении
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(N)*time.Second)
	defer cancel() // отменяем контекст по завершении

	var wg sync.WaitGroup  // создаем группу ожидания
	writer(ctx, &wg, data) // запускаем функцию для записи в канал
	reader(&wg, data)      // запускаем функцию для чтения из канала

	wg.Wait() // ожидаем завершения всех горутин

	fmt.Printf("Завершено через %d секунд(ы)\n", N)
}
