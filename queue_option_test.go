package delayqueue

import (
	"testing"
	"time"
)

func TestNewQueueOptions(t *testing.T) {
	opts := NewQueueOptions(WithFrequency(time.Minute))
	m := opts.Frequency()
	if (m != time.Second * 60) {
		t.Fatal("设置队列 Frequency 选项失败")
	}
}
