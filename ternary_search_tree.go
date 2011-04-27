package collections

type (
	node struct {
		key byte
		value interface{}
		left, middle, right *node
	}
	TernarySearchTree struct {
		size int
		root *node
	}
)
// Add a value to the tree at the specified key
func (this *TernarySearchTree) Add(key string, value interface{}) bool {	
	k := []byte(key)
	// If the key length is 0 don't add it
	if len(k) == 0 {
		return false
	}
	
	t := this.root
	// If we don't have a root yet
	if t == nil {
		t = &node{k[0],nil,nil,nil,nil}
		this.root = t
	}
	
	for _, b := range k {
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
			if t.middle == nil {
				t.middle = &node{b,nil,nil,nil,nil}
			}
			t = t.middle
		}
	}
	t.value = value
	this.size++
	return true
}

func (this *TernarySearchTree) Get(key string, exact bool) interface{} {
	k := []byte(key)
	
	if len(k) == 0 {
		return nil
	}
	
	t := this.root
	if t == nil {
		return nil
	}
	
	var v interface{}
	for _, b := range k {
		if t.value != nil {
			v = t.value
		}
		if b > t.key {
			t = t.right
		} else if b < t.key {
			t = t.left
		} else {
			t = t.middle
		}
		if t == nil {
			if exact {
				return nil
			} else {
				return v
			}
		}
	}
	if t.value != nil {
		v = t.value
	}
	return v
}