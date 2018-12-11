package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

const (
	example = `Step C must be finished before step A can begin.
Step C must be finished before step F can begin.
Step A must be finished before step B can begin.
Step A must be finished before step D can begin.
Step B must be finished before step E can begin.
Step D must be finished before step E can begin.
Step F must be finished before step E can begin.
`
)

type Node struct {
	Letter   string
	Visited  bool
	Requires Nodes
}

type Nodes map[string]*Node

type Queue []string

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Thanks to my brilliant girlfriend Sarah <3

	graph := NewGraph(file)
	//graph := NewGraph(strings.NewReader(example))

	start := GetStartNode(graph)
	queue := &Queue{start.Letter}

	var solution strings.Builder

	for !queue.IsEmpty() {

		current := queue.Dequeue()
		solution.WriteString(current)

		graph[current].Visited = true

		enabledNodes := GetAllEnabledNodes(graph)
		for _, node := range enabledNodes {
			queue.Add(node)
		}
		queue.Sort()
	}

	fmt.Printf("order of instructions: %v\n", solution.String())
}

func NewGraph(r io.Reader) Nodes {
	nodes := Nodes{}

	// Step C must be finished before step A can begin.
	// Therefore: A has requirement of C
	fs := bufio.NewScanner(r)
	for fs.Scan() {
		line := strings.TrimSpace(fs.Text())

		required := line[5:6]
		letter := line[36:37]

		requiredNode, exists := nodes[required]
		if !exists {
			requiredNode = NewNode(required)
			nodes[required] = requiredNode
		}

		node, exists := nodes[letter]
		if !exists {
			node = NewNode(letter)
			nodes[letter] = node
		}
		node.Requires[required] = requiredNode
	}

	return nodes
}

func GetStartNode(nodes Nodes) *Node {
	for _, node := range nodes {
		if len(node.Requires) == 0 {
			return node
		}
	}
	return nil
}

func GetAllEnabledNodes(graph Nodes) Nodes {
	enabled := Nodes{}
	for _, node := range graph {
		if IsEnabled(node) {
			enabled[node.Letter] = node
		}
	}
	return enabled
}

func IsEnabled(node *Node) bool {
	if node.Visited {
		return false
	}
	for _, required := range node.Requires {
		if !required.Visited {
			return false
		}
	}
	return true
}

func NewNode(letter string) *Node {
	return &Node{Letter: letter, Requires: make(Nodes)}
}

func (q *Queue) Add(node *Node) {
	for _, s := range *q {
		if s == node.Letter {
			return
		}
	}
	*q = append(*q, node.Letter)
}

func (q Queue) Len() int {
	return len(q)
}

func (q *Queue) Dequeue() string {
	if q.Len() == 0 {
		return ""
	}
	s := (*q)[0]
	*q = (*q)[1:]
	return s
}

func (q *Queue) Sort() {
	sort.Strings(*q)
}

func (q Queue) IsEmpty() bool {
	return len(q) == 0
}
