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

func TestNew(t *testing.T) {
	table := New(1000, one, two)
	if s := table.Size(); s != 1000 {
		t.Errorf("Table was not expected size.")
	}
}
