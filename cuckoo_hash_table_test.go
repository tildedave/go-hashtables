package main

import (
	"testing"
)

func one(elem Element) int {
	k := elem.Value.(int)
	return k % 11
}

func two(elem Element) int {
	k := elem.Value.(int)
	return (k / 11) % 11
}

func TestLog2(t *testing.T) {
	if log2(8) != 3 {
		t.Errorf("Wrong log2 of 8, %d", log2(8))
	}

	if log2(16) != 4 {
		t.Errorf("Wrong log2 of 16, %d", log2(16))
	}
}

func TestNew(t *testing.T) {
	table := New(1000, 0)
	table.func1 = one
	table.func2 = two

	if s := table.Size(); s != 1000 {
		t.Errorf("Table was not expected size.")
	}
}
