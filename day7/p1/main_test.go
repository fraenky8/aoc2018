package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue_Add(t *testing.T) {
	tests := []struct {
		name           string
		q              func() *Queue
		nodes          Nodes
		expectedLength int
	}{
		{"one item added", func() *Queue { return &Queue{} }, Nodes{"A": NewNode("A")}, 1},
		{"two items", func() *Queue { return &Queue{} }, Nodes{"A": NewNode("A"), "B": NewNode("B")}, 2},
		{"two items with duplicates", func() *Queue { return &Queue{"A"} }, Nodes{"A": NewNode("A"), "B": NewNode("B")}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := tt.q()
			for _, node := range tt.nodes {
				q.Add(node)
			}
			assert.Equal(t, tt.expectedLength, q.Len())
		})
	}
}

func TestQueue_Dequeue(t *testing.T) {
	tests := []struct {
		name           string
		q              func() *Queue
		want           string
		expectedLength int
	}{
		{"no items", func() *Queue { return &Queue{} }, "", 0},
		{"one item", func() *Queue { return &Queue{"A"} }, "A", 0},
		{"two items", func() *Queue { return &Queue{"A", "B"} }, "A", 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := tt.q()
			if got := q.Dequeue(); got != tt.want {
				t.Errorf("Queue.Dequeue() = %v, want %v", got, tt.want)
			}
			assert.Equal(t, tt.expectedLength, q.Len())
		})
	}
}

func TestQueue_Sort(t *testing.T) {
	tests := []struct {
		name     string
		q        func() *Queue
		expected func() *Queue
	}{
		{"no items", func() *Queue { return &Queue{} }, func() *Queue { return &Queue{} }},
		{"one item", func() *Queue { return &Queue{"C"} }, func() *Queue { return &Queue{"C"} }},
		{"two items", func() *Queue { return &Queue{"C", "A"} }, func() *Queue { return &Queue{"A", "C"} }},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sorted := tt.q()
			sorted.Sort()
			assert.Equal(t, sorted, tt.expected())
		})
	}
}
