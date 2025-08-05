package main

import (
	"fmt"
	"math/big"
	"math/rand"
)

type BigInt struct {
	value big.Int
}

// конструктор для создания нового BigInt
func NewBigInt() *BigInt {
	return &BigInt{
		value: *big.NewInt(0),
	}
}

// конструктор для создания нового BigInt из int
func NewBigIntFromInt(value int) *BigInt {
	return &BigInt{
		value: *big.NewInt(int64(value)),
	}
}

// нижеперечисленные методы меняют значение по указателю, но тем не менее,
// я решил возвращать результат вычисления для удобства

// метод для сложения значений двух структур типа BigInt
// ресивер - указатель на структуру А типа BigInt, аргумент - указатель на структуру B типа BigInt
// возвращаемое значение - результат сложения
func (a *BigInt) Add(b *BigInt) *big.Int {
	return a.value.Add(&a.value, &b.value)
}

// метод для вычитания значений двух структур типа BigInt
// ресивер - указатель на структуру А типа BigInt, аргумент - указатель на структуру B типа BigInt
// возвращаемое значение - результат вычитания
func (a *BigInt) Subtract(b *BigInt) *big.Int {
	return a.value.Sub(&a.value, &b.value)
}

// метод для деления значений двух структур типа BigInt
// ресивер - указатель на структуру А типа BigInt, аргумент - указатель на структуру B типа BigInt
// возвращаемое значение - результат деления
func (a *BigInt) Divide(b *BigInt) *big.Int {
	return a.value.Div(&a.value, &b.value)
}

// метод для умножения значений двух структур типа BigInt
// ресивер - указатель на структуру А типа BigInt, аргумент - указатель на структуру B типа BigInt
// возвращаемое значение - результат умножения
func (a *BigInt) Multiply(b *BigInt) *big.Int {
	return a.value.Mul(&a.value, &b.value)
}

// метод для удобного представления значения в виде числа
func (a *BigInt) String() string {
	return a.value.String()
}

func main() {
	powerOfA := rand.Intn(20) + 20                                                     // генерируем значение для возведения А в случайную степень (от 20 до 40) - минимум 1 048 576, максимум - 1 099 511 627 776
	powerOfB := rand.Intn(20) + 20                                                     // генерируем значение для возведения B в случайную степень (от 20 до 40)
	bigA, bigB := NewBigIntFromInt((1 << powerOfA)), NewBigIntFromInt((1 << powerOfB)) // создаем экземпляры BigInt, являющиеся 2 возведенной в степень powerOf(x)

	startingA := bigA // записываем исходное значение переменной А, чтобы сравнить ее с тем, что получится после 4 операций

	fmt.Printf("A = %v\nB = %v\n\n", bigA.value.String(), bigB.value.String())

	fmt.Printf("%v + %v = ", bigA.value.String(), bigB.value.String())
	add := bigA.Add(bigB) // складываем и сохраняем копию результата в переменную add
	fmt.Printf("%v\n\n", add.String())

	fmt.Printf("%v * %v = ", bigA.value.String(), bigB.value.String())
	multiply := bigA.Multiply(bigB) // умножаем и сохраняем копию результата в переменную multiply
	fmt.Printf("%v\n\n", multiply.String())

	fmt.Printf("%v / %v = ", bigA.value.String(), bigB.value.String())
	divide := bigA.Divide(bigB) // делим и сохраняем копию результата в переменную divide
	fmt.Printf("%v\n\n", divide.String())

	fmt.Printf("%v - %v = ", bigA.value.String(), bigB.value.String())
	subtract := bigA.Subtract(bigB) // вычитаем и сохраняем копию результата в переменную subtract
	fmt.Printf("%v\n\n", subtract.String())

	endA := bigA // сохраняем конечное значение переменной А, для сравнения с исходным значением, если они равны - значит вычисления были точными

	if startingA == endA {
		fmt.Printf("После всех операций и обратных операций, значение не изменилось: A = %v\n", startingA.value.String())
	} else {
		fmt.Printf("После всех операций и обратных операций, значение А изменилось, было: %v, стало: %v\n", startingA.value.String(), endA.value.String())
	}

}
