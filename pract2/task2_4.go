package main

import (
    "fmt"
    "math"
)

// Інтерфейс для функції
type Equation interface {
    Value(x float64) float64
    Derivative(x float64) float64
}

// Структура для рівняння 3x - 4lnx - 5 = 0
type MyEquation struct{}

func (MyEquation) Value(x float64) float64 {
    return 3*x - 4*math.Log(x) - 5
}

func (MyEquation) Derivative(x float64) float64 {
    return 3 - 4/x
}

// Функція методу Ньютона (приймає покажчик на інтерфейс)
func NewtonMethod(eq Equation, x0 float64, eps float64, maxIter int) float64 {
    x := x0
    for i := 0; i < maxIter; i++ {
        fx := eq.Value(x)
        dfx := eq.Derivative(x)
        if math.Abs(dfx) < 1e-10 {
            fmt.Println("Похідна занадто мала, метод не сходиться.")
            break
        }
        xNext := x - fx/dfx
        if math.Abs(xNext-x) < eps {
            return xNext
        }
        x = xNext
    }
    return x
}

func task2_4() {
    eq := &MyEquation{}
    start := 2.0
    eps := 1e-6
    maxIter := 100

    root := NewtonMethod(eq, start, eps, maxIter)

    fmt.Printf("Метод Ньютона:\n")
    fmt.Printf("Знайдений корінь: %.4f\n", root)
    fmt.Printf("Перевірка: f(%.4f) = %.6f\n", root, eq.Value(root))
    fmt.Println("Точне значення за умовою: 3.2300")
}
