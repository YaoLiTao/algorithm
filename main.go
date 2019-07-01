package main

import (
	"algorithm/sort"
	"fmt"
	"math/rand"
)

func main() {
	var data []int

	data = getRandInt()
	sort.QuickSort(data, 0, len(data)-1)
	fmt.Println(data)

	data = getRandInt()
	sort.BubbleSort(data[:])
	fmt.Println(data)

	data = getRandInt()
	sort.SelectionSort(data[:])
	fmt.Println(data)

	data = getRandInt()
	sort.StraightInsertionSort(data[:])
	fmt.Println(data)
}

func getRandInt() []int {
	data := [100]int{}
	for i := 0; i < len(data); i++ {
		data[i] = rand.Intn(1000)
	}
	fmt.Println("randInt: ", data)
	return data[:]
}
