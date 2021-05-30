package red_black_tree

type Node struct {
	parent      *Node
	left, right *Node
	val         int64
	color       COLOR
}

func NewNode(parent, left, right *Node, val int64, color COLOR) *Node {
	return &Node{
		parent: parent,
		left:   left,
		right:  right,
		val:    val,
		color:  color,
	}
}

func (n *Node) GrandParent() *Node {
	if n == nil || n.parent == nil {
		return nil
	}

	return n.parent.parent
}

func (n *Node) Uncle() *Node {
	grantParent := n.GrandParent()
	if grantParent == nil {
		return nil
	}

	if grantParent.left == n.parent {
		return grantParent.right
	} else {
		return grantParent.left
	}
}

func (n *Node) Sibling() *Node {
	if n == nil || n.parent == nil {
		return nil
	}

	if n == n.parent.left {
		return n.parent.right
	} else {
		return n.parent.left
	}
}
