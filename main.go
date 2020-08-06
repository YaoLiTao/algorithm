package main

import (
	"algorithm/search"
	"fmt"
	"math/rand"
)

func main() {
	//var data []int
	//var begin time.Time

	//data = getRandInt()
	//begin = time.Now()
	//sort.QuickSort(data, 0, len(data)-1)
	//fmt.Println("QuickSort time:", time.Now().Sub(begin))
	//fmt.Println("QuickSort", data[:50])
	//
	//data = getRandInt()
	//begin = time.Now()
	//sort.BubbleSort(data[:])
	//fmt.Println("BubbleSort time:", time.Now().Sub(begin))
	//fmt.Println("BubbleSort", data[:50])
	//
	//data = getRandInt()
	//begin = time.Now()
	//sort.SelectionSort(data[:])
	//fmt.Println("SelectionSort time:", time.Now().Sub(begin))
	//fmt.Println("SelectionSort", data[:50])
	//
	//data = getRandInt()
	//begin = time.Now()
	//sort.StraightInsertionSort(data[:])
	//fmt.Println("StraightInsertionSort time:", time.Now().Sub(begin))
	//fmt.Println("StraightInsertionSort", data[:50])
	//
	//data = getRandInt()
	//temp := make([]int, len(data))
	//begin = time.Now()
	//sort.MergeSort(data, temp, 0, len(data)-1)
	//fmt.Println("MergeSort time:", time.Now().Sub(begin))
	//fmt.Println("MergeSort: ", data[:50])

	//data = []int{3, 23, 5, 9, 55, 1, 43, 27}
	//fmt.Println(data)
	//begin = time.Now()
	//data = sort.HeapSort(data)
	//fmt.Println("HeapSort time:", time.Now().Sub(begin))
	//fmt.Println("HeapSort: ", data)

	//data = []int{3, 23, 5, 9, 55, 1, 43, 27, 77, 32, 65, 1000}
	//fmt.Println(data)
	//begin = time.Now()
	//data = sort.BitmapSort(data)
	//fmt.Println("BitmapSort time:", time.Now().Sub(begin))
	//fmt.Println("BitmapSort: ", data)
	//
	//str := "abc"
	//fmt.Println(str)
	//encoded := base64.Encode(str)
	//fmt.Println("yami64编码:", encoded)
	//raw := base64.Decode(encoded)
	//fmt.Println(raw)

	rbTree := search.NewRBTree()
	rbTree.Insert(3, 3)
	rbTree.Insert(2, 2)
	rbTree.Insert(9, 9)
	rbTree.Insert(0, 0)
	rbTree.Insert(1, 1)
	rbTree.Insert(4, 4)
	rbTree.Insert(7, 7)
	rbTree.Insert(5, 5)
	fmt.Println(rbTree.Get(0))
	fmt.Println(rbTree.Get(1))
	fmt.Println(rbTree.Get(2))
	fmt.Println(rbTree.Get(3))
	fmt.Println(rbTree.Get(4))
	fmt.Println(rbTree.Get(5))
	fmt.Println(rbTree.Get(6))
	fmt.Println(rbTree.Get(7))
	fmt.Println(rbTree.Get(8))
	fmt.Println(rbTree.Get(9))
}

func getRandInt() []int {
	data := [100000]int{0}
	for i := 0; i < len(data); i++ {
		data[i] = rand.Intn(1000000)
	}
	//fmt.Println("randInt: ", data)
	return data[:]
}
