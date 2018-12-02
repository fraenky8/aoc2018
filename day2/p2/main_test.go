package main

import (
	"reflect"
	"testing"
)

func Test_common(t *testing.T) {
	tests := []struct {
		name       string
		bs         [][]byte
		wantId1    []byte
		wantId2    []byte
		wantCommon string
		wantErr    bool
	}{
		{"empty input has no correct ids", [][]byte{}, nil, nil, "", true},
		{"example", [][]byte{
			[]byte("abcde"), []byte("fghij"), []byte("klmno"), []byte("pqrst"), []byte("fguij"), []byte("axcye"), []byte("wvxyz"),
		}, []byte("fghij"), []byte("fguij"), "fgij", false},
		{"input without correct ids", [][]byte{
			[]byte("abcde"), []byte("fghij"), []byte("klmno"), []byte("pqrst"), []byte("fguik"), []byte("axcye"), []byte("wvxyz"),
		}, nil, nil, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotId1, gotId2, gotCommon, err := common(tt.bs)
			if (err != nil) != tt.wantErr {
				t.Errorf("common() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotId1, tt.wantId1) {
				t.Errorf("common() gotId1 = %v, want %v", gotId1, tt.wantId1)
			}
			if !reflect.DeepEqual(gotId2, tt.wantId2) {
				t.Errorf("common() gotId2 = %v, want %v", gotId2, tt.wantId2)
			}
			if gotCommon != tt.wantCommon {
				t.Errorf("common() gotCommon = %v, want %v", gotCommon, tt.wantCommon)
			}
		})
	}
}
