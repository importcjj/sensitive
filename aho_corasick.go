package sensitive

// ac 自动机
type ac struct {
	results []string
}

func (ac *ac) fail(node *Node, c rune) *Node {
	var next *Node
	for {
		next = ac.next(node.Failure, c)
		if next == nil {
			if node.IsRootNode() {
				return node
			}
			node = node.Failure
			continue
		}

		return next
	}

}

func (ac *ac) next(node *Node, c rune) *Node {
	next, ok := node.Children[c]
	if ok {
		return next
	}
	return nil
}

func (ac *ac) output(node *Node, runes []rune, position int) {
	if node.IsRootNode() {
		return
	}

	if node.IsPathEnd() {
		ac.results = append(ac.results, string(runes[position+1-node.depth:position+1]))
	}

	ac.output(node.Failure, runes, position)
}

func (ac *ac) firstOutput(node *Node, runes []rune, position int) string {
	if node.IsRootNode() {
		return ""
	}

	if node.IsPathEnd() {
		return string(runes[position+1-node.depth : position+1])
	}

	return ac.firstOutput(node.Failure, runes, position)
}

func (ac *ac) replace(node *Node, runes []rune, position int, replace rune) {
	if node.IsRootNode() {
		return
	}

	if node.IsPathEnd() {
		for i := position + 1 - node.depth; i < position+1; i++ {
			runes[i] = replace
		}
	}
	ac.replace(node.Failure, runes, position, replace)
}
