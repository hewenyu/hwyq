package queue

import (
	"sync"
	"time"
)

type (
	//Queue 队列
	Queue struct {
		top    *node
		rear   *node
		length int
		sync.RWMutex
	}
	//双向链表节点
	node struct {
		pre   *node
		next  *node
		value interface{}
		tims  time.Time
	}
)

// Create a new queue
func New() *Queue {
	return &Queue{
		top:    nil,
		rear:   nil,
		length: 0,
	}
}

//获取队列长度
func (this *Queue) Len() int {

	this.RLock()
	defer this.RUnlock()

	return this.length
}

//返回true队列不为空
func (this *Queue) Any() bool {
	this.RLock()
	defer this.RUnlock()

	return this.length > 0
}

//返回队列顶端元素
func (this *Queue) Peek() interface{} {

	this.RLock()
	defer this.RUnlock()

	if this.top == nil {
		return nil
	}
	return this.top.value
}

/*
Push 入队操作
*/
func (this *Queue) Push(v interface{}) {

	this.Lock()
	defer this.Unlock()

	n := &node{nil, nil, v, time.Now().Local()}
	if this.length == 0 {
		this.top = n
		this.rear = this.top
	} else {
		n.pre = this.rear
		this.rear.next = n
		this.rear = n
	}
	this.length++
}

/*
Pop 出队操作
*/
func (this *Queue) Pop() interface{} {

	this.Lock()
	defer this.Unlock()

	if this.length == 0 {
		return nil
	}
	n := this.top
	if this.top.next == nil {
		this.top = nil
	} else {
		this.top = this.top.next
		this.top.pre.next = nil
		this.top.pre = nil
	}
	this.length--
	return n.value
}
