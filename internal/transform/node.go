package transform

// TODO put this in a subpackage and write tests for it.

// node is a tree data structure. A non-nil value indicates a leaf node.
type node struct {
	value    *string
	branches map[rune]*node
}

// putValue inserts a key/value pair into the tree, each of the key's runes
// representing a branch in the tree.
func (n *node) putValue(key []rune, value string) {

	if n == nil || len(key) == 0 {
		return
	}

	if n.branches == nil {
		n.branches = make(map[rune]*node)
	}

	targetNode := new(node)
	if existingNode, ok := n.branches[key[0]]; ok {
		targetNode = existingNode
	}

	if len(key) == 1 {
		targetNode.value = &value
	} else {
		targetNode.putValue(key[1:], value)
	}

	n.branches[key[0]] = targetNode
}

func (n *node) putNode(key []rune, newNode *node) {

	if n == nil || newNode == nil || len(key) == 0 {
		return
	}

	if n.branches == nil {
		n.branches = make(map[rune]*node)
	}

	if len(key) == 1 {
		n.branches[key[0]] = newNode
		return
	}

	n.branches[key[0]].putNode(key[1:], newNode)
}

func (n *node) getValue(key string) string {

	if n == nil || len(key) == 0 || n.branches == nil {
		return ""
	}

	keyRunes := []rune(key)

	var (
		targetNode *node
		ok         bool
	)
	if targetNode, ok = n.branches[keyRunes[0]]; !ok {
		return ""
	}

	if len(key) == 1 {
		if targetNode.value != nil {
			return *targetNode.value
		}

		return ""
	}

	return targetNode.getValue(key[1:])
}

func (n *node) getNode(key []rune) *node {

	if n == nil || len(key) == 0 || n.branches == nil {
		return nil
	}

	var (
		targetNode *node
		ok         bool
	)
	if targetNode, ok = n.branches[key[0]]; !ok {
		return nil
	}

	if len(key) == 1 {
		return targetNode
	}

	return targetNode.getNode(key[1:])
}

func (n *node) copy() *node {

	if n == nil {
		return nil
	}

	branches := make(map[rune]*node)
	for k, v := range n.branches {
		branches[k] = v.copy()
	}

	return &node{
		value:    n.value,
		branches: branches,
	}
}

func (n *node) mapify() map[string]string {

	if n == nil {
		return nil
	}

	result := make(map[string]string)
	n.crawlBranches("", result)

	return result
}

func (n *node) crawlBranches(key string, m map[string]string) {

	if n == nil {
		return
	}

	if n.value != nil {
		m[key] = *n.value
	}

	for k, v := range n.branches {
		v.crawlBranches(key+string(k), m)
	}
}
