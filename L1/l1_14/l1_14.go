package main

import (
	"fmt"

	"github.com/datslim/wb-tech-misc/misc"
)

// функция для определения типа переменной в рантайме.
// any - это алиас для пустого интерфейса (любой тип)
func detectType(variable any) string {
	switch variable.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan int:
		return "chan int"
	case chan string:
		return "chan string"
	default:
		return "unknown type"
	}
}

// проверяем работу нашей функции
// используем nil как неизвестный тип
func main() {
	variablesOfDifferentTypes := []any{25, "Lorem Ipsum", true, make(chan int), make(chan string), nil}

	for _, currentVariable := range variablesOfDifferentTypes {
		fmt.Printf("%v is %v\n", currentVariable, detectType(currentVariable))
	}

	randomVariablesOfDifferentTypes := make([]any, 10)
	misc.FillSetWithRandomValues(randomVariablesOfDifferentTypes)

	// при попадании числа с плавающей точкой, мы получим unknown type поскольку в задании,
	// их не нужно определять
	for _, currentVariable := range randomVariablesOfDifferentTypes {
		fmt.Printf("%v is %v\n", currentVariable, detectType(currentVariable))
	}
}
