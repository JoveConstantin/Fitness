package personaldata

import (
	"fmt"
)

// Ниже создайте структуру Personal
type Personal struct {
	Name   string
	Weight float64
	Height float64
}

// Ниже создайте метод Print()
func (p Personal) Print() {
	fmt.Println("Имя:", p.Name)
	fmt.Println("Вес:", p.Weight)
	fmt.Println("Рост:", p.Height)
}
