package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, _, checksum := checksum(file)

	fmt.Printf("checksum: %v\n", checksum)
}

func checksum(r io.Reader) (twos int, threes int, checksum int) {

	s := bufio.NewScanner(r)
	for s.Scan() {
		m := map[rune]int{}

		s := strings.TrimSpace(s.Text())
		for _, c := range s {
			m[c]++
		}

		twosCounted := false
		threesCounted := false

		for _, i := range m {
			if i == 2 && !twosCounted {
				twos++
				twosCounted = true
			} else if i == 3 && !threesCounted {
				threes++
				threesCounted = true
			}
		}
	}

	return twos, threes, twos * threes
}
