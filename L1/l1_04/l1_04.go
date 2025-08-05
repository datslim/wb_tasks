package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

// функция для считывания задач из канала и вывода результата
// аргументы: ctx - контекст для завершения программы, workerID - идентификатор воркера,
// wg - для синхронизации горутин, jobs - канал для получения задач
func worker(ctx context.Context, workerID int, wg *sync.WaitGroup, jobs <-chan int) {
	defer wg.Done() // уменьшаем счетчик WaitGroup

	for {
		select {
		case result, ok := <-jobs: // получаем задачу из канала
			if !ok { // если не можем получить задачу из канала, то завершаем процесс
				return
			} else {
				fmt.Printf("Воркер %d содержит %d\n", workerID, result) // иначе выводим результат
			}
		case <-ctx.Done(): // когда получили команду на SIGINT, завершаем работу
			fmt.Printf("SIGINT. Воркер %d завершается.\n", workerID)
			return
		}
	}
}

// функция для бесконечного заполнения канала случайными числами
// аргументы: канал для отправки задач
func infiniteFillChan(ctx context.Context, jobs chan<- int) {
	for {
		select {
		case jobs <- rand.Intn(100): // отправляем случайное число в канал
			time.Sleep(time.Second * 1)
		case <-ctx.Done(): // когда получили команду на SIGINT, завершаем работу
			close(jobs) // закрываем канал, чтобы воркеры завершились
			return
		}
	}
}

func main() {
	// проверяем количество аргументов
	switch {
	case len(os.Args) < 2:
		fmt.Println("Аргументы не могут быть пустыми. Введите количество worker'ов")
		os.Exit(1) // завершаем программу с кодом 1
	case len(os.Args) > 2:
		fmt.Println("Не поддерживается больше 1 аргумента. Введите количество worker'ов")
		os.Exit(1)
	}

	// преобразуем аргумент в число
	workersNumber, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Аргумент должен быть числом. Введите количество worker'ов") // выводим ошибку, если аргумент не число
		os.Exit(1)
	}

	jobs := make(chan int) // создаем небуферизированный канал для отправки задач
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() { // запускаем горутину для обработки сигналов
		<-sigChan // ждем сигнал завершения
		fmt.Println("Получен сигнал завершения.")
		cancel() // отменяем контекст
	}()

	go infiniteFillChan(ctx, jobs) // запускаем горутину для бесконечного заполнения канала случайными числами
	var wg sync.WaitGroup

	for workerID := range workersNumber { //
		wg.Add(1)
		go worker(ctx, workerID, &wg, jobs) // запускаем горутину для выполнения задачи
	}

	wg.Wait() // ожидаем завершения всех горутин

}
