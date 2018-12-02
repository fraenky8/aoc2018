package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	ints := []int{}

	fs := bufio.NewScanner(file)
	for fs.Scan() {
		s := strings.TrimSpace(fs.Text())
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		ints = append(ints, n)
	}

	frequency := twice(ints)

	fmt.Printf("twice: %v\n", frequency)
}

func twice(ints []int) int {

	frequency := 0
	seen := map[int]bool{}
	abort := false

	seen[frequency] = true
	for !abort {
		for _, n := range ints {
			frequency += n
			if _, ok := seen[frequency]; ok {
				abort = true
				break
			}
			seen[frequency] = true
		}
	}

	return frequency
}
