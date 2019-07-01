package sort

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
冒泡排序
*/
func BubbleSort(data []int) {
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data)-i-1; j++ {
			if data[j] > data[j+1] {
				x := data[j]
				data[j] = data[j+1]
				data[j+1] = x

			}
		}
	}
}

/**
选择排序
*/
func SelectionSort() {

}
