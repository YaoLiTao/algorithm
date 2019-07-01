package main

import (
	"algorithm/sort"
	"fmt"
	"math/rand"
)

func main() {
	data := [100]int{}
	for i := 0; i < len(data); i++ {
		data[i] = rand.Intn(1000)
	}
	fmt.Println(data)
	//sort.QuickSort(data[:], 0, len(data)-1)
	//sort.BubbleSort(data[:])
	//sort.SelectionSort(data[:])
	sort.StraightInsertionSort(data[:])
	fmt.Println(data)
}
