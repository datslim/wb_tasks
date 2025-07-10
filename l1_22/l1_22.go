package main

import (
	"fmt"
	"math/big"
	"math/rand"
)

type BigInt struct {
	value big.Int
}

func NewBigInt() *BigInt {
	return &BigInt{
		value: *big.NewInt(0),
	}
}

func NewBigIntFromInt(value int) *BigInt {
	return &BigInt{
		value: *big.NewInt(int64(value)),
	}
}

func (a *BigInt) Add(b *BigInt) {
	a.value.Add(&a.value, &b.value)
}

func (a *BigInt) Subtract(b *BigInt) {
	a.value.Sub(&a.value, &b.value)
}

func (a *BigInt) Divide(b *BigInt) {
	a.value.Div(&a.value, &b.value)
}

func (a *BigInt) Multiply(b *BigInt) {
	a.value.Mul(&a.value, &b.value)
}

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
	bigA.Add(bigB) // складываем
	fmt.Printf("%v\n\n", bigA.String())

	fmt.Printf("%v * %v = ", bigA.value.String(), bigB.value.String())
	bigA.Multiply(bigB) // умножаем
	fmt.Printf("%v\n\n", bigA.String())

	fmt.Printf("%v / %v = ", bigA.value.String(), bigB.value.String())
	bigA.Divide(bigB) // делим
	fmt.Printf("%v\n\n", bigA.String())

	fmt.Printf("%v - %v = ", bigA.value.String(), bigB.value.String())
	bigA.Subtract(bigB) // вычитаем
	fmt.Printf("%v\n\n", bigA.String())

	endA := bigA // сохраняем конечное значение переменной А, для сравнения с исходным значением, если они равны - значит вычисления были точными

	if startingA == endA {
		fmt.Printf("После всех операций и обратных операций, значение не изменилось: A = %v\n", startingA.value.String())
	} else {
		fmt.Printf("После всех операций и обратных операций, значение А изменилось, было: %v, стало: %v\n", startingA.value.String(), endA.value.String())
	}

}
