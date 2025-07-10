package main

import (
	"fmt"
	"sync"
)

type OrdersMap struct {
	mu     sync.RWMutex
	orders map[int]string // ключ - айди заказа, значение - заказанный товар
}

// функция для безопасного получения значения из словаря
func (m *OrdersMap) Get(key int) (string, bool) {
	m.mu.RLock() // блокируем на запись
	defer m.mu.RUnlock()
	value, ok := m.orders[key]
	return value, ok // возвращаем значение и true\false в зависимости от результата
}

// функция для безопасной записи значения в словарь
func (m *OrdersMap) Set(key int, value string) {
	m.mu.Lock() // полностью блокируем
	defer m.mu.Unlock()
	m.orders[key] = value
}

// конструктор для создания нового безопасного словаря
func NewOrdersMap(orders map[int]string) *OrdersMap {
	return &OrdersMap{
		orders: orders,
	}
}

// для тестирования нашего словаря, будем использовать три горутины
// конеурентно записывающих значения в словарь
func main() {
	var wg sync.WaitGroup // используем вейт-группу для синхронизации
	orders := NewOrdersMap(make(map[int]string))

	// конкурентно записываем значения
	for i := range 100 {
		wg.Add(3)

		for range 3 { // используем три горутины для записи
			go func(key int) {
				defer wg.Done()
				orders.Set(key, fmt.Sprintf("value is key = %d", key))
			}(i)
		}
	}

	wg.Wait()

	// конкурентно читаем значения
	for i := range 100 {
		wg.Add(1)
		go func(key int) {
			defer wg.Done()

			value, ok := orders.Get(key)
			if !ok {
				fmt.Printf("Ключ %d не был найден во время конкурентного чтения\n", key)
			}

			if value != fmt.Sprintf("value is key = %d", key) {
				fmt.Printf("Во время конкурентного чтения для ключа %d обнаружена ошибка: %s\n", key, value)
			}

		}(i)
	}

	wg.Wait()

	// проверяем, сохранились ли все ключи
	for key := range 100 {
		value, ok := orders.Get(key)
		if !ok {
			fmt.Printf("Ключ %d не был найден\n", key)
		}

		fmt.Printf("Ключ %d был найден со значением: %s\n", key, value)
	}

}
