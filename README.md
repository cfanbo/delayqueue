# 延时队列

 [![Build Status](https://travis-ci.com/cfanbo/delayqueue.svg?branch=master)](https://travis-ci.com/cfanbo/delayqueue)
 [![GoDoc](https://godoc.org/github.com/cfanbo/delayqueue.svg?status.svg)](https://godoc.org/github.com/cfanbo/delayqueue)
 


基于Golang实现的延时队列

## 功能
延时队列是指在指定的时间点进行消息消费，具体消费逻辑由用户自己来实现。

传统解决方法一般采用cron来实现，但有以下缺点：
- 轮训效率太低，每次都需要扫库。
- 如果扫库频率太高，则后端数据库压力过大，如果频率太低，则存在有效性时间差较大的问题


## 使用场景
* 用户网上购买时后，如果收货后，15天内未对交易进行评论，则系统进行默认5星评论。  
* 订单超过30分钟未支付，则系统进行自动取消  
  

## 实现原理
主要使用到的两个数据结构是 环形队列 和 集合。其中环形队列是由数组来实现。

### 基本概念
> currentSlot 表示当前操作的环位置，这里是数组的索引值  
> timer 定时器，默认每秒移动一个slot

系统主要由三部分组成，分别为slot、Elements和Element。  
> Slots 代表一个环, 由多个slot组成，每个slot对应一个Elements     
> Elements slot对应的值  
> Element 组成Elements集合的元素   

每个环节点slot就是一个数据集合 Elements，这个集合内的数据则表示当前时间点需要进行消费的信息集合，有可能是下次循环到这个节点的时间进行消费。  

**环与集合的关系**  
slots[0] = Elements  
slots[1] = Elements  
slots[...] = ...  

一个Elements是由一个或多个 Element 元素组成，每个 Element 元素都有一个 cycleNum 字段，用来表示此元素是立即消费还是以后消费，其值也可以理解成环的循环周期。
如果cycleNum字段值为0，则表示立即消费，如果cycleNum=2则表示还需要两个环周期才能消费，每次循环都进行 cycleNum-- 操作，直到为0时结束。

**集合与元素的关系**  
Elements = {Element、Element、Element}

所以整个延时队列看起来是这个样子:  

    slots[0] = []*Elements{*Element{}, *Element{}, *Element{}...}
    slots[1] = []*Elements{*Element{}, *Element{}, *Element{}...}
    slots[2] = []*Elements{*Element{}, *Element{}, *Element{}...}
    ...

### 实现原理
系统会有一个定时器timer，每1秒(可通过delayqueue.WithFrequency 函数调整)会移动一个slot, 此时currentSlot的值加1，表示下一个节点位置。   
然后遍历当前环点中的所有元素，如果当前元素生命周期cycleNum=0，则立即消费，否则将cycleNum--, 直到循环完集合中的所有元素。  

同时每次添加新元素时，都要以当前时间所在的slot位置为起点，假如当前时间为 00:05:10, 在第 310 (5*60+10) 个slot, 这时添加一个元素时间为 00:02:50,  
由于每秒移动一个slot, 而新添加元素时间slot为179(2*60+50), 则将这个元素放在当前位置后往数的第179个slot, 即这个环的第 310+179=489个slot中。
如果添加的时间大于当前时间的多个环周期时，只需要将环周期对应的slot个数减去即可，环的周期数使用 cycleNum 值来表示。   

## 演示代码

    package main
    
    import (
    	"fmt"
    	"time"
    
    	"github.com/cfanbo/delayqueue"
    )
    
    func consume(entry delayqueue.Entry) {
    	fmt.Println("当前：", time.Now().Format("2006-01-02 15:04:05"))
    	fmt.Println("消费：", entry.ConsumeTime().Format("2006-01-02 15:04:05"))
    	fmt.Println("消费内容", entry.Body())
    	fmt.Println("=======================")
    }
    
    func main() {
    	q := delayqueue.New()
    	q.Put(time.Now().Add(time.Second*2), "2秒后")
    	q.Put(time.Now().Add(time.Second*15), "15秒后")
    	q.Put(time.Now().Add(time.Second*8), "8秒后")
    	q.Put(time.Now().Add(time.Second*43), "43秒后")
    	q.Put(time.Now().Add(time.Second*50), "50秒后")
    	q.Put(time.Now().Add(time.Second*28), "28秒后")
    
        ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
        defer cancel()
        q.Run(ctx, consume)
    }


支持用户自定义间隔时间，如每分钟，每小时，只要是time.NewTicker()支持的 time.Duration 类型即可。  
调用方法如下：  
  
    // 在New() 函数里调用 WithFrequency() 函数即可
    q := delayqueue.New(delayqueue.WithFrequency(time.Minute))
    q.Put(time.Now().Add(time.Minute * 2), "2分钟后消费此内容")

## 说明
1. 系统支持频率周期类型是time.Duration, 如 time.Second、time.Minute 和 time.Hour。
2. 队列暂不支持数据持久化，所以若停止服务或者退出重启，则队列数据将全部丢失。  
 建议在数据消费后对其状态进行变更存储，以便在下次服务启动成功后，立即将需要处理的数据写入延时队列。
  
