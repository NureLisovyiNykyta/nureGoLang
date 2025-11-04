package main

import "fmt"

type Person struct {
    Name     string
    Address  string
    Birthday string
}

type Student struct {
    Person
    Course int
    Group  string
    Scores []float64
}

// Явно задані параметри
func PrintStudentExplicit(name string, course int, group string) {
    fmt.Printf("Студент: %s, Курс: %d, Група: %s\n", name, course, group)
}

// Неявно задані параметри (через структуру)
func PrintStudentImplicit(s Student) {
    fmt.Printf("Студент: %s, Курс: %d, Група: %s\n", s.Name, s.Course, s.Group)
}

// Частково явно, частково неявно
func PrintStudentMixed(name string, s Student) {
    fmt.Printf("Студент: %s, Курс: %d, Група: %s\n", name, s.Course, s.Group)
}

// Явно передані параметри
func AverageExplicit(numbers ...float64) float64 {
    sum := 0.0
    for _, n := range numbers {
        sum += n
    }
    return sum / float64(len(numbers))
}

// Неявно — беремо з поля структури
func AverageImplicit(s Student) float64 {
    sum := 0.0
    for _, n := range s.Scores {
        sum += n
    }
    return sum / float64(len(s.Scores))
}

// Частково явно — додаємо ще одне число вручну
func AverageMixed(s Student, extra float64) float64 {
    sum := extra
    for _, n := range s.Scores {
        sum += n
    }
    return sum / float64(len(s.Scores)+1)
}

type Data struct {
    Person
    Numbers []int
}

// Явно
func MaxIndexExplicit(numbers []int) int {
    maxIdx := 0
    for i, n := range numbers {
        if n > numbers[maxIdx] {
            maxIdx = i
        }
    }
    return maxIdx
}

// Неявно
func MaxIndexImplicit(d Data) int {
    maxIdx := 0
    for i, n := range d.Numbers {
        if n > d.Numbers[maxIdx] {
            maxIdx = i
        }
    }
    return maxIdx
}

// Частково явно
func MaxIndexMixed(d Data, extra int) int {
    numbers := append(d.Numbers, extra)
    maxIdx := 0
    for i, n := range numbers {
        if n > numbers[maxIdx] {
            maxIdx = i
        }
    }
    return maxIdx
}


func task2() {
    student := Student{
        Person: Person{Name: "Лісовий Нікіта", Address: "Київ", Birthday: "2004-02-14"},
        Course: 3,
        Group:  "ПІ-31",
        
        Scores: []float64{90, 80, 100},
    }

    fmt.Println("=== Явно передані параметри ===")
    PrintStudentExplicit("Лісовий Нікіта", 3, "ПІ-31")

    fmt.Println("=== Неявно (через структуру) ===")
    PrintStudentImplicit(student)

    fmt.Println("=== Частково явно, частково неявно ===")
    PrintStudentMixed("Лісовий Нікіта", student)

    fmt.Println("=== Обчислення середнього балу ===")
    fmt.Println("Явно:", AverageExplicit(90, 80, 100))
    fmt.Println("Неявно:", AverageImplicit(student))
    fmt.Println("Частково явно:", AverageMixed(student, 90))

    data := Data{
        Person:  Person{Name: "Лісовий Нікіта", Address: "Київ", Birthday: "2004-02-14"},
        Numbers: []int{5, 8, 12, 3, 10},
    }

    fmt.Println("=== Пошук індексу максимального елемента ===")
    fmt.Println("Явно:", MaxIndexExplicit([]int{5, 8, 12, 3, 10}))
    fmt.Println("Неявно:", MaxIndexImplicit(data))
    fmt.Println("Частково явно:", MaxIndexMixed(data, 0))
}