package search

type RBTree struct {
	size     int64
	rootNode *node
}

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
		tree.rootNode = &node{
			data:   data,
			status: true,
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
