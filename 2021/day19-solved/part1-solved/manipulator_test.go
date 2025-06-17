package day

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		name      string
		rotationX int
		beacon    [3]int
		want      [3]int
	}{
		{
			name:      "no rotation",
			rotationX: 0,
			beacon:    [3]int{0, 1, 2},
			want:      [3]int{0, 1, 2},
		},
		{
			name:      "0, 2, 1",
			rotationX: 1,
			beacon:    [3]int{0, 1, 2},
			want:      [3]int{0, 2, 1},
		},
		{
			name:      "1, 0, 2",
			rotationX: 2,
			beacon:    [3]int{0, 1, 2},
			want:      [3]int{1, 0, 2},
		},
		{
			name:      "1, 2, 0",
			rotationX: 3,
			beacon:    [3]int{0, 1, 2},
			want:      [3]int{1, 2, 0},
		},
		{
			name:      "2, 0, 1",
			rotationX: 4,
			beacon:    [3]int{0, 1, 2},
			want:      [3]int{2, 0, 1},
		},
		{
			name:      "2, 1, 0",
			rotationX: 5,
			beacon:    [3]int{0, 1, 2},
			want:      [3]int{2, 1, 0},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			m := manipulator{
				rotationX: tc.rotationX,
			}
			got := m.rotate(tc.beacon)
			assert.Equal(t, tc.want, got, "rotate(%v) by %v => %v; want %v", tc.beacon, allRotations[tc.rotationX], got, tc.want)
		})
	}
}

func TestFlip(t *testing.T) {
	testCases := []struct {
		name       string
		flipationX int
		beacon     [3]int
		want       [3]int
	}{
		{
			name:       "no flip",
			flipationX: 0,
			beacon:     [3]int{1, 2, 3},
			want:       [3]int{1, 2, 3},
		},
		{
			name:       "flip x",
			flipationX: 1,
			beacon:     [3]int{1, 2, 3},
			want:       [3]int{-1, 2, 3},
		},
		{
			name:       "flip y",
			flipationX: 2,
			beacon:     [3]int{1, 2, 3},
			want:       [3]int{1, -2, 3},
		},
		{
			name:       "flip z",
			flipationX: 3,
			beacon:     [3]int{1, 2, 3},
			want:       [3]int{1, 2, -3},
		},
		{
			name:       "flip x and y",
			flipationX: 4,
			beacon:     [3]int{1, 2, 3},
			want:       [3]int{-1, -2, 3},
		},
		{
			name:       "flip x and z",
			flipationX: 5,
			beacon:     [3]int{1, 2, 3},
			want:       [3]int{-1, 2, -3},
		},
		{
			name:       "flip y and z",
			flipationX: 6,
			beacon:     [3]int{1, 2, 3},
			want:       [3]int{1, -2, -3},
		},
		{
			name:       "flip all",
			flipationX: 7,
			beacon:     [3]int{1, 2, 3},
			want:       [3]int{-1, -2, -3},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			m := manipulator{
				flipationX: tc.flipationX,
			}
			got := m.flip(tc.beacon)
			assert.Equal(t, tc.want, got, "flip(%v) by %v => %v; want %v", tc.beacon, allFlips[tc.flipationX], got, tc.want)
		})
	}
}

func TestTranslate(t *testing.T) {
	testCases := []struct {
		name        string
		translation [3]int
		beacon      [3]int
		want        [3]int
	}{
		{
			name:        "none",
			translation: [3]int{0, 0, 0},
			beacon:      [3]int{4, 5, 6},
			want:        [3]int{4, 5, 6},
		},
		{
			name:        "positive",
			translation: [3]int{3, 4, 5},
			beacon:      [3]int{4, 5, 6},
			want:        [3]int{7, 9, 11},
		},
		{
			name:        "negative",
			translation: [3]int{-1, -2, -3},
			beacon:      [3]int{4, 5, 6},
			want:        [3]int{3, 3, 3},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			m := manipulator{
				translation: tc.translation,
			}
			got := m.translate(tc.beacon)
			assert.Equal(t, tc.want, got, "translate(%v) by %v => %v; want %v", tc.beacon, tc.translation, got, tc.want)
		})
	}
}

func TestGetTranslation(t *testing.T) {
	testCases := []struct {
		name string
		b1   [3]int
		b2   [3]int
		want [3]int
	}{
		{
			name: "equal",
			b1:   [3]int{0, 0, 0},
			b2:   [3]int{0, 0, 0},
			want: [3]int{0, 0, 0},
		},
		{
			name: "positive diff",
			b1:   [3]int{10, 20, 30},
			b2:   [3]int{4, 5, 6},
			want: [3]int{6, 15, 24},
		},
		{
			name: "negative diff",
			b1:   [3]int{0, 0, 0},
			b2:   [3]int{4, 5, 6},
			want: [3]int{-4, -5, -6},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			m := manipulator{}
			got := m.getTranslation(tc.b1, tc.b2)
			assert.Equal(t, tc.want, got, "b1 %v minus b2 %v => %v; want %v", tc.b1, tc.b2, got, tc.want)
		})
	}
}
