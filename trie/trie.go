package trie

func main() {
	trie := Constructor()
	trie.Insert("apple")
	trie.Search("apple")
	trie.Search("app")
	trie.StartsWith("app")
	trie.Insert("app")
	trie.Search("app")
}

type Trie struct {
	value    byte
	children []*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{
		value:    0,
		children: []*Trie{},
	}
}

func (this *Trie) findChild(w byte) *Trie {
	for _, child := range this.children {
		if child.value == w {
			return child
		}
	}
	return nil
}

func (this *Trie) Insert(word string) {
	node := this
	for i := 0; i < len(word); i++ {
		w := word[i]

		child := node.findChild(w)

		if child == nil {
			child = &Trie{
				value:    w,
				children: []*Trie{},
			}
			node.children = append(node.children, child)
		}
		node = child
	}

	node.isEnd = true
}

func (this *Trie) Search(word string) bool {
	result := this.findTarget(word)

	if result == nil {
		return false
	}
	return result.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	return this.findTarget(prefix) != nil
}

func (this *Trie) findTarget(target string) *Trie {
	node := this
	for i := 0; i < len(target); i++ {
		w := target[i]
		child := node.findChild(w)
		if child == nil {
			return nil
		}
		node = child
	}

	return node
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
