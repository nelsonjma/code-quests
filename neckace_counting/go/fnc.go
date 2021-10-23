package main

import "container/list"

/* map operation

var newprimes *list.List = _map(primes, func(i interface{}) interface{} {
	return i.(int) * 2
})
*/
func _map(l *list.List, f func(v interface{}) interface{}) *list.List {

	newL := list.New()

	for e := l.Front(); e != nil; e = e.Next() {
		newL.PushFront(f(e.Value.(interface{})))
	}

	return newL
}

/* each operation with interface to any type

_each(newprimes, func(i interface{}) {
	fmt.Println(i)
})
*/
func _each(l *list.List, f func(v interface{})) {

	for e := l.Front(); e != nil; e = e.Next() {
		f(e.Value.(interface{}))
	}
}

/* filter with interface to any type

var filteredData *list.List = _filter(primes, func(i interface{}) bool {
	if i.(int) > 2 {
		return true
	}
	return false
})
*/
func _filter(l *list.List, f func(v interface{}) bool) *list.List {
	newL := list.New()

	for e := l.Front(); e != nil; e = e.Next() {
		if f(e.Value.(interface{})) {
			newL.PushFront(e.Value)
		}
	}

	return newL
}
