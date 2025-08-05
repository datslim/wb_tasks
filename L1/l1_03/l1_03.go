package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

// функция для считывания задач из канала и вывода результата
// аргументы: workerID - идентификатор воркера, wg - для синхронизации горутин, jobs - канал для получения задач
func worker(workerID int, wg *sync.WaitGroup, jobs <-chan int) {
	defer wg.Done() // уменьшаем счетчик WaitGroup

	for result := range jobs { // получаем задачу из канала
		fmt.Printf("Воркер %d содержит %d\n", workerID, result) // выводим результат
	}
}

// функция для бесконечного заполнения канала случайными числами
// аргументы: канал для отправки задач
func infiniteFillChan(jobs chan<- int) {
	for {
		jobs <- rand.Intn(100)      // отправляем случайное число в канал
		time.Sleep(time.Second * 1) // спим одну секунду
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

	jobs := make(chan int)    // создаем небуферизированный канал для отправки задач
	go infiniteFillChan(jobs) // запускаем горутину для бесконечного заполнения канала случайными числами
	var wg sync.WaitGroup
	for workerID := range workersNumber { //
		wg.Add(1)
		go worker(workerID, &wg, jobs) // запускаем горутину для выполнения задачи
	}

	wg.Wait() // ожидаем завершения всех горутин

}
