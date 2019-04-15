package trees

// TrieTree defines a trie tree.
type TrieTree struct {
	root *TrieNode // the root node won't contain any character.
}

// TrieNode defines a trie node.
type TrieNode struct {
	val      rune          // the character
	count    int           // count counts the number of times a word appears.
	flag     bool          // flag determines whether this node represents a word or not.
	children [26]*TrieNode // because the alphabet has 26 characters, we use array to store the children.
}

// NewTrieTree creates a new trie tree with no strings
func NewTrieTree() *TrieTree {
	return &TrieTree{root: &TrieNode{rune('-'), 0, false, [26]*TrieNode{}}}
}

// Insert inserts a string into the trie tree.
func (t *TrieTree) Insert(str string) {
	node := t.root
	for i, c := range str {
		idx := int(rune(c) - 'a')
		// if this character is not in the children, create it.
		if node.children[idx] == nil {
			node.children[idx] = &TrieNode{rune(c), 0, false, [26]*TrieNode{}}
		}
		// this node represents a full word.
		if i == len(str)-1 {
			node.children[idx].count++
			node.children[idx].flag = true
		}

		// move down
		node = node.children[idx]
	}
}

// Search searches a given str against a trie tree and see whether it exists.
func (t *TrieTree) Search(str string) bool {
	if str == "" {
		return false
	}
	found, _ := t.search(t.root, str)
	return found
}

// StartsWith returns all strings that starts with the prefix.
func (t *TrieTree) StartsWith(prefix string) []string {
	// we will firstly check whether we can fid the prefix in the trie tree, if can't find then return empty string.
	node := t.root
	res := make([]string, 0)
	for i := range prefix {
		idx := int(rune(prefix[i]) - 'a')
		if node.children[idx] == nil {
			return res
		}
		node = node.children[idx]
	}

	for i := range node.children {
		t.dfs(node.children[i], prefix, &res)
	}

	return res
}

// Count counts the number of apperances of a given str.
func (t *TrieTree) Count(str string) int {
	if str == "" {
		return 0
	}
	_, count := t.search(t.root, str)
	return count
}

// GetAll returns all strings stored in the trie tree.
func (t *TrieTree) GetAll() []string {
	all := make([]string, 0)
	for i := range t.root.children {
		t.dfs(t.root.children[i], "", &all)
	}
	return all
}

func (t *TrieTree) search(node *TrieNode, str string) (found bool, count int) {
	if node == nil {
		return false, 0
	}
	idx := int(rune(str[0]) - 'a')

	// we find the character in the children.
	if node.children[idx] != nil {
		// found it, if not str reaches the end, we can check whether this node represents a word.
		if len(str) == 1 {
			if node.children[idx].flag {
				return true, node.children[idx].count
			}
			return false, 0
		}
		// if we stil have more character to match, we search further down the trie tree.
		return t.search(node.children[idx], str[1:])
	}
	// can't find, return false.
	return false, 0
}

func (t *TrieTree) dfs(node *TrieNode, prefix string, res *[]string) {
	if node == nil {
		return
	}
	prefix = prefix + string(node.val)
	// if this node represents a word, add it to the res.
	if node.flag {
		*res = append(*res, prefix)
	}
	for i := range node.children {
		t.dfs(node.children[i], prefix, res)
	}
}
