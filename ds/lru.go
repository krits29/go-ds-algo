package main

import (
	"fmt"
)

// implement LRU cache
// use map to seek elements in the cache with O(1)
// use doubly linked list for maintaining the cache

type Element struct {
	Val  interface{}
	Key  interface{}
	Next *Element // next element in the linked list cache
	Prev *Element // prev element
}

type Map map[interface{}]*Element

type LRU struct {
	Size int
	Cap  int
	Head *Element
	Tail *Element
	List Map
}

func (cache *LRU) init(s int) {
	cache.Cap = s
	cache.List = make(Map)
}

// if exists, return the val, move to the top of the list
func (cache *LRU) get(k interface{}) *Element {
	if cache.List != nil {
		return cache.List[k]
	}
	return nil
}

func getNewElement(k interface{}, v interface{}) *Element {
	return &Element{Key: k, Val: v}
}

// if exists, update the value else add to the list
// move to the top of the list
// check on the size, if over capacity, remove the last element
func (cache *LRU) put(k interface{}, v interface{}) {
	e := cache.get(k)
	if e == nil { // if doesn't exists check for the capacity
		cache.maintainCapacity()
		e = getNewElement(k, v)
		cache.List[k] = e
		cache.Size++
	} else {
		e.Val = v // update the value
		cache.deleteFromList(e)
	}
	// move to the top of the list
	cache.moveTop(e)

}

func (cache *LRU) deleteFromList(e *Element) {
	p := e.Prev
	n := e.Next

	if p != nil {
		p.Next = n
	}
	if n != nil {
		n.Prev = p
	}
	if e == cache.Tail { // if tail, reassign
		cache.Tail = p
	}
	if e == cache.Head {
		cache.Head = n
	}
	e.Next = nil
	e.Prev = nil
}

func (cache *LRU) maintainCapacity() {

	if cache.Size >= cache.Cap {
		//remove the tail element
		e := cache.Tail
		if e != nil {
			if e.Prev != nil {
				e.Prev.Next = nil
			}
			cache.Tail = e.Prev
			e.Prev = nil
			delete(cache.List, e.Key) // delete from map
			cache.Size--              // reset cache size
		}
	}
}

func (cache *LRU) moveTop(e *Element) {

	if cache.Head == nil {
		cache.Head = e
		cache.Tail = e
	} else {
		e.Next = cache.Head
		cache.Head.Prev = e
		cache.Head = e
	}
}

func (cache *LRU) printCache() {
	if cache.List != nil {
		e := cache.Head
		for e != nil {
			fmt.Println(e.Key)
			e = e.Next
		}
	}

}
func main() {
	fmt.Println("creating new cache")
	cache := &LRU{}
	cache.init(3)
	cache.put(4, "hello")
	cache.put(5, "asdf")
	cache.put(6, "asdf")
	cache.put(6, "asdfasdfadsfsd")
	cache.put(8, "asdf")
	cache.put(9, "asdf")
	cache.put(10, "asdf")
	cache.put(10, "asddfdff")
	cache.printCache()
}
