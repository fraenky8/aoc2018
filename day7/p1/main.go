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

type Stack []string

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Thanks to my brilliant girlfriend Sarah <3

	graph := NewGraph(file)
	//graph := NewGraph(strings.NewReader(example))

	startNodes := GetStartNodes(graph)

	fmt.Println("order of instructions:")
	for i, start := range startNodes {
		fmt.Printf("\tsolution %d: %v\n", i+1, FindSolution(graph, start))
		ResetGraph(graph)
	}
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

// FindSolution is basically a DFS
func FindSolution(graph Nodes, start *Node) string {
	stack := &Stack{start.Letter}

	var solution strings.Builder

	for !stack.IsEmpty() {

		current := stack.Pop()
		solution.WriteString(current)

		graph[current].Visited = true

		enabledNodes := GetAllEnabledNodes(graph)
		for _, node := range enabledNodes {
			stack.Push(node)
		}
		stack.Sort()
	}

	return solution.String()
}

func ResetGraph(nodes Nodes) {
	for _, node := range nodes {
		node.Visited = false
	}
}

func GetStartNodes(nodes Nodes) []*Node {
	var sn []*Node
	for _, node := range nodes {
		if len(node.Requires) == 0 {
			sn = append(sn, node)
		}
	}
	return sn
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

func (s *Stack) Push(node *Node) {
	for _, e := range *s {
		if e == node.Letter {
			return
		}
	}
	*s = append(*s, node.Letter)
}

func (s Stack) Len() int {
	return len(s)
}

func (s *Stack) Pop() string {
	if s.Len() == 0 {
		return ""
	}
	e := (*s)[0]
	*s = (*s)[1:]
	return e
}

func (s *Stack) Sort() {
	sort.Strings(*s)
}

func (s Stack) IsEmpty() bool {
	return len(s) == 0
}
