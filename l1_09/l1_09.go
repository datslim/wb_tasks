package main

import "fmt"

// функция для записи значений массива в первый канал
func generateValues(out chan<- int, values [5]int) {
	for _, value := range values {
		out <- value
	}
	close(out)
}

// функция для обработки (возведения в квадрат) значений из первого канала и записи в второй
func processValues(in <-chan int, out chan<- int) {
	for value := range in {
		out <- value * 2
	}
	close(out)
}

func main() {
	array := [5]int{1, 2, 3, 4, 5} // массив значений

	ch1 := make(chan int) // создаем первый канал чисел
	ch2 := make(chan int) // создаем второй канал чисел

	go generateValues(ch1, array) // запускаем функцию для записи значений в первый канал
	go processValues(ch1, ch2)    // запускаем функцию для обработки значений из первого канала и записи в второй

	for value := range ch2 { // читаем из второго канала
		fmt.Println(value) // печатаем значение
	}
}
