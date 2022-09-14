// Copyright 2013 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

// delayqueue 是一款基于Golang开发的高性能延时队列
//
// 使用示例:
//
//
//	package main
//
//	import (
//		"fmt"
//		"time"
//
//		"github.com/cfanbo/delayqueue"
//	)
//
//	func consume(entry delayqueue.Entry) {
//		fmt.Println("当前：", time.Now().Format("2006-01-02 15:04:05"))
//		fmt.Println("消费：", entry.ConsumeTime().Format("2006-01-02 15:04:05"))
//		fmt.Println("消费内容", entry.Body())
//		fmt.Println("=======================")
//	}
//
//	func main() {
// 		// 分钟级别 delayqueue.New(delayqueue.WithFrequency(time.Minute))
//		q := delayqueue.New()
//		q.Put(time.Now().Add(time.Second*2), "2秒后")
//		q.Put(time.Now().Add(time.Second*15), "15秒后")
//		q.Put(time.Now().Add(time.Second*8), "8秒后")
//		q.Put(time.Now().Add(time.Second*43), "43秒后")
//		q.Put(time.Now().Add(time.Second*50), "50秒后")
//		q.Put(time.Now().Add(time.Second*28), "28秒后")
//
//		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
//		defer cancel()
//		q.Run(ctx, consume)
//	}
//
package delayqueue
