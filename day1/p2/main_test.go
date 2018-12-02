package main

import "testing"

func Test_twice(t *testing.T) {
	tests := []struct {
		name     string
		ints     []int
		expected int
	}{
		{"example 0", []int{1, -1}, 0},
		{"example 10", []int{+3, +3, +4, -2, -4}, 10},
		{"example 5", []int{-6, +3, +8, +5, -6}, 5},
		{"example 14", []int{+7, +7, -2, -7, -4}, 14},
		{"example 2", []int{1, -2, 3, 1}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if actual := twice(tt.ints); actual != tt.expected {
				t.Errorf("twice() = %v, expected %v", actual, tt.expected)
			}
		})
	}
}
