package search

import (
	"errors"
)

/**
1.节点是红色或黑色。
2.根是黑色。
3.所有叶子都是黑色（叶子是NIL节点）。
4.每个红色节点必须有两个黑色的子节点。（从每个叶子到根的所有路径上不能有两个连续的红色节点。）
5.从任一节点到其每个叶子的所有简单路径都包含相同数目的黑色节点。
*/

const (
	RED   = false
	BLACK = true
)

type RBTree struct {
	size     int64
	rootNode *treeNode
}

type treeNode struct {
	key        int64
	value      int64
	status     bool // false：红 true：黑
	parent     *treeNode
	leftChild  *treeNode
	rightChild *treeNode
}

func NewRBTree() RBTree {
	return RBTree{
		size:     0,
		rootNode: nil,
	}
}

func (tree *RBTree) Insert(key int64, value int64) {
	// 根节点
	if tree.rootNode == nil {
		tree.rootNode = &treeNode{ // 性质2
			key:        key,
			value:      value,
			status:     BLACK,
			parent:     nil,
			leftChild:  nil,
			rightChild: nil,
		}
		return
	}

	cntNode := tree.rootNode
	for cntNode != nil {
		if key < cntNode.key && cntNode.leftChild == nil {
			cntNode.leftChild = &treeNode{
				key:        key,
				value:      value,
				status:     BLACK,
				parent:     cntNode,
				leftChild:  nil,
				rightChild: nil,
			}
			return
		} else if key > cntNode.key && cntNode.rightChild == nil {
			cntNode.rightChild = &treeNode{
				key:        key,
				value:      value,
				status:     BLACK,
				parent:     cntNode,
				leftChild:  nil,
				rightChild: nil,
			}
			return
		} else if key < cntNode.key && cntNode.leftChild != nil {
			cntNode = cntNode.leftChild
		} else if key > cntNode.key && cntNode.rightChild != nil {
			cntNode = cntNode.rightChild
		} else if key == cntNode.key {
			cntNode.value = value
			return
		}
	}
}

func (tree *RBTree) Delete(data int) {

}

func (tree *RBTree) Get(key int64) (int64, error) {
	cntNode := tree.rootNode
	for cntNode != nil {
		if key < cntNode.key {
			cntNode = cntNode.leftChild
		} else if key > cntNode.key {
			cntNode = cntNode.rightChild
		} else {
			return cntNode.value, nil
		}
	}
	return 0, errors.New("value is nil")
}

func grandparent(n *treeNode) *treeNode {
	return n.parent.parent
}

func uncle(n *treeNode) *treeNode {
	if n.parent == grandparent(n).leftChild {
		return grandparent(n).rightChild
	} else {
		return grandparent(n).leftChild
	}
}
