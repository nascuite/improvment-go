package main

import "fmt"

type Person struct {
	Name string
}

func changeName(person *Person) {
	//обновить имя
	//person.Name = "Alice"

	// обновить целиком структуру
	*person = Person{
		Name: "Alice",
	}
	// person начинает хранить новый адрес, уже на "Alice", но только в пределах changeName
	// по старому адресу person изменений нет

	//person = &Person{
	//	Name: "Alice",
	//}
}

func main() {
	person := &Person{
		Name: "Bob",
	}
	fmt.Println(person.Name)
	changeName(person)
	fmt.Println(person.Name)
}
