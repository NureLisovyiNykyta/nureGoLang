package main

import (
	"fmt"
	"sort"
	"strings"
)

type AEROFLOT struct {
	NAZN string
	NUMR int
	TIP  string
}

type ByNAZN []AEROFLOT

func (a ByNAZN) Len() int           { return len(a) }
func (a ByNAZN) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByNAZN) Less(i, j int) bool { return a[i].NAZN < a[j].NAZN }

func task6() {
	var AIRPORT [7]AEROFLOT

	for i := 0; i < 7; i++ {
		var a AEROFLOT
		fmt.Scan(&a.NAZN, &a.NUMR, &a.TIP)
		AIRPORT[i] = a
	}

	sort.Sort(ByNAZN(AIRPORT[:]))

	var searchTip string
	fmt.Scan(&searchTip)

	found := false
	for _, flight := range AIRPORT {
		if strings.EqualFold(flight.TIP, searchTip) {
			fmt.Printf("Flight number: %d, Aircraft type: %s\n", flight.NUMR, flight.TIP)
			found = true
		}
	}

	if !found {
		fmt.Println("No flights found for the specified aircraft type")
	}
}