package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	inf = 9999999999

	example = `1, 1
1, 6
8, 3
3, 4
5, 5
8, 9`
)

type Point struct {
	X, Y int
}
type Points []Point

type Set map[Point]bool
type Data struct {
	Point    Point
	Distance int
}
type Datas []Data
type Row []Datas
type Matrix []Row

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	points := NewPoints(file)
	//points := NewPoints(strings.NewReader(example))

	matrix := CreateMatrix(points)

	infinities := points.Infinities()

	areas := map[Point]int{}

	for _, row := range matrix {
		for _, cell := range row {

			min := cell[0]
			seen := map[int]int{}

			for _, data := range cell {
				seen[data.Distance]++
				if data.Distance < min.Distance {
					min = data
				}
			}

			overlaps := false
			for distance, count := range seen {
				if count > 1 && distance == min.Distance {
					overlaps = true
				}
			}

			if !overlaps {
				areas[min.Point]++
			}
		}
	}

	largestArea := 0
	for p, area := range areas {
		if !IsInfinityPoint(p, infinities) && area > largestArea {
			largestArea = area
		}
	}

	fmt.Printf("size of the largest area: %v\n", largestArea)
}

func CreateMatrix(points Points) Matrix {

	minX := points.MinX()
	maxX := points.MaxX()

	minY := points.MinY()
	maxY := points.MaxY()

	var matrix Matrix

	for y := minY; y < maxY; y++ {
		r := Row{}
		for x := minX; x < maxX; x++ {
			d := Datas{}
			for _, point := range points {
				d = append(d, Data{Point: point, Distance: Distance(point, Point{X: x, Y: y})})
			}
			r = append(r, d)
		}
		matrix = append(matrix, r)
	}

	return matrix
}

func Distance(p1, p2 Point) int {
	return Abs(p1.X-p2.X) + Abs(p1.Y-p2.Y)
}

func Abs(i int) int {
	if i >= 0 {
		return i
	}
	return i * -1
}

func (ps Points) Infinities() Set {
	infs := Set{}

	tms := ps.TopMosts()
	for _, p := range tms {
		infs[p] = true
	}

	lms := ps.LeftMosts()
	for _, p := range lms {
		infs[p] = true
	}

	bms := ps.BottomMosts()
	for _, p := range bms {
		infs[p] = true
	}

	rms := ps.RightMosts()
	for _, p := range rms {
		infs[p] = true
	}

	return infs
}

func IsInfinityPoint(point Point, points Set) bool {
	ok, _ := points[point]
	return ok
}

func (ps Points) TopMosts() Points {

	minY := ps.MinY()
	points := Points{}

	for _, p := range ps {
		if p.Y == minY {
			points = append(points, p)
		}
	}

	return points
}

func (ps Points) LeftMosts() Points {

	minX := ps.MinX()
	points := Points{}

	for _, p := range ps {
		if p.X == minX {
			points = append(points, p)
		}
	}

	return points
}

func (ps Points) BottomMosts() Points {

	maxY := ps.MaxY()
	points := Points{}

	for _, p := range ps {
		if p.Y == maxY {
			points = append(points, p)
		}
	}

	return points
}

func (ps Points) RightMosts() Points {

	maxX := ps.MaxX()
	points := Points{}

	for _, p := range ps {
		if p.X == maxX {
			points = append(points, p)
		}
	}

	return points
}

func NewPoints(r io.Reader) Points {
	points := Points{}
	fs := bufio.NewScanner(r)
	for fs.Scan() {
		points = append(points, NewPoint(fs.Text()))
	}
	return points
}

func NewPoint(line string) Point {
	line = strings.TrimSpace(line)
	if line == "" {
		log.Fatalf("empty line detected!")
	}
	splitted := strings.Split(line, ", ")

	x, err := strconv.Atoi(splitted[0])
	if err != nil {
		log.Fatalf("could not convert x %v: %v", splitted[0], err)
	}

	y, err := strconv.Atoi(splitted[1])
	if err != nil {
		log.Fatalf("could not convert y %v: %v", splitted[1], err)
	}

	return Point{X: x, Y: y}
}

func (ps Points) MaxX() int {
	max := 0
	for _, p := range ps {
		if p.X > max {
			max = p.X
		}
	}
	return max
}

func (ps Points) MinX() int {
	min := inf
	for _, p := range ps {
		if p.X < min {
			min = p.X
		}
	}
	return min
}

func (ps Points) MaxY() int {
	max := 0
	for _, p := range ps {
		if p.Y > max {
			max = p.Y
		}
	}
	return max
}

func (ps Points) MinY() int {
	min := inf
	for _, p := range ps {
		if p.Y < min {
			min = p.Y
		}
	}
	return min
}
