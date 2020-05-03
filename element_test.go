package delayqueue

import (
	"testing"
	"time"
)

func TestElement(t *testing.T) {
	e := NewElement(time.Now(),4,  "this is a test")
	e.subCycleNum()
	if e.cycleNum != 3 {
		t.Fatal("element test failed")
	}
}

func BenchmarkNewElement(b *testing.B) {
	for i:=0; i< b.N;i++ {
		e := NewElement(time.Now(), 4, "this is a test")
		e.subCycleNum()
	}
}


