package tree

import (
	"algorithm/sort"
	"fmt"
)

/**
二叉堆
1.每个父节点均大于子节点
2.左子节点(奇数) left=2*parent+1, 右子节点(偶数) right=2*parent+2
3.第一个非叶子节点 arr.length/2-1
*/
type BinaryHeap struct {
	array []int
	cap   int
	size  int
}

func NewBinaryHeap(cap int) *BinaryHeap {
	return &BinaryHeap{make([]int, cap), cap, 0}
}

/**
最大堆向下调整
*/
func (heap *BinaryHeap) filterDown(index int) {
	if (2*index + 1) < heap.size {
		temp := 0
		if heap.array[2*index+2] >= heap.size || heap.array[2*index+1] > heap.array[2*index+2] {
			temp = 2*index + 1
		} else {
			temp = 2*index + 2
		}
		if heap.array[index] < heap.array[temp] {
			sort.Swap(&heap.array[index], &heap.array[temp])
			heap.filterDown(temp)
		}
	}
}

/**
最大堆向上调整
*/
func (heap *BinaryHeap) filterUp(index int) {
	parent := 0
	if index&1 == 1 {
		parent = (index - 1) / 2
	} else if index&1 == 0 && index != 0 {
		parent = (index - 2) / 2
	}
	if heap.array[parent] < heap.array[index] {
		sort.Swap(&heap.array[parent], &heap.array[index])
		heap.filterUp(parent)
	}
}

/**
返回data在二叉堆中的索引
*/
func (heap *BinaryHeap) GetIndex(data int) {
}

/**
删除最大堆中的data
*/
func (heap *BinaryHeap) Remove(index int) {
	sort.Swap(&heap.array[index], &heap.array[heap.size-1])
	heap.filterDown(index)
	heap.size--
}

/**
将data插入到二叉堆中
*/
func (heap *BinaryHeap) Insert(data int) bool {
	if heap.size < heap.cap {
		heap.array[heap.size] = data
		heap.filterUp(heap.size)
		heap.size++
		return true
	} else {
		return false
	}
}

/**
打印二叉堆
*/
func (heap *BinaryHeap) Print() []int {
	for i := 0; i < heap.size; i++ {
		fmt.Println(heap.array[i])
	}
	return heap.array[0:heap.size]
}
