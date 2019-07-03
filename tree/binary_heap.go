package tree

import "algorithm/sort"

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
func (heap *BinaryHeap) filterDown(start int, end int) {

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
func (heap *BinaryHeap) Remove(data int) {

}

/**
将data插入到二叉堆中
*/
func (heap *BinaryHeap) Insert(data int) {
	if heap.size < heap.cap {
		heap.array[heap.size] = data
		heap.filterUp(heap.size)
		heap.size++
	}
}

/**
打印二叉堆
*/
func (heap *BinaryHeap) Print() {

}
