package collections

import (
	//"fmt"
)

type (
	Trie struct {
		value interface{}
		children map[byte]*Trie
	}
)

// Add an entry to the trie
func (this *Trie) Add(key string, value interface{}) {
	t := this
	k := []byte(key)
	for _, b := range k {
		if t.children == nil {
			t.children = make(map[byte]*Trie)
		}
		n, ok := t.children[b]
		if !ok {
			n = &Trie{nil,nil}
			t.children[b] = n
		}
		t = n
	}
	t.value = value
}

// Find an entry in the trie
func (this *Trie) Get(key string, exact bool) interface{} {
	t := this
	var v interface{}
	var ok bool
	k := []byte(key)
	for _, b := range k {
		if t.value != nil {
			v = t.value
		}
		if t.children == nil {
			if exact { 
				return nil
			} else {
				return v
			}
		}
		t, ok = t.children[b]
		if !ok {
			if exact {
				return nil
			} else {
				return v
			}
		}
	}
	return t.value
}