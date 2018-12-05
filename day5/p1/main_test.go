package main

import (
	"reflect"
	"testing"
)

func Test_reduce(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		wantOut string
	}{
		{"simple example 1", "aA", ""},
		{"simple example 2", "abBA", ""},
		{"simple example 3", "abAB", "abAB"},
		{"simple example 4", "aabAAB", "aabAAB"},
		{"complex example", "dabAcCaCBAcCcaDA", "dabCBAcaDA"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := reduce(tt.in); !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("reduce() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func TestIsUpper(t *testing.T) {
	tests := []struct {
		name string
		b    byte
		want bool
	}{
		{"space should return false", ' ', false},
		{"lower case should return false", 'a', false},
		{"upper case should return true", 'A', true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsUpper(tt.b); got != tt.want {
				t.Errorf("IsUpper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsAscii(t *testing.T) {
	tests := []struct {
		name string
		b    byte
		want bool
	}{
		{"space should return false", ' ', false},
		{"number should return false", '7', false},
		{"square bracket should return false", '[', false},
		{"tick should return false", '`', false},
		{"hash should return false", '#', false},
		{"upper A character should return true", 'A', true},
		{"upper M character should return true", 'M', true},
		{"upper Z character should return true", 'Z', true},
		{"lower a character should return true", 'a', true},
		{"lower m character should return true", 'm', true},
		{"lower z character should return true", 'z', true},
		{"curly bracket should return true", '{', false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAscii(tt.b); got != tt.want {
				t.Errorf("IsAscii() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsSameType(t *testing.T) {
	tests := []struct {
		name string
		a    byte
		b    byte
		want bool
	}{
		{"first one non ascii should not be same type", ' ', 'b', false},
		{"second one non ascii should not be same type", 'a', ' ', false},
		{"both non ascii should not be same type", '$', '#', false},
		{"not same type", 'a', 'b', false},
		{"both lower case should be same type", 'a', 'a', true},
		{"both upper case should be same type", 'A', 'A', true},
		{"mixed upper/lower case should be same type", 'a', 'A', true},
		{"mixed upper/lower case should be same type", 'A', 'a', true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSameType(tt.a, tt.b); got != tt.want {
				t.Errorf("IsSameType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsDestroyable(t *testing.T) {
	tests := []struct {
		name string
		a    byte
		b    byte
		want bool
	}{
		{"wrong type should not be destroyable", 'a', 'b', false},
		{"same type but both lower case should not be destroyable", 'a', 'a', false},
		{"same type but both upper case should not be destroyable", 'A', 'A', false},
		{"first one lower, second one upper case should be destroyable", 'a', 'A', true},
		{"first one upper, second one lower case should be destroyable", 'A', 'a', true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsDestroyable(tt.a, tt.b); got != tt.want {
				t.Errorf("IsDestroyable() = %v, want %v", got, tt.want)
			}
		})
	}
}
