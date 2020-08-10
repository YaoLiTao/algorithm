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

默认新加入的节点为红色
*/

const (
	RED   = false
	BLACK = true
)

type RBTree struct {
	size     int
	rootNode *treeNode
}

type treeNode struct {
	key        int
	value      int
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

func (tree *RBTree) Insert(key int, value int) {
	// 根节点
	if tree.rootNode == nil {
		tree.rootNode = &treeNode{
			key:        key,
			value:      value,
			color:      BLACK, // 根节点为黑色
			parent:     nil,
			leftChild:  nil,
			rightChild: nil,
		}
		tree.size++
		return
	}

	currentNode := tree.rootNode
	for currentNode != nil {
		if key < currentNode.key && currentNode.leftChild == nil {
			var newNode = treeNode{
				key:        key,
				value:      value,
				color:      RED, // 默认新加入节点为红色
				parent:     currentNode,
				leftChild:  nil,
				rightChild: nil,
			}
			currentNode.leftChild = &newNode
			tree.insertCaseOne(&newNode)
			tree.size++
			return
		} else if key > currentNode.key && currentNode.rightChild == nil {
			var newNode = treeNode{
				key:        key,
				value:      value,
				color:      RED,
				parent:     currentNode,
				leftChild:  nil,
				rightChild: nil,
			}
			currentNode.rightChild = &newNode
			tree.insertCaseOne(&newNode)
			tree.size++
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

func (tree *RBTree) insertCaseOne(node *treeNode) {
	// 如果新节点为根节点，变为黑色
	if node.parent == nil {
		node.color = BLACK
	} else {
		tree.insertCaseTwo(node)
	}
}

func (tree *RBTree) insertCaseTwo(node *treeNode) {
	// 父节点为黑色，新节点为红色，无需改变
	if node.parent.color == BLACK {
		return
	} else {
		tree.insertCaseTree(node)
	}
}

func (tree *RBTree) insertCaseTree(node *treeNode) {
	// 如果父节点和叔父节点都为红色，新加入的节点为红色，需要变色
	if uncle(node) != nil && uncle(node).color == RED {
		node.parent.color = BLACK
		uncle(node).color = BLACK
		grandparent(node).color = RED
		tree.insertCaseOne(grandparent(node))
	} else {
		tree.insertCaseFour(node)
	}
}

func (tree *RBTree) insertCaseFour(node *treeNode) {
	// 如果自身为父节点的右节点，并且父节点为祖父节点的左节点，让自身左旋，变为insertCaseFive的状态
	if node == node.parent.rightChild && node.parent == grandparent(node).leftChild {
		tree.rotateLeft(node)
		node = node.leftChild
	} else if node == node.parent.leftChild && node.parent == grandparent(node).rightChild {
		tree.rotateRight(node)
		node = node.rightChild
	}
	tree.insertCaseFive(node)
}

func (tree *RBTree) insertCaseFive(node *treeNode) {
	// 如果自身为父节点的左节点，并且父节点也为祖父节点的左节点，让父节点右旋，然后变色
	if node == node.parent.leftChild && node.parent == grandparent(node).leftChild {
		tree.rotateRight(node.parent)
		node.parent.color = BLACK
		node.parent.rightChild.color = RED
	} else if node == node.parent.rightChild && node.parent == grandparent(node).rightChild {
		tree.rotateLeft(node.parent)
		node.parent.color = BLACK
		node.parent.leftChild.color = RED
	}
}

func (tree *RBTree) Delete(key int) {
	node := tree.searchNode(key)
	if node == nil {
		return
	}

	// 找到这个节点的左子树的最大值元素，并且复制它的值，可以保持树的有序性
	var deleteNode *treeNode
	if node.leftChild != nil && node.rightChild != nil {
		deleteNode = node.leftChild
		for deleteNode.rightChild != nil {
			deleteNode = deleteNode.rightChild
		}
		node.value = deleteNode.value
	} else {
		deleteNode = node
	}

	var child *treeNode
	if node.rightChild == nil {
		child = node.leftChild
	} else {
		child = node.rightChild
	}

	// 和子节点交换位置
	child.parent = node.parent
	if node.parent != nil {
		if node == node.parent.leftChild {
			node.parent.leftChild = child
		} else {
			node.parent.rightChild = child
		}
	} else {
		tree.rootNode = child
	}

	// 如果是红色节点直接删除不需要重新着色
	if node.color == RED {
		return
	}

	// 如果子节点是红色节点，把子节点重新着色为黑色节点即可
	if child.color == RED {
		child.color = BLACK
	} else {
		tree.deleteCaseOne(child)
	}
}

func (tree *RBTree) deleteCaseOne(node *treeNode) {
	if node.parent != nil {
		tree.deleteCaseTwo(node)
	}
}

func (tree *RBTree) deleteCaseTwo(node *treeNode) {
	sib := sibling(node)
	if sib.color == RED {

	}
}

func (tree *RBTree) Get(key int) (int, error) {
	node := tree.searchNode(key)
	if node != nil {
		return node.value, nil
	} else {
		return 0, errors.New("value is nil")
	}
}

func (tree *RBTree) searchNode(key int) *treeNode {
	currentNode := tree.rootNode
	for currentNode != nil {
		if key < currentNode.key {
			currentNode = currentNode.leftChild
		} else if key > currentNode.key {
			currentNode = currentNode.rightChild
		} else {
			return currentNode
		}
	}
	return nil
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
	if grandparent != nil { // 判断父节点是否是根节点
		if grandparent.leftChild == parent {
			grandparent.leftChild = node
		} else {
			grandparent.rightChild = node
		}
	} else {
		tree.rootNode = node
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
	if grandparent != nil {
		if grandparent.leftChild == parent {
			grandparent.leftChild = node
		} else {
			grandparent.rightChild = node
		}
	} else {
		tree.rootNode = node
	}
}

/**
获取祖父节点
*/
func grandparent(node *treeNode) *treeNode {
	return node.parent.parent
}

/**
获取叔父节点
*/
func uncle(node *treeNode) *treeNode {
	if node.parent == grandparent(node).leftChild {
		return grandparent(node).rightChild
	} else {
		return grandparent(node).leftChild
	}
}

/**
获取兄弟节点
*/
func sibling(node *treeNode) *treeNode {
	if node == node.parent.leftChild {
		return node.rightChild
	}
	return node.leftChild
}
