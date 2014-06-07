package main

import (
	"testing"
)

func TestNew(t *testing.T) {
	table := New(1000)
	if s := table.Size(); s != 1000 {
		t.Errorf("Table was not expected size.")
	}
}
