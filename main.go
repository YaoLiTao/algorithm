package main

import (
	"algorithm/yami64"
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
	//fmt.Println("MergeSort time:", time.Now().Sub(begin))
	//fmt.Println("MergeSort: ", data)

	str := "我是皮卡丘！"
	fmt.Println("原始字符串:", str)
	encoded := yami64.Encode("我是皮卡丘！")
	fmt.Println("yami64编码:", encoded)
	raw := yami64.Decode(encoded)
	fmt.Println("yami64解码:", raw)
}

func getRandInt() []int {
	data := [100]int{0}
	for i := 0; i < len(data); i++ {
		data[i] = rand.Intn(1000)
	}
	//fmt.Println("randInt: ", data)
	return data[:]
}
