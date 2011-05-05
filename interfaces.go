package collections

import (
	//"reflect"
)

type (
	Any interface{}
	
	Iterable interface {
		Do(func(Any)bool)
	}
	Collection interface {
		Init()
		Len() int
	}
	Indexed interface {
		At(int) Any
	}
	Dictionary interface {
		Get(key Any) Any
		Has(key Any) bool
		Insert(key Any, value Any)
		Remove(key Any) Any
	}
	Sorted interface {
		Min() Any
		Max() Any
	}
)

/*
// Call the function 'f' for each of the items in the collection.
// The object 'result' will be filled with the next value and the
// passed function will be called.
func For(this Iterable, result Any, f func()bool) {
	v := reflect.ValueOf(result)

	this.Do(func(val Any)bool{
		v.Set(reflect.ValueOf(val))
		return f()
	})
}*/