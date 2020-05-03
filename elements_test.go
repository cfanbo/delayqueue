package delayqueue

import (
	"testing"
	"time"
)

func TestElements(t *testing.T)  {
	eles := NewElements()
	if eles.elements != nil {
		t.Fatal("初始化 slot 失败")
	}

	ele := NewElement(time.Now(), 4, "test")
	eles.Append(ele)
	if eles.elements == nil || eles.Empty() {
		t.Fatal("slot 添加 Element 元素失败")
	}

	if eles.elements[0].cycleNum != 4 || eles.elements[0].data != "test" {
		t.Fatal("添加的元素内容与存储的内容不一致")
	}
}

func BenchmarkNewElements(b *testing.B) {
	eles := NewElements()
	for i :=0; i<b.N; i++ {
		ele := NewElement(time.Now(), i, i)
		eles.Append(ele)
	}
}