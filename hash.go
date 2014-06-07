package main

type Element struct {
	Value interface{}
}

type HashFunc func(elem Element) int

type HashTable interface {
	Insert(elem Element)
	Delete(elem Element)
	Contains(elem Element) bool
}
