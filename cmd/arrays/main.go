package main

import (
	"fmt"
	"math/bits"
)

type Point struct {
	X int
	Y int
}

func main() {
	Newmap := map[string]Point{}
	// Newmap is a mapped by strings array with points
	Newmap
	fmt.Println(Newmap)

	array := []int{1, 2, 3}
	var max int = (1 << bits.UintSize) / -2
	var min int = (1 << bits.UintSize) / 2 - 1
	for _, el := range array {
		if el > max {max = el}
		if el < min {max = el}
		}
	}
}

// Здесь потестируем всякие массивы
/*
slice := []string{"a", "b", "c"}
//перебрали массив и элементы
for index, element := range slice {
	fmt.Printf("element #%d: %s\n", index, element)
}
// перебрали элементы
for _, element := range slice {
	fmt.Printf("element without number: %s\n", element)
}
//перебрали индексы
for index := range slice {
	fmt.Printf("element without number: %s\n", index)
}*/
