package strategy

import (
	"context"
	"errors"
	"math"
	"sync"
	"time"
)

type window struct {
	startTime time.Time
	count     int
}

func newWindow(t time.Time) *window {
	return &window{
		startTime: t,
		count:     0,
	}
}

func (w *window) AddN(n int) {
	w.count += n
}

func (w *window) Count() int {
	return w.count
}

func (w *window) StartTime() time.Time {
	return w.startTime
}

func (w *window) Set(t time.Time, c int) {
	w.SetTime(t)
	w.SetCount(c)
}

func (w *window) SetTime(t time.Time) {
	w.startTime = t
}

func (w *window) SetCount(c int) {
	w.count = c
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
	currTime := now()
	prevTime := currTime.Add(-interval)
	sl := &SlidingWindow{
		prev:       newWindow(prevTime),
		curr:       newWindow(currTime),
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
	currTimeSlide := tNow.Sub(w.curr.StartTime())
	prevTimeSlide := w.interval - currTimeSlide

	prevSlideCount := float64(prevTimeSlide) / float64(w.interval) * float64(w.prev.Count())
	currSlideCount := float64(currTimeSlide) / float64(w.interval) * float64(w.curr.Count())

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
			waitDuration := w.interval - now().Sub(w.curr.StartTime())
			time.Sleep(waitDuration)
			w.Lock()
			// copy curr to prev and start new curr
			w.prev.CopyFrom(w.curr)
			w.curr.Reset()
			w.Unlock()
		}

	}
}

type SyncSlidingWindow struct {
	prev     *window
	curr     *window
	interval time.Duration
	capacity int
}

func NewSyncSlidingWindow(interval time.Duration, capacity int) *SyncSlidingWindow {
	currTime := now()
	prevTime := currTime.Add(-interval)
	return &SyncSlidingWindow{
		prev:     newWindow(prevTime),
		curr:     newWindow(currTime),
		capacity: capacity,
		interval: interval,
	}
}

func (w *SyncSlidingWindow) getNSlides(t time.Time) (bool, int) {
	diff := t.Sub(w.curr.startTime)
	isInCurrentWindow := diff < w.interval
	nSlides := float64(diff) / float64(w.interval)

	return isInCurrentWindow, int(nSlides)
}

func (w *SyncSlidingWindow) Count() (int, int) {
	return w.prev.Count(), w.curr.Count()
}

func (w *SyncSlidingWindow) AddN(n int) (bool, error) {
	tNow := now()

	w.adjustWindows(tNow)

	currTimePortion := tNow.Sub(w.curr.StartTime())
	prevTimePortion := w.interval - currTimePortion

	currTimeCount := float64(currTimePortion) / float64(w.interval) * float64(w.curr.Count())
	prevTimeCount := float64(prevTimePortion) / float64(w.interval) * float64(w.prev.Count())
	totalCount := currTimeCount + prevTimeCount

	if int(math.Round(totalCount))+n > w.capacity {
		return false, errors.New("sliding window is full")
	}

	w.curr.AddN(n)
	return true, nil
}

func (w *SyncSlidingWindow) adjustWindows(t time.Time) {
	if isInCurrentWindow, nSlides := w.getNSlides(t); !isInCurrentWindow {

		// if we slide by 1 window, init new curr and copy curr to prev
		if nSlides == 1 {
			w.prev.CopyFrom(w.curr)
			w.curr = newWindow(t)
		} else if nSlides > 1 { // if greater than 1, set prev to nothing and init new curr
			prevTime := t.Add(-w.interval)
			w.prev = newWindow(prevTime)
			w.curr = newWindow(t)
		}
	}
}
