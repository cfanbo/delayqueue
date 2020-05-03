// 消息实体

package delayqueue

import (
	"time"
)

// 消费实体
type Entry struct {
	// 信息创建时间
	bornTime time.Time

	// 消费时间
	consumeTime time.Time

	// 消息内容
	body interface{}
}

// NewEntry 创建消费实体
func NewEntry(bt time.Time, body interface{}) Entry {
	return Entry{
		consumeTime: time.Now(),
		bornTime: bt,
		body: body,
	}
}

// BornTime 消息创建时间
func (e Entry)BornTime() time.Time {
	return e.bornTime
}

// ConsumeTime 消费时间
func (e Entry)ConsumeTime() time.Time {
	return e.consumeTime
}

// Body 消息内容
func (e Entry)Body() interface{} {
	return e.body
}

