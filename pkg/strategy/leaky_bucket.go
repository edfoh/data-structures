package strategy

import (
	"time"
)

var now = time.Now

type LeakyBucket struct {
	capacity      int
	current       int
	leakPerSecond float64
	lastUpdated   time.Time
}

func NewLeakyBucket(tokensPerInterval int, interval time.Duration, capacity int) *LeakyBucket {
	ratePerSecond := interval.Seconds() * float64(tokensPerInterval)
	return &LeakyBucket{
		current:       0,
		leakPerSecond: ratePerSecond,
		capacity:      capacity,
		lastUpdated:   now(),
	}
}

func (b *LeakyBucket) Count() int {
	timeElapsed := now().Sub(b.lastUpdated)
	amtLeaked := b.leakPerSecond * timeElapsed.Seconds()
	return max(0, b.current-int(amtLeaked))
}

func (b *LeakyBucket) AddN(n int) (bool, int) {
	success := true
	spillover := 0
	count := b.Count()

	if count+n > b.capacity {
		success = false
		spillover = abs(b.capacity - count - n)
	}

	b.current = min(b.capacity, b.current+n)
	b.lastUpdated = now()
	return success, spillover
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
