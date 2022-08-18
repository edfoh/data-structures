package strategy

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSlidingWindow(t *testing.T) {

	testCases := []struct {
		desc          string
		interval      time.Duration
		capacity      int
		sleepInterval time.Duration
		beforeAddN    int
		afterAddN     int
		wantErr       error
		wantCanAdd    bool
		wantPrevCount int
		wantCurrCount int
	}{
		{
			desc:          "AddN in the current window below capacity should pass",
			interval:      1 * time.Second,
			capacity:      10,
			sleepInterval: 500 * time.Millisecond,
			beforeAddN:    5,
			afterAddN:     3,
			wantCanAdd:    true,
			wantPrevCount: 0,
			wantCurrCount: 8,
		},
		{
			desc:          "AddN in the current window over the capacity should fail",
			interval:      1 * time.Second,
			capacity:      10,
			sleepInterval: 500 * time.Millisecond,
			beforeAddN:    5,
			afterAddN:     9,
			wantCanAdd:    false,
			wantPrevCount: 0,
			wantCurrCount: 5,
		},
		{
			desc:          "AddN in the current window below capacity in the second window should pass",
			interval:      1 * time.Second,
			capacity:      10,
			sleepInterval: 1500 * time.Millisecond,
			beforeAddN:    5,
			afterAddN:     4,
			wantCanAdd:    true,
			wantPrevCount: 5,
			wantCurrCount: 4,
		},
		{
			desc:          "AddN in the current window above capacity in the second window should pass",
			interval:      1 * time.Second,
			capacity:      10,
			sleepInterval: 1500 * time.Millisecond,
			beforeAddN:    5,
			afterAddN:     10,
			wantCanAdd:    false,
			wantPrevCount: 5,
			wantCurrCount: 0,
		},
		{
			desc:          "AddN in the current window below capacity in the 3rd window should pass",
			interval:      1 * time.Second,
			capacity:      10,
			sleepInterval: 2500 * time.Millisecond,
			beforeAddN:    5,
			afterAddN:     5,
			wantCanAdd:    true,
			wantPrevCount: 0,
			wantCurrCount: 5,
		},
		{
			desc:          "AddN in the current window over capacity in the 3rd window should pass",
			interval:      1 * time.Second,
			capacity:      10,
			sleepInterval: 2500 * time.Millisecond,
			beforeAddN:    5,
			afterAddN:     15,
			wantCanAdd:    false,
			wantPrevCount: 0,
			wantCurrCount: 0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			sl := NewSlidingWindow(tC.interval, tC.capacity)

			_, gotErr := sl.AddN(tC.beforeAddN)
			require.NoError(t, gotErr)

			time.Sleep(tC.sleepInterval)

			gotCanAdd, gotErr := sl.AddN(tC.afterAddN)
			sl.Stop()

			assert.Equal(t, tC.wantCanAdd, gotCanAdd)
			assert.Equal(t, tC.wantErr, gotErr)

			gotPrevCount, gotCurrCount := sl.GetCount()
			assert.Equal(t, tC.wantPrevCount, gotPrevCount)
			assert.Equal(t, tC.wantCurrCount, gotCurrCount)
		})
	}
}

func TestSyncSlidingWindow(t *testing.T) {
	capacity := 10
	interval := 1 * time.Minute
	testCases := []struct {
		desc            string
		firstAddN       int
		secondAddN      int
		fakeTimeElapsed time.Duration
		wantCanAdd      bool
		wantErr         error
		wantPrevCount   int
		wantCurrCount   int
	}{
		{
			desc:            "adding twice in curr window works",
			firstAddN:       5,
			secondAddN:      4,
			fakeTimeElapsed: 59 * time.Second,
			wantCanAdd:      true,
			wantPrevCount:   0,
			wantCurrCount:   9,
		},
		{
			desc:            "over capacity in curr window returns error",
			firstAddN:       5,
			secondAddN:      6,
			fakeTimeElapsed: 59 * time.Second,
			wantCanAdd:      false,
			wantErr:         errors.New("sliding window is full"),
			wantPrevCount:   0,
			wantCurrCount:   5,
		},
		{
			desc:            "adding in both windows halfway under capacity works",
			firstAddN:       5,
			secondAddN:      5,
			fakeTimeElapsed: 90 * time.Second,
			wantCanAdd:      true,
			wantErr:         nil,
			wantPrevCount:   5,
			wantCurrCount:   5,
		},
		{
			desc:            "adding in both windows halfway over capacity returns erros",
			firstAddN:       5,
			secondAddN:      6,
			fakeTimeElapsed: 90 * time.Second,
			wantCanAdd:      false,
			wantErr:         errors.New("sliding window is full"),
			wantPrevCount:   5,
			wantCurrCount:   0,
		},
		{
			desc:            "adding past 1 window under capacity works",
			firstAddN:       5,
			secondAddN:      5,
			fakeTimeElapsed: 179 * time.Second,
			wantCanAdd:      true,
			wantErr:         nil,
			wantPrevCount:   0,
			wantCurrCount:   5,
		},
		{
			desc:            "adding past 1 window over capacity returns error",
			firstAddN:       5,
			secondAddN:      11,
			fakeTimeElapsed: 179 * time.Second,
			wantCanAdd:      false,
			wantErr:         errors.New("sliding window is full"),
			wantPrevCount:   0,
			wantCurrCount:   0,
		},
		{
			desc:            "adding past 2 windows under capacity works",
			firstAddN:       5,
			secondAddN:      5,
			fakeTimeElapsed: 239 * time.Second,
			wantCanAdd:      true,
			wantErr:         nil,
			wantPrevCount:   0,
			wantCurrCount:   5,
		},
		{
			desc:            "adding past 2 windows over capacity returns error",
			firstAddN:       5,
			secondAddN:      11,
			fakeTimeElapsed: 239 * time.Second,
			wantCanAdd:      false,
			wantErr:         errors.New("sliding window is full"),
			wantPrevCount:   0,
			wantCurrCount:   0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			teardown := fakeTimeSetup(t)
			defer teardown()

			sl := NewSyncSlidingWindow(interval, capacity)
			_, gotErr := sl.AddN(tC.firstAddN)
			require.NoError(t, gotErr)

			setFakeNow(fakeNow.Add(tC.fakeTimeElapsed))

			gotCanAdd, gotErr := sl.AddN(tC.secondAddN)
			assert.Equal(t, tC.wantCanAdd, gotCanAdd)
			assert.Equal(t, tC.wantErr, gotErr)

			gotPrevCount, gotCurrCount := sl.Count()
			assert.Equal(t, tC.wantPrevCount, gotPrevCount)
			assert.Equal(t, tC.wantCurrCount, gotCurrCount)
		})
	}
}
