package sort

import (
	"algorithm/tree"
	"math"
)

func swap(a *int, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

/**
冒泡排序
*/
func BubbleSort(data []int) {
	for i := 0; i < len(data); i++ {
		for j := len(data) - 1; j > i; j-- {
			if data[j-1] > data[j] {
				swap(&data[j-1], &data[j])
			}
		}
	}
}

/**
选择排序
*/
func SelectionSort(data []int) {
	for i := 0; i < len(data); i++ {
		for j := i; j < len(data); j++ {
			if data[i] > data[j] {
				swap(&data[i], &data[j])
			}
		}
	}
}

/**
直接插入排序
*/
func StraightInsertionSort(data []int) {
	j := 0
	k := 0
	for i := 1; i < len(data); i++ {
		for j = i - 1; j >= 0; j-- {
			if data[j] < data[i] {
				break
			}
		}

		if j != i-1 {
			temp := data[i]
			for k = i - 1; k > j; k-- {
				data[k+1] = data[k]
			}
			data[k+1] = temp
		}
	}
}

/**
快速排序
*/
func QuickSort(data []int, l int, r int) {
	if l < r {
		i := l
		j := r
		x := data[i]
		for i < j {
			for ; i < j && data[j] > x; j-- {

			}
			if i < j {
				data[i] = data[j]
				i++
			}
			for ; i < j && data[i] < x; i++ {

			}
			if i < j {
				data[j] = data[i]
				j--
			}
		}
		data[i] = x
		QuickSort(data, l, i-1)
		QuickSort(data, i+1, r)
	}
}

/**
归并排序
*/
func MergeSort(data []int, temp []int, l int, r int) {
	if l < r {
		center := (l + r) / 2
		MergeSort(data, temp, l, center)
		MergeSort(data, temp, center+1, r)

		i := l
		j := center + 1
		for k := l; k < r; k++ {
			if i > center {
				temp[k] = data[j]
				j++
			} else if j > r {
				temp[k] = data[i]
				i++
			} else if data[i] < data[j] {
				temp[k] = data[i]
				i++
			} else {
				temp[k] = data[j]
				j++
			}
		}

		for k := l; k < r; k++ {
			data[k] = temp[k]
		}
	}
}

/**
堆排序
*/
func HeapSort(data []int) []int {
	heap := tree.NewBinaryHeap(len(data))
	return heap.Sort(data)
}

/**
bitmap排序
*/
func BitmapSort(data []int) []int {
	temp := make([]byte, math.MaxInt32)
	for i := 0; i < len(data); i++ {
		index := data[i] / 8
		offset := data[i] % 8
		temp[index] = temp[index] | (0x01 << uint(offset))
	}
	res := make([]int, len(data))
	rIndex := 0
	for i := 0; i < len(temp); i++ {
		for j := 0; j < 8; j++ {
			if (temp[i] >> uint(j) & 0x01) == 1 {
				res[rIndex] = i*8 + j
				rIndex++
			}
		}
	}
	return res
}

/**
希尔排序
*/
func ShellsSort(data []int) {

}

/**
位排序
*/
func BitSort(data []int) {

}

/**
桶排序
*/
func BucketSort(data []int) {

}
