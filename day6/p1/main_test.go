package main

import (
	"reflect"
	"strings"
	"testing"

	"github.com/magiconair/properties/assert"
)

const (
	example_wo_CF = `1, 1
1, 6
3, 4
5, 5`

	example_wo_F = `1, 1
1, 6
8, 3
3, 4
5, 5`
)

func TestPoints_MaxX(t *testing.T) {
	tests := []struct {
		name      string
		ps        Points
		expectedX int
	}{
		{"example", NewPoints(strings.NewReader(example)), 8},
		{"example without CF", NewPoints(strings.NewReader(example_wo_CF)), 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expectedX, tt.ps.MaxX())
		})
	}
}

func TestPoints_MaxY(t *testing.T) {
	tests := []struct {
		name      string
		ps        Points
		expectedY int
	}{
		{"example", NewPoints(strings.NewReader(example)), 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expectedY, tt.ps.MaxY())
		})
	}
}

func TestPoints_TopMosts(t *testing.T) {
	tests := []struct {
		name  string
		ps    Points
		wantP Points
	}{
		{"example A", NewPoints(strings.NewReader(example)), Points{{X: 1, Y: 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotP := tt.ps.TopMosts(); !reflect.DeepEqual(gotP, tt.wantP) {
				t.Errorf("Points.TopMosts() = %v, want %v", gotP, tt.wantP)
			}
		})
	}
}

func TestPoints_LeftMosts(t *testing.T) {
	tests := []struct {
		name  string
		ps    Points
		wantP Points
	}{
		{"example A", NewPoints(strings.NewReader(example)), Points{{X: 1, Y: 1}, {X: 1, Y: 6}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotP := tt.ps.LeftMosts(); !reflect.DeepEqual(gotP, tt.wantP) {
				t.Errorf("Points.LeftMosts() = %v, want %v", gotP, tt.wantP)
			}
		})
	}
}

func TestPoints_BottomMost(t *testing.T) {
	tests := []struct {
		name  string
		ps    Points
		wantP Points
	}{
		{"example A", NewPoints(strings.NewReader(example)), Points{{X: 8, Y: 9}}},
		{"example without F", NewPoints(strings.NewReader(example_wo_F)), Points{{X: 1, Y: 6}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotP := tt.ps.BottomMosts(); !reflect.DeepEqual(gotP, tt.wantP) {
				t.Errorf("Points.BottomMosts() = %v, want %v", gotP, tt.wantP)
			}
		})
	}
}

func TestPoints_RightMost(t *testing.T) {
	tests := []struct {
		name  string
		ps    Points
		wantP Points
	}{
		{"example A", NewPoints(strings.NewReader(example)), Points{{X: 8, Y: 3}, {X: 8, Y: 9}}},
		{"example without C", NewPoints(strings.NewReader(example_wo_CF)), Points{{X: 5, Y: 5}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotP := tt.ps.RightMosts(); !reflect.DeepEqual(gotP, tt.wantP) {
				t.Errorf("Points.RightMosts() = %v, want %v", gotP, tt.wantP)
			}
		})
	}
}

func TestPoint_Distance(t *testing.T) {
	tests := []struct {
		name string
		p1   Point
		p2   Point
		want int
	}{
		{"A, B", Point{X: 1, Y: 1}, Point{X: 1, Y: 6}, 5},
		{"A, C", Point{X: 1, Y: 1}, Point{X: 8, Y: 3}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Distance(tt.p1, tt.p2); got != tt.want {
				t.Errorf("Point.Distance() = %v, want %v", got, tt.want)
			}
		})
	}
}
