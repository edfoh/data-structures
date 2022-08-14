package datastruct

import (
	"errors"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPriorityQueue_Enqueue(t *testing.T) {
	tests := []struct {
		name         string
		priorities   []int
		wantDequeues []string
		wantErrs     []error
	}{
		{
			name:         "enqueue and dequeue works",
			priorities:   []int{50, 20, 10, 40, 80},
			wantDequeues: []string{"80", "50", "40", "20", "10"},
			wantErrs:     []error{nil, nil, nil, nil, nil},
		},
		{
			name:         "dequeue error when nothing to dequeue",
			priorities:   []int{50},
			wantDequeues: []string{"50", ""},
			wantErrs:     []error{nil, errors.New("heap is empty")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pq := NewPriorityQueue(10)

			for _, priority := range tt.priorities {
				pq.Enqueue(strconv.Itoa(priority), priority)
			}

			for i, wantDequeue := range tt.wantDequeues {
				gotDequeue, gotErr := pq.Dequeue()
				require.Equal(t, tt.wantErrs[i], gotErr)
				assert.Equal(t, wantDequeue, gotDequeue)
			}
		})
	}
}
