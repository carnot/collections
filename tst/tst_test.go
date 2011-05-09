package tst

import (
	//"fmt"
	"testing"
)

func TestCases(t *testing.T) {
	tree := New()
	tree.Insert("test", 1)
	if tree.Len() != 1 {
		t.Errorf("expecting len 1")
	}
	if !tree.Has("test") {
		t.Errorf("expecting to find key=test")
	}
	
	tree.Insert("testing", 2)
	tree.Insert("abcd", 0)
		
	found := false
	tree.Do(func(test Any)bool {
		if test.(int) == 1 {
			found = true
		}
		return true
	})
	if !found {
		t.Errorf("expecting iterator to find test")
	}
	
	tree.Remove("testing")
	tree.Remove("abcd")
	
	v := tree.Remove("test")
	if tree.Len() != 0 {
		t.Errorf("expecting len 0")
	}
	if tree.Has("test") {
		t.Errorf("expecting not to find key=test")
	}
	if v.(int) != 1 {
		t.Errorf("expecting value=1")
	}
}