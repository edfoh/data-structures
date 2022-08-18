package strategy

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	fakeNow = time.Date(2022, 01, 01, 12, 0, 0, 0, time.UTC)
)

func TestLeakyBucket(t *testing.T) {

	testCases := []struct {
		desc              string
		tokensPerInterval int
		interval          time.Duration
		capacity          int
		fakeTimeElapsed   time.Time
		initialAddN       int
		afterTimeAddN     int
		wantSuccess       bool
		wantSpillover     int
	}{
		{
			desc:              "when adding to bucket with capacity, should succeed without spillover",
			tokensPerInterval: 1,
			interval:          1 * time.Second,
			capacity:          100,
			fakeTimeElapsed:   fakeNow.Add(5 * time.Second),
			initialAddN:       10,
			afterTimeAddN:     50,
			wantSuccess:       true,
			wantSpillover:     0,
		},
		{
			desc:              "when adding to bucket at capacity, should succeed without spillover",
			tokensPerInterval: 1,
			interval:          10 * time.Second,
			capacity:          100,
			fakeTimeElapsed:   fakeNow.Add(9 * time.Second),
			initialAddN:       100,
			afterTimeAddN:     90,
			wantSuccess:       true,
			wantSpillover:     0,
		},
		{
			desc:              "when adding to bucket that just tips over capacity, should fail with 1 spillover",
			tokensPerInterval: 1,
			interval:          10 * time.Second,
			capacity:          100,
			fakeTimeElapsed:   fakeNow.Add(9 * time.Second),
			initialAddN:       100,
			afterTimeAddN:     91,
			wantSuccess:       false,
			wantSpillover:     1,
		},
		{
			desc:              "when adding full capacity to bucket that is at half capacity, should fail with half capacity spillover",
			tokensPerInterval: 1,
			interval:          10 * time.Second,
			capacity:          100,
			fakeTimeElapsed:   fakeNow.Add(5 * time.Second),
			initialAddN:       100,
			afterTimeAddN:     100,
			wantSuccess:       false,
			wantSpillover:     50,
		},
		{
			desc:              "when empty initially and adding to capacity, should succeed",
			tokensPerInterval: 1,
			interval:          10 * time.Second,
			capacity:          100,
			fakeTimeElapsed:   fakeNow.Add(1 * time.Second),
			initialAddN:       0,
			afterTimeAddN:     100,
			wantSuccess:       true,
			wantSpillover:     0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			teardown := fakeTimeSetup(t)
			defer teardown()

			bucket := NewLeakyBucket(tC.tokensPerInterval, tC.interval, tC.capacity)
			bucket.AddN(tC.initialAddN)
			setFakeNow(tC.fakeTimeElapsed)

			gotSuccess, gotSpillover := bucket.AddN(tC.afterTimeAddN)

			assert.Equal(t, tC.wantSuccess, gotSuccess)
			assert.Equal(t, tC.wantSpillover, gotSpillover)
		})
	}
}

func setFakeNow(t time.Time) {
	now = func() time.Time {
		return t
	}
}

func fakeTimeSetup(t *testing.T) func() {
	t.Helper()

	now = func() time.Time {
		return fakeNow
	}

	return func() {
		now = time.Now
	}
}
