package collections

import (
	"fmt"
)

type (
	splayNode struct {
		key Any
		value Any
		parent, left, right *splayNode
	}
	splayNodeIterator struct {
		step int
		node *splayNode
		prev *splayNodeIterator
	}
	
	SplayTree struct {
		length int
		root *splayNode
		less func(Any,Any)bool
	}
)

func (this *splayNode) rotateLeft() {
	parent := this.parent
	pivot := this.right
	child := pivot.left
	
	if pivot == nil {
		return
	}
	
	// Update the parent
	if parent != nil {
		if parent.left == this {
			parent.left = pivot
		} else {
			parent.right = pivot
		}
	}
	
	// Update the pivot
	pivot.parent = parent
	pivot.left = this
	
	// Update the child
	if child != nil {
		child.parent = this
	}
	
	// Update this
	this.parent = pivot
	this.right = child
}
func (this *splayNode) rotateRight() {
	parent := this.parent
	pivot := this.left
	child := pivot.right
	
	if pivot == nil {
		return
	}
	
	// Update the parent
	if parent != nil {
		if parent.left == this {
			parent.left = pivot
		} else {
			parent.right = pivot
		}
	}
	
	// Update the pivot
	pivot.parent = parent
	pivot.right = this
	
	if child != nil {
		child.parent = this
	}
	
	// Update this
	this.parent = pivot
	this.left = child	
}

func (this *SplayTree) splay(node *splayNode) {
	// Already root, nothing to do
	if node.parent == nil {
		this.root = node
		return
	}
	
	p := node.parent
	g := p.parent
	
	// Zig
	if p == this.root {
		if node == p.left {
			p.rotateRight()
		} else {
			p.rotateLeft()
		}
	} else {		
		// Zig-zig
		if node == p.left && p == g.left {
			g.rotateRight()
			p.rotateRight()
		} else if node == p.right && p == g.right {
			g.rotateLeft()
			p.rotateLeft()
		// Zig-zag
		} else if node == p.right && p == g.left {
			p.rotateLeft()
			g.rotateRight()
		} else if node == p.left && p == g.right {
			p.rotateRight()
			g.rotateLeft()
		}
	}
	this.splay(node)
}
// Swap two nodes in the tree
func (this *SplayTree) swap(n1, n2 *splayNode) {
	p1 := n1.parent
	l1 := n1.left
	r1 := n1.right
	
	p2 := n2.parent
	l2 := n2.left
	r2 := n2.right
	
	// Update node links
	n1.parent = p2
	n1.left = l2
	n1.right = r2
	
	n2.parent = p1
	n2.left = l1
	n2.right = r1
	
	// Update parent links
	if p1 != nil {
		if p1.left == n1 {
			p1.left = n2
		} else {
			p1.right = n2
		}
	}	
	if p2 != nil {
		if p2.left == n2 {
			p2.left = n1
		} else {
			p2.right = n1
		}
	}
	
	if n1 == this.root {
		this.root = n2
	} else if n2 == this.root {
		this.root = n1
	}
}

func (this *splayNode) String() string {
	str := "{" + fmt.Sprint(this.key) + ":" + fmt.Sprint(this.value) + "|"
	if this.left != nil {
		str += this.left.String()
	}
	str += "|"
	if this.right != nil {
		str += this.right.String()
	}
	str += "}"
	return str
}

