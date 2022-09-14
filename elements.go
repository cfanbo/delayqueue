// slot单元，每个slot中会存储多个element, slot数据为slice

package delayqueue

import (
	"sync"
)

// Elements 队列slot 为Element元素的集合
type Elements struct {
	mu       sync.Mutex
	elements []*Element
}

// NewElements 初始化slot，初始化以前 slot 的值为nil
func NewElements() *Elements {
	return &Elements{}
}

// Append 添加新元素 Element 到slot
func (e *Elements) Append(ele *Element) {
	e.mu.Lock()
	defer e.mu.Unlock()

	if e.elements == nil {
		e.elements = make([]*Element, 0)
	}
	e.elements = append(e.elements, ele)
}

// Detection 遍历检测每个 solt 中的每个元素
// 如果元素的当前生命周期为0，则表示立即执行，返回将生命周期进行 cycleNum--
func (e *Elements) Detection(ch chan<- Entry) {
	e.mu.Lock()
	defer e.mu.Unlock()

	k := 0
	for _, ele := range e.elements {
		if ele.cycleNum == 0 {
			// 写入chan
			entry := NewEntry(ele.bornTime, ele.data)
			ch <- entry

		} else {
			// 减少生命周期-1
			ele.subCycleNum()

			// 切片左侧是所有有效数据
			e.elements[k] = ele
			k++
		}
	}
	e.elements = e.elements[:k]
}

// Empty 判断slot是否空slot
func (e *Elements) Empty() bool {
	return e.Len() == 0
}

// Len 返回slot里面元素的个数
func (e *Elements) Len() int {
	e.mu.Lock()
	defer e.mu.Unlock()
	return len(e.elements)
}
