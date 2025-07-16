package main

import "fmt"

// интерфейс для всех фигур
type Shape interface {
	Area() float64
}

// структура квадрата
type Square struct {
	sideLength float64
}

// метод для вычисления площади квадрата (реализует интерфейс Shape)
func (s *Square) Area() float64 {
	return s.sideLength * s.sideLength
}

// структура прямоугольника
type Rectangle struct {
	width  float64
	height float64
}

// у прямоугольника другой метод для вычисления площади - SquareArea (не реализует интерфейс Shape)
func (r *Rectangle) SquareArea() float64 {
	return r.width * r.height
}

// адаптер для прямоугольника для совместимости с интерфейсом Shape
type RectangleAdapter struct {
	rectangle *Rectangle // встраиваем прямоугольник в адаптер
}

// метод для вычисления площади прямоугольника (реализует интерфейс Shape)
func (r *RectangleAdapter) Area() float64 {
	return r.rectangle.SquareArea() // вычисляем и возвращаем площадь прямоугольника через его метод SquareArea
}

func main() {
	shapes := []Shape{ // создаем слайс фигур
		&Square{sideLength: 10.9},                                     // создаем квадрат
		&RectangleAdapter{rectangle: &Rectangle{width: 5, height: 3}}, // создаем прямоугольник с адаптером
	}

	for _, shape := range shapes { // проходим по всем фигурам и вычисляем их площадь
		fmt.Println(shape.Area()) // используем один и тот же метод Area для всех фигур
	}
}
