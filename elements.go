// slot单元，每个slot中会存储多个element, slot数据为slice

package delayqueue

import (
	"sync"
)

type Elements struct {
	mu sync.Mutex
	elements []*Element
}

// 初始化slot，初始化以前slot的值为nil
func NewElements() *Elements {
	return &Elements{}
}

// Append 添加新元素到slot
func (e *Elements)Append(ele *Element) {
	e.mu.Lock()
	defer e.mu.Unlock()

	if e.elements == nil {
		e.elements = make([]*Element, 0)
	}
	e.elements = append(e.elements, ele)
}

// Detection 遍历检测每个solt中的元素
// 如果元素的当前生命周期为0，则表示立即执行，返回将生命周期进行 cycleNum--
func (e *Elements)Detection(ch chan<- Entry) {
	e.mu.Lock()
	defer e.mu.Unlock()

	for i, ele := range e.elements {

		if ele.cycleNum == 0 {
			// 写入chan
			entry := NewEntry(ele.bornTime, ele.data)
			ch<-entry

			// 执行回调并移除此队列
			// 第一个元素
			if i == 0 {
				e.elements = e.elements[1:]
				continue
			}

			// 最后一个元素
			len := len(e.elements) - 1
			if i == len {
				e.elements = e.elements[:len-1]
				continue
			}

			// 中间元素
			e.elements = append(e.elements[:i], e.elements[i+1:]...)
		} else {
			// 减少生命周期-1
			ele.subCycleNum()
		}
	}
}

// Empty 判断slot是否空slot
func (e *Elements)Empty() bool {
	return e.Len() == 0
}

// Len 返回slot里面元素的个数
func (e *Elements)Len() int {
	return len(e.elements)
}