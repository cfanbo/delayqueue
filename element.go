// slot中存储的元素

package delayqueue

import "time"

// 队列元素结构体
type Element struct {
	// 创建时间
	bornTime time.Time

	// 消费时间
	consumeTime time.Time

	// 生存周期
	cycleNum int

	// 存储数据
	data interface{}
}

// 创建新的队列元素
func NewElement(t time.Time, cycleNum int, data interface{}) *Element {
	return &Element{
		bornTime: time.Now(),
		consumeTime: t,
		cycleNum: cycleNum,
		data: data,
	}
}

// subCycleNum 生命周期减1
func (e *Element) subCycleNum() {
	if e.cycleNum < 1 {
		return
	}
	e.cycleNum--
}
