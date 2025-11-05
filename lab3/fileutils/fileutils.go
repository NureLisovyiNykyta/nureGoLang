package fileutils

import (
	"encoding/json"
	"fmt"
	"os"
)

type Stadium struct {
	Name     string
	Address  string
	Capacity int
	Sports   []string
}

// Створення файлу та запис структурованих даних
func CreateFile(filename string, data []Stadium) {
	file, err := os.Create(filename)
	if err != nil {
		panic(fmt.Sprintf("Помилка створення файлу: %v", err))
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(data); err != nil {
		panic(fmt.Sprintf("Помилка запису даних: %v", err))
	}
	fmt.Println("Файл успішно створено.")
}

// Зчитування та виведення файлу
func ReadFile(filename string) []Stadium {
	file, err := os.Open(filename)
	if err != nil {
		panic(fmt.Sprintf("Помилка відкриття файлу: %v", err))
	}
	defer file.Close()

	var data []Stadium
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		panic(fmt.Sprintf("Помилка зчитування файлу: %v", err))
	}
	return data
}

// Видалення елемента за назвою
func DeleteByName(filename, name string) {
	data := ReadFile(filename)
	var newData []Stadium
	for _, s := range data {
		if s.Name != name {
			newData = append(newData, s)
		}
	}
	WriteFile(filename, newData)
	fmt.Printf("Елемент '%s' видалено.\n", name)
}

// Додавання K нових елементів
func AddStadiums(filename string, newItems []Stadium) {
	data := ReadFile(filename)
	data = append(data, newItems...)
	WriteFile(filename, data)
	fmt.Printf("%d елементів додано.\n", len(newItems))
}

// Запис у файл (оновлення)
func WriteFile(filename string, data []Stadium) {
	file, err := os.Create(filename)
	if err != nil {
		panic(fmt.Sprintf("Помилка запису файлу: %v", err))
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(data); err != nil {
		panic(fmt.Sprintf("Помилка кодування JSON: %v", err))
	}
}