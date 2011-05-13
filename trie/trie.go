package trie

import (
	//"fmt"
)

type (
	Action func([]byte,Any)bool
	Any interface{}
	
	Trie struct {
		root *node
		length int
	}
	node struct {
		value Any
		key byte
		children [256]*node
	}
	iterator struct {
		step int
		node *node
		prev *iterator
	}
)

func New() *Trie {
	return &Trie{nil,0}
}
func (this *Trie) Do(action Action) {
	InOrder(action)
}
func (this *Trie) Get(key []byte) Any {

}
func (this *Trie) Has(key []byte) bool {

}
func (this *Trie) Init() {
	this.root = nil
	this.length = 0
}
func (this *Trie) InOrder(action Action) {

}
func (this *Trie) Insert(key []byte, value Any) {

}
func (this *Trie) Len() int {
	return this.length
}
func (this *Trie) PostOrder(action Action) {

}
func (this *Trie) PreOrder(action Action) {
	if this.length == 0 {
		return
	}
	
	d := 0
	i := &iterator{0,this.root,nil}
	for i != nil {
		if i.node.value != nil {
			bs := make([]byte, d)
			t := i
			for j := d; j >= 0; j-- {
				bs[j] = t.node.key
				t = i.prev
			}
			action(bs, i.node.value)
		}
		for i.step < 256; i.step++ {
			c := i.node.children[j]
			if c != nil {
				i = &iterator{0,c,i}
				d++
				break
			}
		}
		i = i.prev
		d--
	}	
}
func (this *Trie) Remove() {
	
}




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