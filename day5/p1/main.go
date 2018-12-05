package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	b, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	b = bytes.Replace(b, []byte("\n"), []byte(""), -1)

	out := reduce(string(b))

	fmt.Println(out)
	fmt.Println(len(out))
}

func reduce(in string) string {
	if in == "" {
		return ""
	}

	changed := false

	for i := 0; i < len(in)-1; i++ {
		if !IsDestroyable(in[i], in[i+1]) {
			continue
		}

		in = in[:i] + in[i+2:]
		changed = true
	}

	if !changed {
		return in
	}

	return reduce(in)
}

func IsDestroyable(a, b byte) bool {
	if !IsSameType(a, b) {
		return false
	}

	if IsUpper(a) && IsLower(b) {
		return true
	}

	if IsLower(a) && IsUpper(b) {
		return true
	}

	return false
}

func IsUpper(b byte) bool {
	if !IsAscii(b) {
		return false
	}
	bb := []byte{b}
	return bytes.ToUpper(bb)[0] == bb[0]
}

func IsLower(b byte) bool {
	return !IsUpper(b)
}

func IsAscii(b byte) bool {
	return b >= 'A' && (b <= 'Z' || b >= 'a') && b <= 'z'
}

func IsSameType(a, b byte) bool {
	if !IsAscii(a) || !IsAscii(b) {
		return false
	}
	aa := []byte{a}
	bb := []byte{b}
	return bytes.ToUpper(aa)[0] == bytes.ToUpper(bb)[0]
}
