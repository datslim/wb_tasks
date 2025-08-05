package main

import (
	"fmt"
	"sync"
)

// пример остановки выполнения горутины по условию
// умножаем i (2) на 2 до тех пор пока не достигнем значения больше 1000
func exitByCondition(wg *sync.WaitGroup) {
	i := 2
	defer wg.Done()
	for {
		fmt.Println(i)
		i *= 2
		if i > 1000 { // при i больше 1000, выходим из цикла -> вызываем wg.Done (завершаем горутину)
			fmt.Printf("i (%d) больше 1000. выходим из горутины\n", i)
			return
		}
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go exitByCondition(&wg) // запускаем функцию в горутине
	wg.Wait()
}
