package search

/**
1.节点是红色或黑色。
2.根是黑色。
3.所有叶子都是黑色（叶子是NIL节点）。
4.每个红色节点必须有两个黑色的子节点。（从每个叶子到根的所有路径上不能有两个连续的红色节点。）
5.从任一节点到其每个叶子的所有简单路径都包含相同数目的黑色节点。
*/

type RBTree struct {
	size     int64
	rootNode *node
}

const (
	RED   = false
	BLACK = true
)

type node struct {
	data   int
	status bool // false：红 true：黑
	parent *node
	left   *node
	right  *node
}

func NewRBTree() RBTree {
	return RBTree{
		size:     0,
		rootNode: nil,
	}
}

func (tree *RBTree) Insert(data int) {
	if tree.rootNode == nil {
		tree.rootNode = &node{ // 性质2
			data:   data,
			status: BLACK,
			parent: nil,
			left:   nil,
			right:  nil,
		}
		return
	}
}

func (tree *RBTree) Delete(data int) {

}

func grandParent(n *node) *node {
	return n.parent.parent
}

func uncle(n *node) *node {
	if n.parent == grandParent(n).left {
		return grandParent(n).right
	} else {
		return grandParent(n).left
	}
}
