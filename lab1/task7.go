package main

import (
	"fmt"
)

type student struct {
	name  *string
	kurs  int
	rating float32
}

func task7() {
	var n int
	fmt.Print("Введіть кількість студентів: ")
	fmt.Scan(&n)

	students := make([]student, n)

	for i := 0; i < n; i++ {
		var name string
		var kurs int
		var rating float32
		fmt.Print("Введіть ім'я, курс та рейтинг студента: ")
		fmt.Scan(&name, &kurs, &rating)
		students[i].name = &name
		students[i].kurs = kurs
		students[i].rating = rating
	}

	for _, s := range students {
		if s.kurs == 3 {
			fmt.Printf("Ім'я: %s, Рейтинг: %.2f\n", *s.name, s.rating)
		}
	}
}