package main

import (
	"image"
	"reflect"
	"testing"
)

func TestNewClaim(t *testing.T) {
	tests := []struct {
		name string
		line string
		want Claim
	}{
		{"example 1", "#123 @ 3,2: 5x4", Claim{id: 123, r: image.Rect(3, 2, 8, 6)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClaim(tt.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClaim() = %v, want %v", got, tt.want)
			}
		})
	}
}
