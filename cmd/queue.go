package main

import "sync"

type Queue[T any] struct {
	slice []T
	lock  sync.Mutex

	Size int
}

func (q *Queue[T]) Push(val T) {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.slice = append(q.slice, val)
	q.Size++
}

func (q *Queue[T]) Pop() (*T, bool) {
	q.lock.Lock()
	defer q.lock.Unlock()

	if len(q.slice) == 0 {
		return nil, false
	}

	val := q.slice[0]
	q.slice = q.slice[1:]

	q.Size--

	return &val, true
}
