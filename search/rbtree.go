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
	key    uint64
	value  string
	status bool // false：红 true：黑
	parent *treeNode
	left   *treeNode
	right  *treeNode
}

func NewRBTree() RBTree {
	return RBTree{
		size:     0,
		rootNode: nil,
	}
}

func (tree *RBTree) Insert(key uint64, value string) {
	// 根节点
	if tree.rootNode == nil {
		tree.rootNode = &treeNode{ // 性质2
			key:    key,
			value:  value,
			status: BLACK,
			parent: nil,
			left:   nil,
			right:  nil,
		}
		return
	}

	cntNode := tree.rootNode
	for cntNode != nil {
		if key < cntNode.key && cntNode.left != nil {
			cntNode = cntNode.left
			break
		} else if key < cntNode.key && cntNode.left == nil {
			// todo 左新节点
			cntNode.left = &treeNode{
				key:    key,
				value:  value,
				status: BLACK,
				parent: cntNode,
				left:   nil,
				right:  nil,
			}
			return
		} else if key > cntNode.key && cntNode.right != nil {
			cntNode = cntNode.right
			break
		} else if key > cntNode.key && cntNode.right == nil {
			// todo 右新节点
			cntNode.right = &treeNode{
				key:    key,
				value:  value,
				status: BLACK,
				parent: cntNode,
				left:   nil,
				right:  nil,
			}
			return
		} else if key == cntNode.key {
			// 更新值
			cntNode.value = value
			return
		}
	}
}

func (tree *RBTree) Delete(data int) {

}

func (tree *RBTree) Get(key uint64) (string, error) {
	cntNode := tree.rootNode
	for cntNode != nil {
		if key < cntNode.key {
			cntNode = cntNode.left
		} else if key > cntNode.key {
			cntNode = cntNode.left
		} else {
			return cntNode.value, nil
		}
	}
	return "", errors.New("value is nil")
}

func grandparent(n *treeNode) *treeNode {
	return n.parent.parent
}

func uncle(n *treeNode) *treeNode {
	if n.parent == grandparent(n).left {
		return grandparent(n).right
	} else {
		return grandparent(n).left
	}
}
