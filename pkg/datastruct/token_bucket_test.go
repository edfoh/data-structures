package datastruct

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTokenBucket_TakeN(t *testing.T) {
	tests := []struct {
		name                string
		maxTokens           float64
		refillRatePerSecond float64
		takesN              []int
		takeNSleep          time.Duration
		wantErrs            []error
	}{
		{
			name:                "taking less than max tokens works",
			maxTokens:           10,
			refillRatePerSecond: 1,
			takesN:              []int{5},
			takeNSleep:          0,
			wantErrs:            []error{nil},
		},
		{
			name:                "taking more than max token fails",
			maxTokens:           10,
			refillRatePerSecond: 0.1,
			takesN:              []int{11},
			takeNSleep:          0,
			wantErrs:            []error{errors.New("ran out of tokens")},
		},
		{
			name:                "taking max tokens, refill should allow take max again",
			maxTokens:           10,
			refillRatePerSecond: 10,
			takesN:              []int{10, 10},
			takeNSleep:          1 * time.Second,
			wantErrs:            []error{nil, nil},
		},
		{
			name:                "taking max tokens, refill at slower rate, take max should fail",
			maxTokens:           10,
			refillRatePerSecond: 8,
			takesN:              []int{10, 10},
			takeNSleep:          1 * time.Second,
			wantErrs:            []error{nil, errors.New("ran out of tokens")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bucket := NewTokenBucket(tt.maxTokens, tt.refillRatePerSecond)

			for i, takeN := range tt.takesN {
				gotErr := bucket.TakeN(takeN)
				assert.Equal(t, tt.wantErrs[i], gotErr)
				time.Sleep(tt.takeNSleep)
			}
		})
	}
}
