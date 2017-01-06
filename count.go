package main

import (
	"bytes"
	"github.com/gonum/graph"
	. "github.com/gonum/graph/simple"
	"strconv"
)

// enumerate_subsets all subsets of the set {1 .. n}
// and puts them in the channel ch in the form of
// the elements of the subset as keys
// and the values as all false
func enumerate_subsets(n int, ch chan map[int]bool) {
	defer close(ch)
	if n == 0 {
		ch <- nil
	} else {
		ch_rec := make(chan map[int]bool)
		go enumerate_subsets(n-1, ch_rec)
		for subset := range ch_rec {
			ch <- subset
			subset2 := make(map[int]bool, len(subset)+1)
			for k := range subset {
				subset2[k] = false
			}
			subset2[n] = false
			ch <- subset2
		}
	}
}

// Takes an UndirectedGraph g and a list of nodes
// (Nodes in the form of map[int]bool where it is initially all false)
// And determines if they form a connected subgraph.
func are_connected(g *UndirectedGraph, nodes map[int]bool) bool {
	var recurse func(n graph.Node)
	recurse = func(n graph.Node) {
		if already_found, ok := nodes[n.ID()]; ok && !already_found {
			nodes[n.ID()] = true
			for _, neighbor := range g.From(n) {
				recurse(neighbor)
			}
		}
	}

	for nid := range nodes {
		recurse(Node(nid))
		break
	}
	for _, v := range nodes {
		if !v {
			return false
		}
	}
	return true
}

func connected_subset_filter(g *UndirectedGraph, in chan map[int]bool, out chan int) {
	defer close(out)
	for subset := range in {
		size := len(subset)
		switch {
		case size <= 1:
			out <- size
		case size == 2:
			var nodes [2]Node
			i := 0
			for k := range subset {
				nodes[i] = Node(k)
				i++
			}

			if g.HasEdgeBetween(nodes[0], nodes[1]) {
				out <- size
			}
		default:
			if are_connected(g, subset) {
				out <- size
			}
		}
	}
}

type CountSequence map[int]int

func (seq CountSequence) String() string {
	buf := new(bytes.Buffer)
	for i := 0; i < len(seq); i++ {
		if i == len(seq)-1 {
			buf.WriteString(strconv.Itoa(seq[i]))
		} else {
			buf.WriteString(strconv.Itoa(seq[i]))
			buf.WriteByte(',')
		}
	}
	return buf.String()
}

func count_subgraphs(g *UndirectedGraph) CountSequence {
	ch := make(chan map[int]bool)
	ch_out := make(chan int)
	go enumerate_subsets(len(g.Nodes()), ch)
	go connected_subset_filter(g, ch, ch_out)
	results := make(map[int]int)
	for res := range ch_out {
		results[res]++
	}
	return CountSequence(results)
}
