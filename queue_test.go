package delayqueue

import (
	"fmt"
	"testing"
	"time"
)

func consume(entry Entry) {
	fmt.Println("当前：", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println("消费：", entry.ConsumeTime().Format("2006-01-02 15:04:05"))
	fmt.Println("消费内容：", entry.Body())
	fmt.Println("=======================")
}

func TestQueue(t *testing.T) {
	q := New()

	q.Put(time.Now().Add(time.Second * 2), "2秒后")
	q.Put(time.Now().Add(time.Second * 15), "15秒后")
	q.Put(time.Now().Add(time.Second * 8), "8秒后")
	q.Put(time.Now().Add(time.Second * 43), "43秒后")

	q.Put(time.Now().Add(time.Second * 50), "50秒后")
	q.Put(time.Now().Add(time.Second * 28), "28秒后")
	//q.Debug(true)

	q.Run(consume)
}
