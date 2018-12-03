package main

import (
	"bufio"
	"fmt"
	"image"
	"log"
	"os"
	"strconv"
	"strings"
)

type Claim struct {
	id       int
	overlaps bool
	r        image.Rectangle
}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	claims := make([]Claim, 0)

	fs := bufio.NewScanner(file)
	for fs.Scan() {
		line := strings.TrimSpace(fs.Text())
		claims = append(claims, NewClaim(line))
	}

	for i := 0; i < len(claims); i++ {
		for j := i + 1; j < len(claims)-1; j++ {
			if claims[i].r.Intersect(claims[j].r) != image.ZR {
				claims[i].overlaps = true
				claims[j].overlaps = true
				break
			}
		}
	}

	var c Claim
	for _, claim := range claims {
		if !claim.overlaps {
			c = claim
			break
		}
	}

	fmt.Printf("non overlapping claim: %+v\n", c)
}

func NewClaim(line string) Claim {
	// #123 @ 3,2: 5x4
	splitted := strings.Split(line, " ")

	if len(splitted) != 4 {
		log.Fatalf("wrong length %v of line %v\n", len(splitted), line)
	}

	id, err := strconv.Atoi(splitted[0][1:])
	if err != nil {
		log.Fatalf("could not convert to id: %v\n", err)
	}

	commaPos := strings.Index(splitted[2], ",")

	x0, err := strconv.Atoi(splitted[2][0:commaPos])
	if err != nil {
		log.Fatalf("could not convert to x0: %v\n", err)
	}

	y0, err := strconv.Atoi(splitted[2][commaPos+1 : len(splitted[2])-1])
	if err != nil {
		log.Fatalf("could not convert to y0: %v\n", err)
	}

	xPos := strings.Index(splitted[3], "x")

	w, err := strconv.Atoi(splitted[3][0:xPos])
	if err != nil {
		log.Fatalf("could not convert to width: %v\n", err)
	}

	h, err := strconv.Atoi(splitted[3][xPos+1:])
	if err != nil {
		log.Fatalf("could not convert to height: %v\n", err)
	}

	return Claim{
		id: id,
		r:  image.Rect(x0, y0, x0+w, y0+h),
	}
}
