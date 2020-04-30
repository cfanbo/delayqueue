package queue

import (
	"testing"
)

func TestElement(t *testing.T) {
	e := NewElement(4, "this is a test")
	e.subCycleNum()
	if e.cycleNum != 3 {
		t.Fatal("element test failed")
	}
}

func BenchmarkNewElement(b *testing.B) {
	for i:=0; i< b.N;i++ {
		e := NewElement(4, "this is a test")
		e.subCycleNum()
	}
}


