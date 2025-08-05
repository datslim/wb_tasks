package main

import (
	"fmt"
	"math"
)

// структура для хранения координат точки
type Point struct {
	x, y float64 // приватные поля
}

// конструктор для создания точки
// аргументы: x и y
// возвращаемое значение: указатель на точку
func NewPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

// поскольку поля нашей структуры должны быть приватными,
// то для работы с ними из другого пакета, нужно реализовать
// методы для получения приватных полей (геттеры)

// функция для получения приватного поля x
func (p *Point) GetX() float64 {
	return p.x
}

// функция для получения приватного поля y
func (p *Point) GetY() float64 {
	return p.y
}

// метод для вычисления расстояния между двумя точками
// аргументы: другая точка
// возвращаемое значение: расстояние типа float64
// формула нахождения расстояния между двумя точками на плоскости:
// √((x₂ - x₁)² + (y₂ - y₁)²)
func (p *Point) Distance(other Point) float64 {
	return math.Sqrt(math.Pow((other.GetX()-p.GetX()), 2) + math.Pow(other.GetY()-p.GetY(), 2))
}

func main() {
	// простой пример с двумя положительными числами, без плавающей точки
	a := NewPoint(2, 3) // a (2,3)
	b := NewPoint(5, 7) // b (5,7)
	fmt.Printf("Расстояние между А (%f, %f) и B (%f, %f): %f\n\n", a.GetX(), a.GetY(), b.GetX(), b.GetY(), a.Distance(*b))

	// пример с отрицательными координатами и плавающей точкой
	negativeFloatA := NewPoint(2.7, 3.7)   // a (2.7,3.7)
	negativeFloatB := NewPoint(-2.7, -3.7) // b (-2.7,-3.7)
	fmt.Printf("Расстояние между А (%f, %f) и B (%f, %f): %f\n\n", negativeFloatA.GetX(), negativeFloatA.GetY(), negativeFloatB.GetX(), negativeFloatB.GetY(), negativeFloatA.Distance(*negativeFloatB))

	// пример c одинаковыми точками (расстояние 0)
	samePointsA := NewPoint(2, 3) // a (0,0)
	samePointsB := NewPoint(2, 3) // b (0,0)
	fmt.Printf("Расстояние между А (%f, %f) и B (%f, %f): %f\n\n", samePointsA.GetX(), samePointsA.GetY(), samePointsB.GetX(), samePointsB.GetY(), samePointsA.Distance(*samePointsB))

	// пример с точками на одной горизонтальной оси (y одинаковый)
	horizontalA := NewPoint(0, 5) // a (0,5)
	horizontalB := NewPoint(8, 5) // b (8,5)
	fmt.Printf("Расстояние между А (%f, %f) и B (%f, %f): %f\n\n", horizontalA.GetX(), horizontalA.GetY(), horizontalB.GetX(), horizontalB.GetY(), horizontalA.Distance(*horizontalB))

	// пример с точками на одной вертикальной оси (x одинаковый)
	verticalA := NewPoint(3, 0)  // a (3,0)
	verticalB := NewPoint(3, 12) // b (3,12)
	fmt.Printf("Расстояние между А (%f, %f) и B (%f, %f): %f\n\n", verticalA.GetX(), verticalA.GetY(), verticalB.GetX(), verticalB.GetY(), verticalA.Distance(*verticalB))

	// пример с очень большими координатами
	bigCoordsA := NewPoint(1000000, 2000000) // a (1000000,2000000)
	bigCoordsB := NewPoint(3000000, 4000000) // b (3000000,4000000)
	fmt.Printf("Расстояние между А (%f, %f) и B (%f, %f): %f\n\n", bigCoordsA.GetX(), bigCoordsA.GetY(), bigCoordsB.GetX(), bigCoordsB.GetY(), bigCoordsA.Distance(*bigCoordsB))

	// пример с очень маленькими координатами (близкими к нулю)
	smallCoordsA := NewPoint(0.0001, 0.0002) // a (0.0001,0.0002)
	smallCoordsB := NewPoint(0.0003, 0.0004) // b (0.0003,0.0004)
	fmt.Printf("Расстояние между А (%f, %f) и B (%f, %f): %f\n\n", smallCoordsA.GetX(), smallCoordsA.GetY(), smallCoordsB.GetX(), smallCoordsB.GetY(), smallCoordsA.Distance(*smallCoordsB))
}
