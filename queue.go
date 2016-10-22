package trygo

import (
	"container/list"
	"reflect"
)

// Queue struct
type Queue struct {
	sem  chan int
	list *list.List
}

var tFunc func(val interface{}) bool

// NewQueue create a new Queue and return.
func NewQueue() *Queue {
	sem := make(chan int, 1)
	list := list.New()
	return &Queue{sem, list}
}

// Size get size of the queue
func (q *Queue) Size() int {
	return q.list.Len()
}

// Enqueue puts an element into queue.
func (q *Queue) Enqueue(val interface{}) *list.Element {
	q.sem <- 1
	e := q.list.PushFront(val)
	<-q.sem
	return e
}

// Dequeue puts an element out of the queue.
func (q *Queue) Dequeue() *list.Element {
	q.sem <- 1
	e := q.list.Back()
	q.list.Remove(e)
	<-q.sem
	return e
}

// Query returns the element in the queue only if func queueFunc(element) returns true..
func (q *Queue) Query(queryFunc interface{}) *list.Element {
	q.sem <- 1
	e := q.list.Front()
	for e != nil {
		if reflect.TypeOf(queryFunc) == reflect.TypeOf(tFunc) {
			if queryFunc.(func(val interface{}) bool)(e.Value) {
				<-q.sem
				return e
			}
		} else {
			<-q.sem
			return nil
		}
		e = e.Next()
	}
	<-q.sem
	return nil
}

// Contain tests if this item in the queue.
func (q *Queue) Contain(val interface{}) bool {
	q.sem <- 1
	e := q.list.Front()
	for e != nil {
		if e.Value == val {
			<-q.sem
			return true
		}
		e = e.Next()
	}
	<-q.sem
	return false
}
