package binarytree

// BinaryTree ...
type BinaryTree struct {
	root *Node
}

// NewBinaryTree ...
func NewBinaryTree() *BinaryTree {
	return &BinaryTree{}
}

// Insert ...
func (b *BinaryTree) Insert(i int) {
	if b.root == nil {
		b.root = NewNode(i, nil, nil)
		return
	}

	b.root.Insert(i)
}

// InOrderTraversal ...
func (b *BinaryTree) InOrderTraversal(node *Node, callback func(int)) {
	if node.left != nil {
		b.InOrderTraversal(node.left, callback)
	}
	callback(node.value)
	if node.right != nil {
		b.InOrderTraversal(node.right, callback)
	}
}

// Node ...
type Node struct {
	value int
	left  *Node
	right *Node
}

// NewNode ...
func NewNode(i int, l *Node, r *Node) *Node {
	return &Node{
		value: i,
		left:  l,
		right: r,
	}
}

// Insert ...
func (n *Node) Insert(i int) {
	if i < n.value {
		if n.left == nil {
			n.left = NewNode(i, nil, nil)
			return
		}
		n.left.Insert(i)
	} else {
		if n.right == nil {
			n.right = NewNode(i, nil, nil)
			return
		}
		n.right.Insert(i)
	}
}
