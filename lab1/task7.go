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
		fmt.Print("Введіть ім'я, курс та рейтинг (1-5) студента: ")
		fmt.Scan(&name, &kurs, &rating)
		students[i].name = &name
		students[i].kurs = kurs
		students[i].rating = rating
	}

	for _, s := range students {
		if s.rating < 3 {
			fmt.Printf("Ім'я: %s, Рейтинг: %.2f\n", *s.name, s.rating)
		}
	}

	fmt.Println("Введіть кількість студентів для видалення з кінця списку:")
	var count int
	fmt.Scan(&count)

	if count > n {
		count = n
	}

	new_students := make([]student, n-count)

	fmt.Println("Оновлений список студентів:")
	for i := 0; i < n-count; i++ {
		new_students[i] = students[i]
		fmt.Printf("Ім'я: %s, Курс: %d, Рейтинг: %.2f\n", *new_students[i].name, new_students[i].kurs, new_students[i].rating)
	}
	students = new_students
}