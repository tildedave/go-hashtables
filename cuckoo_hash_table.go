package main

import (
	"math/rand"
	"time"
)

type CuckooHashTable struct {
	table []Element

	rand  *rand.Rand
	func1 HashFunc
	func2 HashFunc
}

func log2(size int) uint {
	r := 0
	for size != 0 {
		size >>= 1
		r++
	}

	// Number of iterations to reach 0 value minus 1, e.g. log2(1) = 0
	return uint(r - 1)
}

func generateHashFunction(r *rand.Rand, size int) HashFunc {
	// From http://www.keithschwarz.com/interesting/code/cuckoo-hashmap/CuckooHashMap.java.html
	// Universal hash family
	//    HashCode = ((HIGH + A) * (LOW * B)) / (2^(32 - k))
	// For A, B random

	return func(elem Element) int {
		value := elem.Value.(int32)
		high := value >> 16
		low := value & 0x0000FFFF
		A := r.Int31()
		B := r.Int31()

		return int((high + A) + (low+B)>>(2^(32-log2(size))))
	}
}

func New(size int, seed int64) *CuckooHashTable {
	var r *rand.Rand

	if seed == 0 {
		seed = time.Now().UnixNano()
	}
	r = rand.New(rand.NewSource(seed))

	func1, func2 := generateHashFunction(r, size), generateHashFunction(r, size)
	table := CuckooHashTable{
		table: make([]Element, size),
		func1: func1,
		func2: func2,
		rand:  r,
	}
	return &table
}

func (ht *CuckooHashTable) Size() int {
	return len(ht.table)
}
