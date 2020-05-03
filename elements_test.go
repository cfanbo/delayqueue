package delayqueue

import (
	"fmt"
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

func TestElements_Detection(t *testing.T) {
	eles := NewElements()

	for i:=0; i< 10; i++ {
		e := &Element{
			bornTime:    time.Now(),
			consumeTime: time.Now(),
			cycleNum:    int(i / 3),
			data:        "aa",
		}
		eles.Append(e)
	}
	for k, v := range eles.elements {
		fmt.Println(k, v)
	}

	k := 0
	for _, e := range eles.elements {
		if e.cycleNum == 0 {
			fmt.Println("delete", e.cycleNum)
		} else {
			// 存储
			eles.elements[k] = e
			k++
		}
	}
	eles.elements = eles.elements[:k]
	for k, v := range eles.elements {
		fmt.Println(k, v)
	}

	if (len(eles.elements) != 7) {
		t.Fatal("测试for语句中删除切片元素失败")
	}
}


func BenchmarkNewElements(b *testing.B) {
	eles := NewElements()
	for i :=0; i<b.N; i++ {
		ele := NewElement(time.Now(), i, i)
		eles.Append(ele)
	}
}