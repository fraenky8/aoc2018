package main

import (
	"io"
	"strings"
	"testing"
)

func Test_checksum(t *testing.T) {
	tests := []struct {
		name         string
		r            io.Reader
		wantTwos     int
		wantThrees   int
		wantChecksum int
	}{
		{"no twos, no threes", strings.NewReader("abcdef"), 0, 0, 0},
		{"one twos, one threes", strings.NewReader("bababc"), 1, 1, 1},
		{"one twos, zero threes", strings.NewReader("abbcde"), 1, 0, 0},
		{"zero twos, one threes", strings.NewReader("abcccd"), 0, 1, 0},
		{"two twos, zero threes, twos count only once", strings.NewReader("aabcdd"), 1, 0, 0},
		{"one twos, zero threes", strings.NewReader("abcdee"), 1, 0, 0},
		{"zero twos, two threes, threes count only one", strings.NewReader("ababab"), 0, 1, 0},
		{"all together", strings.NewReader("abcdef\nbababc\nabbcde\nabcccd\naabcdd\nabcdee\nababab"), 4, 3, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTwos, gotThrees, gotChecksum := checksum(tt.r)
			if gotTwos != tt.wantTwos {
				t.Errorf("checksum() gotTwos = %v, want %v", gotTwos, tt.wantTwos)
			}
			if gotThrees != tt.wantThrees {
				t.Errorf("checksum() gotThrees = %v, want %v", gotThrees, tt.wantThrees)
			}
			if gotChecksum != tt.wantChecksum {
				t.Errorf("checksum() gotChecksum = %v, want %v", gotChecksum, tt.wantChecksum)
			}
		})
	}
}
