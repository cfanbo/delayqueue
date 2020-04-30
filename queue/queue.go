// delay queue
// 队列共有 SLOT_NUM 个slot, 队列存储是一个数组类型变量slots
// 每个slot又是一个slice数据类型的 Elements，存储的是 Element

package queue

import (
	"fmt"
	"sync"
	"time"
)

const (
	// 队列总的slot数，每1秒移动一个slot
	SlotsCount = 3600



)

var once sync.Once

// computeCycleNum 计算生存周期
func computeDealTimeCycleNum(t time.Time) (slotNum, cycleNum int) {
	// 计算当前相差秒数
	now := time.Now().Unix()
	time := t.Unix()
	if now - time >= 0 {
		return 0, 0
	}

	// 相差时间
	diff := time - now
	cycleNum = int(diff) / SlotsCount
	slotNum = int(diff) % SlotsCount

	return
}

// consumeFunc 默认消费函数
func consumeFunc(entry Entry) {
	fmt.Printf("%#v\n", entry)
}

type Queue struct {
	// 数组 共3600个slot, 每秒移动一个slot
	slots [SlotsCount]*Elements

	// 当前正在执行的slot
	currentSlot int

	// mutex
	mu sync.Mutex

	// 定时器
	ticker *time.Ticker

	// 接收chan
	ch chan Entry

	// 消费函数
	consumeFunc func(entry Entry)

	// 调试
	debug bool
}

var singleton *Queue

// NewQueue 创建一个队列
func NewQueue() *Queue {
	opt := NewQueueOptions()
	once.Do(func() {
		singleton = &Queue{
			ticker: time.NewTicker(opt.frequency),
			slots: [SlotsCount]*Elements{},
			ch: make(chan Entry, 100),
			consumeFunc: consumeFunc,
		}
	})

	return singleton
}

// Debug 调度模式
func (q *Queue)Debug(b bool) {
	q.debug = b
}

// Put 写入元素
func (q *Queue) Put(t time.Time, data interface{}) {
	// 计算存储元素所在的slot位置和生命周期
	slotNum, cycleNum := computeDealTimeCycleNum(t)
	ele := NewElement(t, cycleNum, data)

	// 放入指定的slot中
	// 由于是从当前时间开始计算，所以要从当前slot开始计算，往后数第 slotNum 个slot
	// 当前slot位置 + 计算下次运行时间的slot
	if (q.slots[q.currentSlot + slotNum] == nil) {
		q.slots[q.currentSlot + slotNum] = NewElements()
	}
	q.slots[q.currentSlot + slotNum].Append(ele)
}

// Run 启动服务
func (q *Queue)Run(f func(entry Entry)) {
	// define consumeFunc
	if f != nil {
		q.consumeFunc = f
	}

	// detection slot 每次移动一个slot
	go func() {
		for {
			select {
			case <-q.ticker.C:
				// debug
				if (q.debug) {
					go q.info()
				}

				// 检测slot
				go q.consumeSlot(q.currentSlot)

				// 下移一位slot
				if q.currentSlot >= (10 - 1) {
					q.currentSlot = 0
				} else {
					q.currentSlot++
				}
			}
		}
	}()

	// consume
	for {
		select {
			case ele := <-q.ch:
				q.consumeFunc(ele)
		}
	}
}

// sonsumeSlot 检测指定的slot
func (q *Queue)consumeSlot(slotIndex int) {
	if q.slots[slotIndex] == nil {
		// 当前slot从未使用
		return
	} else if q.slots[slotIndex].Empty() {
		// slot 为空
		return
	}

	// 遍历slot中的所有元素(切片类型)
	q.slots[slotIndex].Detection(q.ch)
}

// info 打印debug信息
func (q *Queue)info() {
	q.mu.Lock()
	defer q.mu.Unlock()

	fmt.Println("====", time.Now().Format("2006-01-02 15:04:05"), "====")
	for k, eles := range q.slots {
		var count int
		if eles == nil {
			count = -1
		} else {
			count = eles.Len()
		}
		fmt.Printf("%d: 元素数量 %d\n", k, count)
	}
	fmt.Println()
}