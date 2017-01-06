package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/gonum/graph/simple"
	"os"
	"strconv"
	"strings"
)

func parse_graph(s string) (*simple.UndirectedGraph, error) {
	g := simple.NewUndirectedGraph(0.0, 0.0)
	for _, edge := range strings.Split(s, "  ") {
		pair := strings.Split(edge, " ")
		if len(pair) != 2 {
			return nil, errors.New("Bad format.")
		}
		v1, err1 := strconv.Atoi(pair[0])
		v2, err2 := strconv.Atoi(pair[1])
		if err1 != nil || err2 != nil {
			return nil, errors.New("Bad format.")
		}
		g.SetEdge(simple.Edge{F: simple.Node(v1), T: simple.Node(v2), W: 0.0})
	}
	return g, nil
}

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("One connected graph per line:")
		fmt.Println("1 2  1 4  1 3  1 5  1 6")
		fmt.Println("The above is a claw graph with 6 vertices, 5 edges.")
		fmt.Println("Notice the two spaces between edges.")
		fmt.Println("Node number must start with 1 and go up to N")
		fmt.Println("Need file with graphs in the above special format as argument.")
		return
	}
	file, err := os.Open(args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var g *simple.UndirectedGraph
		var err error
		g, err = parse_graph(scanner.Text())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(count_subgraphs(g))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
