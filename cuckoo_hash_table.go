package main

type CuckooHashTable struct {
	table []Element
	func1 HashFunc
	func2 HashFunc
}

func New(size int, func1 HashFunc, func2 HashFunc) *CuckooHashTable {
	table := CuckooHashTable{
		table: make([]Element, size),
		func1: func1,
		func2: func2,
	}
	return &table
}

func (ht *CuckooHashTable) Size() int {
	return len(ht.table)
}