func NewSplayTree(less func(Any,Any)bool) *SplayTree {
	return &SplayTree{0,nil,less}
}
func (this *SplayTree) Get(key Any) Any {
	if this.length == 0 {
		return nil
	}
	
	node := this.root
	for node != nil {
		if this.less(key, node.key) {
			node = node.left
			continue
		}
		
		if this.less(node.key, key) {
			node = node.right
			continue
		}
		
		this.splay(node)
		return node.value
	}
	return nil
}
func (this *SplayTree) Has(key Any) bool {
	return this.Get(key) != nil
}
func (this *SplayTree) Init() {
	this.length = 0
	this.root = nil
}
func (this *SplayTree) Insert(key Any, value Any) {
	if this.length == 0 {
		this.root = &splayNode{key,value,nil,nil,nil}
		this.length = 1
		return
	}
	
	node := this.root
	for {
		if this.less(key, node.key) {
			if node.left == nil {
				node.left = &splayNode{key,value,node,nil,nil}
				this.length++
				node = node.left
				break
			}
			node = node.left
			continue
		}
		
		if this.less(node.key, key) {
			if node.right == nil {
				node.right = &splayNode{key,value,node,nil,nil}
				this.length++
				node = node.right
				break
			}
			node = node.right
			continue
		}
		
		node.value = value
		break
	}
	this.splay(node)
}
func (this *SplayTree) PreOrder(f func(Any,Any)bool) {
	if this.length == 1 {
		return
	}
	i := &splayNodeIterator{0,this.root,nil}
	for i != nil {
		switch i.step {
		// Value
		case 0:
			i.step++
			if !f(i.node.key,i.node.value) {
				break
			}
		// Left
		case 1:
			i.step++
			if i.node.left != nil {
				i = &splayNodeIterator{0,i.node.left,i}
			}
		// Right
		case 2:
			i.step++
			if i.node.right != nil {
				i = &splayNodeIterator{0,i.node.right,i}
			}
		default:
			i = i.prev
		}
	}
}
func (this *SplayTree) InOrder(f func(Any,Any)bool) {
	if this.length == 1 {
		return
	}
	i := &splayNodeIterator{0,this.root,nil}
	for i != nil {
		switch i.step {
		// Left
		case 0:
			i.step++
			if i.node.left != nil {
				i = &splayNodeIterator{0,i.node.left,i}
			}
		// Value
		case 1:
			i.step++
			if !f(i.node.key,i.node.value) {
				break
			}
		// Right
		case 2:
			i.step++
			if i.node.right != nil {
				i = &splayNodeIterator{0,i.node.right,i}
			}
		default:
			i = i.prev
		}
	}
}
func (this *SplayTree) PostOrder(f func(Any,Any)bool) {
	if this.length == 1 {
		return
	}
	i := &splayNodeIterator{0,this.root,nil}
	for i != nil {
		switch i.step {
		// Left
		case 0:
			i.step++
			if i.node.left != nil {
				i = &splayNodeIterator{0,i.node.left,i}
			}
		// Right
		case 1:
			i.step++
			if i.node.right != nil {
				i = &splayNodeIterator{0,i.node.right,i}
			}
		// Value
		case 2:
			i.step++
			if !f(i.node.key,i.node.value) {
				break
			}
		default:
			i = i.prev
		}
	}
}
func (this *SplayTree) Do(f func(Any,Any)bool) {
	this.InOrder(f)
}
func (this *SplayTree) Len() int {
	return this.length
}
func (this *SplayTree) Remove(key Any) {
	if this.length == 0 {
		return
	}
	
	node := this.root
	for node != nil {
		if this.less(key, node.key) {
			node = node.left
			continue
		}
		if this.less(node.key, key) {
			node = node.right
			continue
		}
		
		// First splay the parent node
		if node.parent != nil {
			this.splay(node.parent)
		}
		
		// No children
		if node.left == nil && node.right == nil {
			// guess we're the root node
			if node.parent == nil {
				this.root = nil
				break
			}
			if node.parent.left == node {
				node.parent.left = nil
			} else {
				node.parent.right = nil
			}
		} else if node.left == nil {
			// root node
			if node.parent == nil {
				this.root = node.right
				break
			}
			if node.parent.left == node {
				node.parent.left = node.right
			} else {
				node.parent.right = node.right
			}
		} else if node.right == nil {
			// root node
			if node.parent == nil {
				this.root = node.left
				break
			}
			if node.parent.left == node {
				node.parent.left = node.left
			} else {
				node.parent.right = node.left
			}
		} else {
			// find the successor
			s := node.right
			for s.left != nil {
				s = s.left
			}
			
			np := node.parent
			nl := node.left
			nr := node.right
			
			sp := s.parent
			sr := s.right
			
			// Update parent
			s.parent = np
			if np == nil {
				this.root = s
			} else {
				if np.left == node {
					np.left = s
				} else {
					np.right = s
				}
			}
			
			// Update left
			s.left = nl
			s.left.parent = s
			
			// Update right
			if nr != s {
				s.right = nr
				s.right.parent = s
			}
			
			// Update successor parent
			if sp.left == s {
				sp.left = sr
			} else {
				sp.right = sr
			}			
		}
			
		break
	}
	
	if node != nil {
		this.length--
	}
}
func (this *SplayTree) String() string {
	if this.length == 0 {
		return "{}"
	}
	return this.root.String()
}