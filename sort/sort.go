package sort

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
		for j := 0; j < len(data)-i-1; j++ {
			if data[j] > data[j+1] {
				swap(&data[i], &data[j])
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
		for ; i < j; {
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
func MergeSort(data []int) {

}

/**
堆排序
 */
func HeapSort(data []int) {

}

/**
希尔排序
 */
func ShellsSort(data []int) {

}
