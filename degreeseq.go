package main

import (
	"bytes"
	"fmt"
	"github.com/gonum/graph/simple"
	"sort"
)

type DegreeSequence map[int]int

func (seq DegreeSequence) String() string {
	degs := make([]int, len(seq))
	i := 0
	for k := range seq {
		degs[i] = k
		i++
	}
	sort.Ints(degs)
	buf := new(bytes.Buffer)
	for i, deg := range degs {
		if i == len(degs)-1 {
			buf.WriteString(fmt.Sprintf("%d:%d", deg, seq[deg]))
		} else {
			buf.WriteString(fmt.Sprintf("%d:%d,", deg, seq[deg]))
		}
	}
	return buf.String()
}

func degree_sequence(g *simple.UndirectedGraph) DegreeSequence {
	results := make(DegreeSequence)
	for node := range g.Nodes() {
		deg := g.Degree(simple.Node(node))
		results[deg]++
	}
	return results
}
