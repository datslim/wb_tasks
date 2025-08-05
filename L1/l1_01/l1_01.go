package main

import "fmt"

type Human struct {
	Gender string
	Age    int
}

type Action struct {
	ActionType string
	Human      // встраиваем структуру Human в структуру Action (embedded struct)
}

// конструктор для Human
func newHuman(gender string, age int) *Human {
	return &Human{
		Gender: gender,
		Age:    age,
	}
}

// конструктор для Action
func newAction(action string, human *Human) *Action {
	return &Action{
		ActionType: action,
		Human:      *human,
	}
}

// метод для получения возраста из Human
func (h *Human) GetAge() int {
	return h.Age
}

// метод для получения пола из Human
func (h *Human) GetGender() string {
	return h.Gender
}

// метод для получения типа действия из Action
func (a *Action) GetActionType() string {
	return a.ActionType
}

func main() {
	Sasha := newHuman("Мужчина", 25)
	Dancing := newAction("Танцы", Sasha)
	fmt.Printf("Возраст: %d\n", Dancing.GetAge()) // напрямую вызываем метод GetAge для Action
	fmt.Printf("Пол: %s\n", Dancing.GetGender())  // напрямую вызываем метод GetGender для Action
	fmt.Printf("Тип действия: %s\n", Dancing.GetActionType())

}
