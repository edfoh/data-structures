package strategy

import (
	"context"
	"errors"
	"sync"
	"time"
)

var now = time.Now

type window struct {
	startTime time.Time
	count     int
}

func newWindowBefore(interval time.Duration) *window {
	prevTime := now().Add(-interval)
	return &window{
		startTime: prevTime,
		count:     0,
	}
}

func newNowWindow() *window {
	return &window{
		startTime: now(),
		count:     0,
	}
}

func (w *window) AddN(n int) {
	w.count += n
}

func (w *window) Count() int {
	return w.count
}

func (w *window) Reset() {
	w.startTime = now()
	w.count = 0
}

func (w *window) CopyFrom(from *window) {
	w.startTime = from.startTime
	w.count = from.count
}

type SlidingWindow struct {
	sync.Mutex
	prev       *window
	curr       *window
	stopped    bool
	interval   time.Duration
	capacity   int
	ctx        context.Context
	cancelFunc context.CancelFunc
}

func NewSlidingWindow(interval time.Duration, capacity int) *SlidingWindow {

	ctx, cancelFunc := context.WithCancel(context.Background())

	sl := &SlidingWindow{
		prev:       newWindowBefore(interval),
		curr:       newNowWindow(),
		interval:   interval,
		capacity:   capacity,
		ctx:        ctx,
		cancelFunc: cancelFunc,
	}
	go sl.processProgressive()
	return sl
}

func (w *SlidingWindow) GetCount() (int, int) {
	return w.prev.Count(), w.curr.Count()
}

func (w *SlidingWindow) AddN(n int) (bool, error) {
	if w.stopped {
		return false, errors.New("sliding window has stopped")
	}
	w.Lock()
	defer w.Unlock()

	tNow := now()
	currTimeSlide := tNow.Sub(w.curr.startTime)
	prevTimeSlide := w.interval - currTimeSlide

	prevSlideCount := float64(w.interval-prevTimeSlide) / float64(w.interval) * float64(w.prev.Count())
	currSlideCount := float64(w.interval-currTimeSlide) / float64(w.interval) * float64(w.curr.Count())

	windowCount := prevSlideCount + currSlideCount

	if int(windowCount)+n > w.capacity {
		return false, nil
	}

	w.curr.AddN(n)
	return true, nil
}

func (w *SlidingWindow) Stop() {
	if w.stopped {
		return
	}

	w.Lock()
	defer w.Unlock()

	defer w.cancelFunc()
	w.stopped = true
}

func (w *SlidingWindow) processProgressive() {

	for {
		select {
		case <-w.ctx.Done():
			return
		default:
			waitDuration := w.interval - now().Sub(w.curr.startTime)
			time.Sleep(waitDuration)
			w.Lock()
			// copy curr to prev and start new curr
			w.prev.CopyFrom(w.curr)
			w.curr.Reset()
			w.Unlock()
		}

	}
}
