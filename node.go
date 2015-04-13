package trog

type rawNode struct {
	Id        NodeId
	ParentIds []NodeId // Every ancestor id, all the way up to the root
	ChildIds  []NodeId // Every direct child id of this node
	Parents   []*Node  // Every ancestor, all the way up to the root
	Children  []*Node  // Every direct child of this node
}

type Node struct {
	Pairs    []Pair  // Key/value pairs
	Flags    []Flag  // Value-only flags
	Parent   *Node   // Direct ancestor of this node
	Children []*Node // Direct children of this node
	Parents  []*Node // Every ancestor, all the way up to the root
}

// Create a new, empty, node.
func NewNode() *Node {
	return &Node{}
}

// Add a new node between the given node and its parent.
func (node *Node) Split() *Node {
	newNode := NewNode()
	parent := node.Parent

	if parent != nil {
		parent.delChild(node)
		parent.addChild(newNode)
	}
	newNode.addChild(node)
	return newNode
}

// Add a new child to the given node.
func (node *Node) AddChild() *Node {
	newNode := NewNode()
	node.addChild(newNode)
	return newNode
}

func (node *Node) addChild(child *Node) {
	child.Parent = node
	node.Children = append(node.Children, child)
}

func (node *Node) delChild(child *Node) {
	for i := 0; i < len(node.Children); i++ {
		if node.Children[i] == child {
			node.Children[i] = node.Children[len(node.Children)-1]
			node.Children = node.Children[:len(node.Children)-1]
			child.Parent = nil
		}
	}
}
