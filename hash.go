package main

type Element struct {
	Value interface{}
}

type HashFunc func(elem Element) uint32

type HashTable interface {
	Insert(elem Element)
	Remove(elem Element)
	Contains(elem Element) bool
}
