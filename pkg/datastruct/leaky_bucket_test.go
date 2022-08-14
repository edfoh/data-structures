package datastruct

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLeakyBucket(t *testing.T) {
	testCases := []struct {
		desc                      string
		capacity                  int
		ratePerSecond             int
		enqueueItems              []int
		wantDequeued              []int
		wantErrs                  []error
		wantDequeueTimeElapsedMin time.Duration
		wantDequeueTimeElapsedMax time.Duration
	}{
		{
			desc:                      "can process within capacity",
			capacity:                  10,
			ratePerSecond:             5,
			enqueueItems:              []int{1, 2, 3, 4, 5},
			wantDequeued:              []int{1, 2, 3, 4, 5},
			wantErrs:                  []error{nil, nil, nil, nil, nil},
			wantDequeueTimeElapsedMin: 0,
			wantDequeueTimeElapsedMax: 1 * time.Second,
		},
		{
			desc:                      "cannot enqueue if queue is full",
			capacity:                  1,
			ratePerSecond:             5,
			enqueueItems:              []int{1, 2},
			wantDequeued:              []int{1},
			wantErrs:                  []error{nil, errors.New("queue is full")},
			wantDequeueTimeElapsedMin: 0,
			wantDequeueTimeElapsedMax: 1 * time.Second,
		},
		{
			desc:                      "rate of 1 item per second should take approx 5 seconds for 5 items",
			capacity:                  10,
			ratePerSecond:             1,
			enqueueItems:              []int{1, 2, 3, 4, 5},
			wantDequeued:              []int{1, 2, 3, 4, 5},
			wantErrs:                  []error{nil, nil, nil, nil, nil},
			wantDequeueTimeElapsedMin: 4 * time.Second,
			wantDequeueTimeElapsedMax: 5 * time.Second,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			out := make(chan int)
			bucket := NewLeakyBucket(out, tC.capacity, tC.ratePerSecond)

			bucket.Start()

			for i, enqueuedItem := range tC.enqueueItems {
				gotErr := bucket.Enqueue(enqueuedItem)
				require.Equal(t, tC.wantErrs[i], gotErr)
			}

			var gotDequeued []int
			go func() {
				for {
					if len(gotDequeued) == len(tC.wantDequeued) {
						bucket.Stop()
						close(out)
						break
					}
				}
			}()

			now := time.Now()
			for o := range out {
				gotDequeued = append(gotDequeued, o)
			}
			gotDequeueTimeElapsed := time.Since(now)

			assert.ElementsMatch(t, tC.wantDequeued, gotDequeued)
			timeComparison := func() bool {
				return gotDequeueTimeElapsed >= tC.wantDequeueTimeElapsedMin && gotDequeueTimeElapsed <= tC.wantDequeueTimeElapsedMax
			}
			assert.Condition(t, timeComparison)
		})
	}
}
