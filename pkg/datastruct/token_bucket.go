package datastruct

import (
	"errors"
	"math"
	"sync"
	"time"
)

type TokenBucket struct {
	sync.Mutex
	maxTokens           float64
	currentTokens       float64
	refillRatePerSecond float64
	lastTakenTime       time.Time
}

func NewTokenBucket(maxTokens float64, refillRatePerSecond float64) *TokenBucket {
	return &TokenBucket{
		maxTokens:           maxTokens,
		currentTokens:       maxTokens,
		refillRatePerSecond: refillRatePerSecond,
		lastTakenTime:       time.Now(),
	}
}

func (b *TokenBucket) TakeN(n int) error {
	b.Lock()
	defer b.Unlock()

	b.refill()
	if b.currentTokens-float64(n) < 0 {
		return errors.New("ran out of tokens")
	}

	b.currentTokens -= float64(n)
	return nil
}

func (b *TokenBucket) refill() {
	timeElapsed := time.Since(b.lastTakenTime)
	refillNumTokens := math.Floor(timeElapsed.Seconds() * b.refillRatePerSecond)

	b.currentTokens = math.Min(b.currentTokens+refillNumTokens, b.maxTokens)
}
