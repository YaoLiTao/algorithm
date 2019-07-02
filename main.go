package main

import (
	"algorithm/sort"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var data []int

	data = getRandInt()
	begin := time.Now()
	sort.QuickSort(data, 0, len(data)-1)
	fmt.Println("QuickSort time:", time.Now().Sub(begin))
	//fmt.Println("QuickSort", data)

	data = getRandInt()
	begin = time.Now()
	sort.BubbleSort(data[:])
	fmt.Println("BubbleSort time:", time.Now().Sub(begin))
	//fmt.Println("BubbleSort", data)

	data = getRandInt()
	begin = time.Now()
	sort.SelectionSort(data[:])
	fmt.Println("SelectionSort time:", time.Now().Sub(begin))
	//fmt.Println("SelectionSort", data)

	data = getRandInt()
	begin = time.Now()
	sort.StraightInsertionSort(data[:])
	fmt.Println("StraightInsertionSort time:", time.Now().Sub(begin))
	//fmt.Println("StraightInsertionSort", data)

	data = getRandInt()
	temp := [80000]int{}
	copy(temp[:], data)
	begin = time.Now()
	sort.MergeSort(data, temp[:], 0, len(data)-1)
	fmt.Println("MergeSort time:", time.Now().Sub(begin))
	//fmt.Println("MergeSort: ", data)
}

func getRandInt() []int {
	data := [80000]int{0}
	for i := 0; i < len(data); i++ {
		data[i] = rand.Intn(1000000)
	}
	//fmt.Println("randInt: ", data)
	return data[:]
}
