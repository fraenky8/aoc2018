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

	frequency := 0

	fs := bufio.NewScanner(file)
	for fs.Scan() {
		s := strings.TrimSpace(fs.Text())
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		frequency += n
	}

	fmt.Printf("frequency: %v\n", frequency)
}
