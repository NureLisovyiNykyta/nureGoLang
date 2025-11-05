package main

import (
	"fmt"
	"nureGoLang/lab3/fileutils"
)

// Демонстрація роботи
func task2() {
	filename := "stadiums.json"

	stadiums := []fileutils.Stadium{
		{"Олімпійський", "Київ, вул. Велика Васильківська, 55", 70050, []string{"Футбол", "Легка атлетика"}},
		{"Металіст", "Харків, Плеханівська 65", 40000, []string{"Футбол"}},
		{"Арена Львів", "Львів, Стрийська 199", 35000, []string{"Футбол", "Концерти"}},
	}

	fileutils.CreateFile(filename, stadiums)
	fmt.Println("\n=== Вміст створеного файлу ===")
	fmt.Println(fileutils.ReadFile(filename))

	fmt.Println("\n=== Видалення ===")
	fileutils.DeleteByName(filename, "Металіст")
	fmt.Println("=== Після видалення ===")
	fmt.Println(fileutils.ReadFile(filename))

	
	fmt.Println("\n=== Додавання нових елементів ===")
	newStadiums := []fileutils.Stadium{
		{"Дніпро Арена", "Дніпро, Набережна Перемоги", 31000, []string{"Футбол"}},
		{"Чорноморець", "Одеса, вул. Льва Толстого, 3", 34000, []string{"Футбол"}},
	}
	fileutils.AddStadiums(filename, newStadiums)
	fmt.Println("=== Після додавання ===")
	fmt.Println(fileutils.ReadFile(filename))
}