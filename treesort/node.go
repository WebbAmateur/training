package treesort

type Node struct {
	val   int
	left  *Node
	right *Node
}

func NewNode(value int) *Node {
	node := Node{val: value, left: nil, right: nil}
	return &node
}

func Insert(node *Node, value int) *Node {
	// If the tree is empty, return a new Node
	if node == nil {
		return NewNode(value)
	}

	// Else recurse, down the tree
	if value < node.val {
		node.left = Insert(node.left, value)
	} else if value > node.val {
		node.right = Insert(node.right, value)
	}

	return node
}

func TreeSort(values []int) *Node {
	var root *Node = nil

	root = Insert(root, values[0])

	for index := 1; index < len(values); index++ {
		root = Insert(root, values[index])
	}

	return root
}

func TreeWalk(node *Node, values []int) []int {
	if node != nil {
		values = TreeWalk(node.left, values)
		values = append(values, node.val)
		values = TreeWalk(node.right, values)
	}

	return values
}
