package main

import (
	"fmt"
)

func main() {
	// int array
	var x [5]int
	x[4] = 100
	fmt.Println(x) //auto fill 0 with non-assignment

	// go through array
	var y [5]float64
	y[0] = 98
	y[1] = 93
	y[2] = 77
	y[3] = 82
	y[4] = 83
	var total float64 = 0
	for i := 0; i < len(y); i++ {
		total += y[i]
	}
	fmt.Println("Total =", total)
	fmt.Println("Average =", total/float64(len(y))) //Mind len(y) is int

	// slice不固定大小的陣列( have append copy function....etc)
	z := make([]float64, 4, 10)
	fmt.Println(z)

	fmt.Println(y[1:4]) // same as python

	//append
	fmt.Println(append(z, 100, 200, 300, 400, 555, 666, 777, 888))

	// Map, need memory malloc by using make
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"

	fmt.Println(elements["Li"])
	fmt.Println(elements)

	// add if
	name, ok := elements["Li"]
	fmt.Println(name, ok)

	// Map in Map
	people := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
		"Be": map[string]string{
			"name":  "Beryllium",
			"state": "solid",
		},
		"B": map[string]string{
			"name":  "Boron",
			"state": "solid",
		},
		"C": map[string]string{
			"name":  "Carbon",
			"state": "solid",
		},
		"N": map[string]string{
			"name":  "Nitrogen",
			"state": "gas",
		},
		"O": map[string]string{
			"name":  "Oxygen",
			"state": "gas",
		},
		"F": map[string]string{
			"name":  "Fluorine",
			"state": "gas",
		},
		"Ne": map[string]string{
			"name":  "Neon",
			"state": "gas",
		},
	}

	if el, ok := people["Li"]; ok {
		fmt.Println(el["name"], el["state"])
		fmt.Println(el)
	}
}
