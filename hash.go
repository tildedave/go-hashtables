package main

type Element struct {
	Value interface{}
}

type HashFunc func(elem Element) int32

type HashTable []Element

func New(size int) *HashTable {
	table := make(HashTable, size)
	return &table
}

func (ht HashTable) Size() int {
	return len(ht)
}
