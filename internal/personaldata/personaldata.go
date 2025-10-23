package personaldata

import "fmt"

type Personal struct {
	Name   string  // имя пользователя;
	Weight float64 // вес пользователя;
	Height float64 // рост пользователя.
}

func (p Personal) Print() {
	fmt.Printf("Имя: %s\n", p.Name)
	fmt.Printf("Вес: %.2f кг.\n", p.Weight)
	fmt.Printf("Рост: %.2f м.\n", p.Height)
}
