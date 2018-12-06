package main

import "testing"

func Test_reduceBy(t *testing.T) {
	tests := []struct {
		name string
		in   string
		r    rune
		want string
	}{
		{"removing a/A", "dabAcCaCBAcCcaDA", 'a', "dbCBcD"},
		{"removing b/B", "dabAcCaCBAcCcaDA", 'b', "daCAcaDA"},
		{"removing c/C", "dabAcCaCBAcCcaDA", 'c', "daDA"},
		{"removing d/D", "dabAcCaCBAcCcaDA", 'd', "abCBAc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reduceBy(tt.in, tt.r); got != tt.want {
				t.Errorf("reduceBy() = %v, want %v", got, tt.want)
			}
		})
	}
}
