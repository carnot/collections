package collections

import (
	"fmt"
)

type (
	node struct {
		key byte
		value Any
		left, middle, right *node
	}
	iterator struct {
		step int
		node *node
		prev *iterator
	}
	TernaryTree struct {
		length int
		root *node
	}	
)

// Create a new ternary search tree
func NewTernaryTree() *TernaryTree {
	tree := &TernaryTree{}
	tree.Init()
	return tree
}
// Get the value at the specified key. Returns nil if not found.
func (this *TernaryTree) Get(key string) Any {
	if this.length == 0 {
		return nil
	}
	
	node := this.root
	bs := []byte(key)
	for i := 0; i < len(bs); {
		b := bs[i]
		if b > node.key {
			if node.right == nil {
				return nil
			}
			node = node.right
		} else if (b < node.key) {
			if node.left == nil {
				return nil
			}
			node = node.left
		} else {
			i++
			if i < len(bs) {
				if node.middle == nil {
					return nil
				}
				node = node.middle
			} else {
				break
			}
		}
	}
	return node.value
}
// Test to see whether or not the given key is contained in the tree.
func (this *TernaryTree) Has(key string) bool {
	return this.Get(key) != nil
}
// Initialize the tree (reset it so that its empty). New will do this for you.
func (this *TernaryTree) Init() {
	this.length = 0
	this.root = nil
}
// Insert a new key value pair into the collection
func (this *TernaryTree) Insert(key string, value Any) {
	// If the value is nil then remove this key from the collection
	if value == nil {
		this.Remove(key)
		return
	}

	if this.length == 0 {
		this.root = &node{0,nil,nil,nil,nil}
	}
	
	t := this.root
	bs := []byte(key)
	for i := 0; i < len(bs); {
		b := bs[i]
		if b > t.key {
			if t.right == nil {
				t.right = &node{b,nil,nil,nil,nil}
			}
			t = t.right
		} else if b < t.key {
			if t.left == nil {
				t.left = &node{b,nil,nil,nil,nil}
			}
			t = t.left
		} else {
			i++
			if i < len(bs) {
				if t.middle == nil {
					t.middle = &node{bs[i],nil,nil,nil,nil}
				}
				t = t.middle				
			}
		}
	}
	
	if t.value == nil {
		this.length++
	}
	t.value = value
}
// Iterate over the collection
func (this *TernaryTree) Do(f func(Any)bool) {
	if this.Len() == 0 {
		return
	}
	i := &iterator{0,this.root,nil}
	for i != nil {
		switch i.step {
		// Left
		case 0:
			i.step++
			if i.node.left != nil {
				i = &iterator{0,i.node.left,i}
				continue
			}
		// Value
		case 1:
			i.step++
			if i.node.value != nil {
				if !f(i.node.value) {
					break
				}
				continue
			}
		// Middle
		case 2:
			i.step++
			if i.node.middle != nil {
				i = &iterator{0,i.node.middle,i}
				continue
			}
		// Right
		case 3:
			i.step++
			if i.node.right != nil {
				i = &iterator{0,i.node.right,i}
				continue
			}
		// Backtrack
		case 4:
			i = i.prev
		}
	}
}
// Dump the tree to a string for easier debuggin
func (this *node) String() string {
	str := "{" + string(this.key)
	if this.value != nil {
		str += ":" + fmt.Sprint(this.value)
	}
	if this.left != nil {
		str += this.left.String()
	} else {
		str += " "
	}
	if this.middle != nil {
		str += this.middle.String()
	} else {
		str += " "
	}
	if this.right != nil {
		str += this.right.String()
	} else {
		str += " "
	}
	str += "}"
	return str
}
func (this *TernaryTree) String() string {
	if this.length == 0 {
		return "{}"
	}
	
	return this.root.String()
}
// Get the number of items stored in the tree
func (this *TernaryTree) Len() int {
	return this.length
}
// Remove a key from the collection
func (this *TernaryTree) Remove(key string) Any {
	if this.length == 0 {
		return nil
	}
	
	var remove *node
	var direction int
	
	t := this.root
	bs := []byte(key)
	for i := 0; i < len(bs); {	
		b := bs[i]
		if b > t.key {
			// Not in the collection
			if t.right == nil {
				return nil
			}
			// This is a branch so we have to keep it
			remove = t
			direction = 1
			// Move to the next node
			t = t.right
		} else if b < t.key {
			// Not in the collection
			if t.left == nil {
				return nil
			}
			// This is a branch so we have to keep it
			remove = t
			direction = -1
			// Move to the next node
			t = t.left
		} else {
			i++
			if i < len(bs) {
				// Not in the collection
				if t.middle == nil {
					return nil
				}
				// Has a value so we need to keep at least this much
				if t.value != nil {
					remove = t
					direction = 0
				}
				// Move to the next node
				t = t.middle
			}
		}
	}
	
	// If this was the only item in the tree, set the root pointer to nil
	if this.length == 1 {
		this.root = nil
	} else {
		if direction == -1 {
			remove.left = nil
		} else if direction == 0 {
			remove.middle = nil
		} else {
			remove.right = nil
		}
	}
	this.length--
	return t.value
}