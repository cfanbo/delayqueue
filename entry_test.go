package delayqueue

import (
	"testing"
	"time"
)

func TestNewEntry(t *testing.T) {
	bornTime := time.Now()
	entry := NewEntry(bornTime, "aaa")
	if bornTime != entry.bornTime {
		t.Fatal("entry bornTime error ")
	}

	if "aaa" != entry.Body() {
		t.Fatal("entry body error")
	}
}
