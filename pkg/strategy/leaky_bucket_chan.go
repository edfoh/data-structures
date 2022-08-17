package strategy

import (
	"errors"
	"sync"
	"time"
)

type LeakyBucketChan[T any] struct {
	sync.RWMutex
	in           chan T
	out          chan<- T
	interval     time.Duration
	pollInterval time.Duration
	running      bool
}

func NewLeakyBucketChan[T any](out chan<- T, cap int, ratePerSecond int) *LeakyBucketChan[T] {
	return &LeakyBucketChan[T]{
		in:           make(chan T, cap),
		out:          out,
		interval:     time.Second / time.Duration(ratePerSecond),
		pollInterval: 1 * time.Second,
	}
}

func (b *LeakyBucketChan[T]) Enqueue(task T) error {
	select {
	case b.in <- task:
		return nil
	default:
		return errors.New("queue is full")
	}
}

func (b *LeakyBucketChan[T]) Start() {
	b.Lock()
	defer b.Unlock()

	b.running = true

	go func() {
		for b.isRunning() {
			select {
			case item := <-b.in:
				go func() { b.out <- item }()
				time.Sleep(b.interval)
			default:
				time.Sleep(b.pollInterval)
			}
		}
	}()
}

func (b *LeakyBucketChan[T]) Stop() {
	b.Lock()
	defer b.Unlock()

	b.running = false
}

func (b *LeakyBucketChan[T]) isRunning() bool {
	b.RLock()
	defer b.RUnlock()

	return b.running
}
