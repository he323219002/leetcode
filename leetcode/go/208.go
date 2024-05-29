package main

import "fmt"

func main() {
	trie := Constructor()
	trie.Insert("abcde")
	searchRes := trie.Search("abcde")
	fmt.Println(searchRes)
	startsRes := trie.StartsWith("abce")
	fmt.Println(startsRes)
}

type Trie struct {
	isEnd bool
	next  [26]*Trie
}

func Constructor() Trie {
	trie := Trie{isEnd: false}
	return trie
}

func (t *Trie) Insert(word string) {
	node := t
	for _, c := range word {
		index := c - 'a'
		if node.next[index] == nil {
			node.next[index] = &Trie{}
		}
		node = node.next[index]
	}
	node.isEnd = true
}

func (t *Trie) Search(word string) bool {
	node := t
	for _, c := range word {
		index := c - 'a'
		if node.next[index] == nil {
			return false
		}
		node = node.next[index]
	}
	return node.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	node := t
	for _, c := range prefix {
		index := c - 'a'
		if node.next[index] == nil {
			return false
		}
		node = node.next[index]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
