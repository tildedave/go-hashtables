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
	table.hash1 = one
	table.hash2 = two

	if s := table.Size(); s != 1000 {
		t.Errorf("Table was not expected size.")
	}
}

func TestInsert(t *testing.T) {
	table := New(22, 0)
	table.hash1 = one
	table.hash2 = two

	table.Insert(Element{20})

	if table.table1[9].Value != 20 {
		t.Errorf("Table did not insert into expected row.")
	}
	if table.table2[1].Value != nil {
		t.Errorf("Second table should not have been modified.")
	}

	table.Insert(Element{31})

	if table.table1[9].Value != 31 {
		t.Errorf("Table should have overwritten original value.")
	}
	if table.table2[1].Value != 20 {
		t.Errorf("First value should have been inserted into second table.")
	}
}

func TestInsertAndContains(t *testing.T) {
	table := New(22, 0)
	table.hash1 = one
	table.hash2 = two

	table.Insert(Element{20})
	if table.Contains(Element{20}) != true {
		t.Errorf("Table did not contain element.")
	}

	table.Insert(Element{31})
	if table.Contains(Element{31}) != true {
		t.Errorf("Table did not contain element.")
	}

	if table.Contains(Element{73}) != false {
		t.Errorf("Table should not have contained non-inserted element.")
	}
}

func TestRemove(t *testing.T) {
	table := New(22, 0)
	table.hash1 = one
	table.hash2 = two

	table.Insert(Element{20})
	table.Insert(Element{31})
	table.Remove(Element{31})
	table.Remove(Element{20})

	if table.Contains(Element{31}) {
		t.Errorf("Table should have removed element.")
	}

	if table.Contains(Element{20}) {
		t.Errorf("Table should have removed element.")
	}
}
