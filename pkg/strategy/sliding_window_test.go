package strategy

import (
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
