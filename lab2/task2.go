package main

import "fmt"

type Fuel struct {
	Name   string
	Volume float64
}

type DayStats struct {
	DayName    string
	Fuels      []Fuel
	CarsServed int
}

type FuelMachine struct {
	ID   int
	Week []DayStats
}

// Метод: обчислення загальної кількості проданого палива
func (fm *FuelMachine) TotalFuel() float64 {
	total := 0.0
	for _, day := range fm.Week {
		for _, f := range day.Fuels {
			total += f.Volume
		}
	}
	return total
}

// Метод: день, коли обслужено найбільше машин
func (fm *FuelMachine) DayWithMaxCars() string {
	maxCars := 0
	bestDay := ""
	for _, day := range fm.Week {
		if day.CarsServed > maxCars {
			maxCars = day.CarsServed
			bestDay = day.DayName
		}
	}
	return bestDay
}

// Функція: автомат з найбільшим обсягом продажів
func MachineWithMaxFuel(machines []*FuelMachine) *FuelMachine {
	var best *FuelMachine
	maxFuel := 0.0

	for _, m := range machines {
		total := m.TotalFuel()
		if total > maxFuel {
			maxFuel = total
			best = m
		}
	}
	return best
}

// Функція: найпопулярніше паливо серед усіх автоматів
func MostPopularFuel(machines []*FuelMachine) string {
	fuelTotals := map[string]float64{}

	for _, m := range machines {
		for _, day := range m.Week {
			for _, f := range day.Fuels {
				fuelTotals[f.Name] += f.Volume
			}
		}
	}

	maxFuel := ""
	maxVal := 0.0
	for name, volume := range fuelTotals {
		if volume > maxVal {
			maxVal = volume
			maxFuel = name
		}
	}

	return maxFuel
}

func task2() {
	machine1 := &FuelMachine{
		ID: 1,
		Week: []DayStats{
			{"Понеділок", []Fuel{{"A92", 50}, {"A95", 40}, {"Euro-5", 20}, {"DIZ", 30}}, 60},
			{"Вівторок", []Fuel{{"A92", 60}, {"A95", 30}, {"Euro-5", 40}, {"DIZ", 50}}, 70},
			{"Середа", []Fuel{{"A92", 30}, {"A95", 50}, {"Euro-5", 20}, {"DIZ", 40}}, 80},
			{"Четвер", []Fuel{{"A92", 80}, {"A95", 60}, {"Euro-5", 30}, {"DIZ", 60}}, 75},
			{"П’ятниця", []Fuel{{"A92", 100}, {"A95", 50}, {"Euro-5", 40}, {"DIZ", 70}}, 90},
			{"Субота", []Fuel{{"A92", 90}, {"A95", 60}, {"Euro-5", 30}, {"DIZ", 80}}, 95},
			{"Неділя", []Fuel{{"A92", 70}, {"A95", 50}, {"Euro-5", 20}, {"DIZ", 60}}, 85},
		},
	}

	machine2 := &FuelMachine{
		ID: 2,
		Week: []DayStats{
			{"Понеділок", []Fuel{{"A92", 40}, {"A95", 60}, {"Euro-5", 30}, {"DIZ", 20}}, 65},
			{"Вівторок", []Fuel{{"A92", 50}, {"A95", 50}, {"Euro-5", 20}, {"DIZ", 40}}, 68},
			{"Середа", []Fuel{{"A92", 60}, {"A95", 40}, {"Euro-5", 20}, {"DIZ", 50}}, 72},
			{"Четвер", []Fuel{{"A92", 70}, {"A95", 60}, {"Euro-5", 30}, {"DIZ", 60}}, 85},
			{"П’ятниця", []Fuel{{"A92", 90}, {"A95", 70}, {"Euro-5", 40}, {"DIZ", 80}}, 100},
			{"Субота", []Fuel{{"A92", 80}, {"A95", 50}, {"Euro-5", 30}, {"DIZ", 70}}, 78},
			{"Неділя", []Fuel{{"A92", 60}, {"A95", 40}, {"Euro-5", 20}, {"DIZ", 50}}, 74},
		},
	}

	machines := []*FuelMachine{machine1, machine2}

	fmt.Println("=== Інформація про АЗС ===")
	fmt.Printf("Загальний обсяг палива автомата №1: %.2f л\n", machine1.TotalFuel())
	fmt .Printf("Загальний обсяг палива автомата №2: %.2f л\n", machine2.TotalFuel())
	fmt.Printf("День з найбільшою кількістю машин (автомат №1): %s\n", machine1.DayWithMaxCars())
	fmt.Printf("День з найбільшою кількістю машин (автомат №2): %s\n", machine2.DayWithMaxCars())

	bestMachine := MachineWithMaxFuel(machines)
	fmt.Printf("Найбільше палива продав автомат №%d (%.2f л)\n", bestMachine.ID, bestMachine.TotalFuel())

	popularFuel := MostPopularFuel(machines)
	fmt.Printf("Найпопулярніше паливо: %s\n", popularFuel)
}