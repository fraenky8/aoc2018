package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"strconv"
	"strings"
)

type Claim struct {
	id int
	r  image.Rectangle
}

type Claims []Claim

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	claims := make(Claims, 0)

	fs := bufio.NewScanner(file)
	for fs.Scan() {
		line := strings.TrimSpace(fs.Text())
		claims = append(claims, NewClaim(line))
	}

	// Thanks to my brilliant girlfriend Sarah <3

	fabric := [1000][1000]int{}

	for _, claim := range claims {
		for y := claim.r.Min.Y; y < claim.r.Max.Y; y++ {
			for x := claim.r.Min.X; x < claim.r.Max.X; x++ {
				fabric[x][y]++
			}
		}
	}

	// create a heatmap for fun :D
	img := image.NewRGBA(image.Rect(0, 0, 1000, 1000))

	in2s := 0
	for i := range fabric {
		for j := range fabric[i] {
			if fabric[i][j] > 1 {
				in2s++
			}
			c := 255 - uint8(fabric[i][j]*30)
			img.Set(i, j, color.RGBA{c, c, c, 0xff})
		}
	}

	fmt.Printf("square inches: %v\n", in2s)

	f, err := os.Create("image.png")
	if err != nil {
		log.Fatal(err)
	}

	err = png.Encode(f, img)
	if err != nil {
		log.Fatal(err)
	}
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
