package main

import "fmt"

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

	for _, currentType := range variablesOfDifferentTypes {
		fmt.Println(detectType(currentType))
	}
}
