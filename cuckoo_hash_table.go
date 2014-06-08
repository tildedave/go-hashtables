package main

import (
	"log"
	"math/rand"
	"time"
)

const MAX_LOOPS = 10

type CuckooHashTable struct {
	table1 []Element
	table2 []Element
	hash1  HashFunc
	hash2  HashFunc

	rand *rand.Rand
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

	A := uint(r.Uint32())
	B := uint(r.Uint32())

	return func(elem Element) uint32 {
		value := elem.Value.(int)
		high := uint(value >> 16)
		low := uint(value & 0x0000FFFF)

		hash := uint32(high*A + low*B)
		return hash >> (32 - log2(size))
	}
}

func New(size int, seed int64) *CuckooHashTable {
	var r *rand.Rand

	if seed == 0 {
		seed = time.Now().UnixNano()
	}
	r = rand.New(rand.NewSource(seed))

	hash1, hash2 := generateHashFunction(r, size/2), generateHashFunction(r, size/2)
	table := CuckooHashTable{
		table1: make([]Element, size/2),
		table2: make([]Element, size/2),
		hash1:  hash1,
		hash2:  hash2,
		rand:   r,
	}
	return &table
}

func (ht *CuckooHashTable) Size() int {
	return len(ht.table1) * 2
}

func (ht *CuckooHashTable) Insert(elem Element) {
	if ht.Contains(elem) {
		return
	}

	success := ht.doInsert(elem)
	if !success {
		ht.rehash()
		ht.Insert(elem)
	}
}

func (ht *CuckooHashTable) doInsert(elem Element) bool {
	var displacedElem Element

	loop := 0
	for ; loop < MAX_LOOPS; loop++ {
		k1 := ht.hash1(elem)
		log.Printf("Inserting %v into table (k1 = %d)", elem, k1)
		if ht.table1[k1].Value == nil {
			log.Printf("Setting table1[%d] = %v", k1, elem)
			ht.table1[k1] = elem
			return true
		}

		displacedElem = ht.table1[k1]
		log.Printf("Setting table1[%d] = %v, displacing %v (cuckoo)", k1, elem, displacedElem)
		ht.table1[k1] = elem
		elem = displacedElem
		k2 := ht.hash2(elem)

		log.Printf("Inserting %v into table (k2 = %d)", elem, k2)

		if ht.table2[k2].Value == nil {
			log.Printf("Setting table2[%d] = %v", k2, elem)
			ht.table2[k2] = elem
			return true
		}

		displacedElem = ht.table2[k2]
		log.Printf("Setting table2[%d] = %v, displacing %v (cuckoo)", k2, elem, displacedElem)
		ht.table2[k2] = elem
		elem = displacedElem
	}

	log.Printf("Failed to insert %v into table", elem)
	return false
}

func (ht *CuckooHashTable) rehash() {
	size := ht.Size()
	items := make([]Element, 0, size)
	for _, item := range append(ht.table1, ht.table2...) {
		if item.Value != nil {
			items = append(items, item)
		}
	}

redo:
	for {
		ht.table1 = make([]Element, size/2)
		ht.table2 = make([]Element, size/2)

		ht.hash1 = generateHashFunction(ht.rand, size/2)
		ht.hash2 = generateHashFunction(ht.rand, size/2)

		for _, item := range items {
			success := ht.doInsert(item)
			if success == false {
				continue redo
			}
		}

		return
	}
}

func (ht *CuckooHashTable) Contains(elem Element) bool {
	k1, k2 := ht.hash1(elem), ht.hash2(elem)
	log.Printf("Checking for %v in table (k1 = %d, k2 = %d)", elem, k1, k2)

	if ht.table1[k1] == elem {
		log.Printf("Found %v at table1[%d]", elem, k1)
		return true
	} else if ht.table2[k2] == elem {
		log.Printf("Found %v at table2[%d]", elem, k2)
		return true
	}

	log.Printf("Did not find %v in table", elem)
	return false
}

func (ht *CuckooHashTable) Remove(elem Element) {
	k1, k2 := ht.hash1(elem), ht.hash2(elem)
	log.Printf("Removing %v from table (k1 = %d, k2 = %d)", elem, k1, k2)

	if ht.table1[k1] == elem {
		log.Printf("Removing %v from table1[%d]", elem, k1)
		ht.table1[k1] = Element{nil}
	} else if ht.table2[k2] == elem {
		log.Printf("Removing %v from table2[%d]", elem, k2)
		ht.table2[k2] = Element{nil}
	}
}
