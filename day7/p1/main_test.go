package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack_Push(t *testing.T) {
	tests := []struct {
		name           string
		stack          func() *Stack
		nodes          Nodes
		expectedLength int
	}{
		{"one item added", func() *Stack { return &Stack{} }, Nodes{"A": NewNode("A")}, 1},
		{"two items", func() *Stack { return &Stack{} }, Nodes{"A": NewNode("A"), "B": NewNode("B")}, 2},
		{"two items with duplicates", func() *Stack { return &Stack{"A"} }, Nodes{"A": NewNode("A"), "B": NewNode("B")}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := tt.stack()
			for _, node := range tt.nodes {
				stack.Push(node)
			}
			assert.Equal(t, tt.expectedLength, stack.Len())
		})
	}
}

func TestStack_Pop(t *testing.T) {
	tests := []struct {
		name           string
		stack          func() *Stack
		want           string
		expectedLength int
	}{
		{"no items", func() *Stack { return &Stack{} }, "", 0},
		{"one item", func() *Stack { return &Stack{"A"} }, "A", 0},
		{"two items", func() *Stack { return &Stack{"A", "B"} }, "A", 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := tt.stack()
			if got := stack.Pop(); got != tt.want {
				t.Errorf("Stack.Pop() = %v, want %v", got, tt.want)
			}
			assert.Equal(t, tt.expectedLength, stack.Len())
		})
	}
}

func TestStack_Sort(t *testing.T) {
	tests := []struct {
		name     string
		stack    func() *Stack
		expected func() *Stack
	}{
		{"no items", func() *Stack { return &Stack{} }, func() *Stack { return &Stack{} }},
		{"one item", func() *Stack { return &Stack{"C"} }, func() *Stack { return &Stack{"C"} }},
		{"two items", func() *Stack { return &Stack{"C", "A"} }, func() *Stack { return &Stack{"A", "C"} }},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sorted := tt.stack()
			sorted.Sort()
			assert.Equal(t, sorted, tt.expected())
		})
	}
}
