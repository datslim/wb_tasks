package main

import (
	"fmt"
	"time"
)

// функция для задержки выполнения на заданное время
func sleep(duration time.Duration) {
	done := make(chan struct{}) // канал для сигнала о завершении

	go func() {
		<-time.After(duration) // ожидаем заданное время (пока в канал не придет сигнал)
		close(done)            // отправляем сигнал о завершении
	}()

	<-done // ожидаем сигнал о завершении (приостанавливаем выполнение функции до получения сигнала)
}

func main() {
	var seconds int
	fmt.Println("Введите количество секунд для задержки:")
	fmt.Scan(&seconds)
	duration := time.Duration(seconds) * time.Second

	sleep(duration)
	fmt.Println("Время вышло.")
}
