package collections

import (
	"fmt"
	"testing"
)

func TestSplay(t *testing.T) {
	tree := NewSplayTree(func(a,b Any)bool {
		return a.(string) < b.(string)
	})
	
	tree.Insert("d", 4)
	fmt.Println(tree)
	tree.Insert("b", 2)
	fmt.Println(tree)
	tree.Insert("a", 1)
	fmt.Println(tree)
	tree.Insert("c", 3)	
	fmt.Println(tree)
	tree.Remove("b")
	fmt.Println(tree)
	
	/*n1 := &splayNode{0,0,nil,nil,nil}
	n2 := &splayNode{1,1,nil,nil,nil}
	n3 := &splayNode{2,2,nil,nil,nil}
	n4 := &splayNode{3,3,nil,nil,nil}
	n5 := &splayNode{4,4,nil,nil,nil}
	n6 := &splayNode{5,5,nil,nil,nil}
	n7 := &splayNode{6,6,nil,nil,nil}
	n1.left = n2
	n2.parent = n1
	
	n1.right = n3
	n3.parent = n1
	
	n2.left = n4
	n4.parent = n2
	
	n2.right = n5
	n5.parent = n2
	
	n3.left = n6
	n6.parent = n3
	
	n3.right = n7
	n7.parent = n3
	
	n1.rotateLeft()
	fmt.Println(n3)*/
}