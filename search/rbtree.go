package search

import "errors"

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
	LEFT  = false
	RIGHT = true
)

type RBTree struct {
	size     int64
	rootNode *treeNode
}

type treeNode struct {
	key        int64
	value      int64
	color      bool // false：红 true：黑
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
			color:      BLACK,
			parent:     nil,
			leftChild:  nil,
			rightChild: nil,
		}
		return
	}

	currentNode := tree.rootNode
	for currentNode != nil {
		if key < currentNode.key && currentNode.leftChild == nil {
			var newNode = treeNode{
				key:        key,
				value:      value,
				color:      BLACK,
				parent:     currentNode,
				leftChild:  nil,
				rightChild: nil,
			}
			currentNode.leftChild = &newNode
			return
		} else if key > currentNode.key && currentNode.rightChild == nil {
			var newNode = treeNode{
				key:        key,
				value:      value,
				color:      BLACK,
				parent:     currentNode,
				leftChild:  nil,
				rightChild: nil,
			}
			currentNode.rightChild = &newNode
			return
		} else if key < currentNode.key && currentNode.leftChild != nil {
			currentNode = currentNode.leftChild
		} else if key > currentNode.key && currentNode.rightChild != nil {
			currentNode = currentNode.rightChild
		} else if key == currentNode.key {
			currentNode.value = value
			return
		}
	}
}

func (tree *RBTree) Delete(data int) {

}

func (tree *RBTree) Get(key int64) (int64, error) {
	currentNode := tree.rootNode
	for currentNode != nil {
		if key < currentNode.key {
			currentNode = currentNode.leftChild
		} else if key > currentNode.key {
			currentNode = currentNode.rightChild
		} else {
			return currentNode.value, nil
		}
	}
	return 0, errors.New("value is nil")
}

func insertAdjustOne(node *treeNode) {
	if node.parent == nil { // insert_case2 父节点为黑色，新节点为红色来保持性质5

	} else { // insert_case3 父节点为红色，证明肯定存在祖父节点（根节点为黑色）

	}
}

/**
左旋
*/
func (tree *RBTree) rotateLeft(node *treeNode) {
	if node.parent == nil { // 自己是根节点，不执行左旋操作
		return
	}

	parent := node.parent
	grandparent := grandparent(node)
	leftChild := node.leftChild

	parent.rightChild = leftChild
	parent.parent = node
	if leftChild != nil {
		leftChild.parent = parent
	}

	node.leftChild = parent
	node.parent = grandparent
	if grandparent == nil { // 判断父节点是否是根节点
		tree.rootNode = node
	} else {
		if grandparent.leftChild == parent {
			grandparent.leftChild = node
		} else {
			grandparent.rightChild = node
		}
	}
}

/**
右旋
*/
func (tree *RBTree) rotateRight(node *treeNode) {
	if node.parent == nil { // 自己是根节点，不执行右旋操作
		return
	}

	parent := node.parent
	grandparent := grandparent(node)
	rightChild := node.rightChild

	parent.leftChild = rightChild
	parent.parent = node
	if rightChild != nil {
		rightChild.parent = parent
	}

	node.rightChild = parent
	node.parent = grandparent
	if grandparent == nil {
		tree.rootNode = node
	} else {
		if grandparent.leftChild == parent {
			grandparent.leftChild = node
		} else {
			grandparent.rightChild = node
		}
	}
}

/**
获取祖父节点
*/
func grandparent(n *treeNode) *treeNode {
	return n.parent.parent
}

/**
获取叔父节点
*/
func uncle(n *treeNode) *treeNode {
	if n.parent == grandparent(n).leftChild {
		return grandparent(n).rightChild
	} else {
		return grandparent(n).leftChild
	}
}
