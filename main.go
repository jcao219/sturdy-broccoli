package main

import (
	"fmt"
	"github.com/gonum/graph"
	. "github.com/gonum/graph/simple"
	"sync"
)

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

func peterson_graph() *UndirectedGraph {
	g := NewUndirectedGraph(0.0, 0.0)
	for i := 2; i <= 10; i++ {
		g.SetEdge(Edge{Node(i), Node(i - 1), 0.0})
	}
	g.SetEdge(Edge{Node(2), Node(10), 0.0})
	g.SetEdge(Edge{Node(3), Node(7), 0.0})
	g.SetEdge(Edge{Node(4), Node(9), 0.0})
	g.SetEdge(Edge{Node(6), Node(10), 0.0})
	g.SetEdge(Edge{Node(1), Node(5), 0.0})
	g.SetEdge(Edge{Node(1), Node(8), 0.0})
	return g
}

func are_connected(g *UndirectedGraph, nodes map[int]bool) bool {
	var mtx sync.Mutex
	var recurse func(graph.Node, *sync.WaitGroup)
	recurse = func(n graph.Node, wg *sync.WaitGroup) {
		defer wg.Done()
		mtx.Lock()
		already_found, ok := nodes[n.ID()]
		if ok {
			if already_found {
				mtx.Unlock()
				return
			}
			var wg_inner sync.WaitGroup
			nodes[n.ID()] = true
			mtx.Unlock()
			for _, neighbor := range g.From(n) {
				wg_inner.Add(1)
				go recurse(neighbor, &wg_inner)
			}
			wg_inner.Wait()
		} else {
			mtx.Unlock()
		}
	}

	var wg sync.WaitGroup
	for nid := range nodes {
		wg.Add(1)
		go recurse(Node(nid), &wg)
	}
	wg.Wait()
	for _, v := range nodes {
		if !v {
			return false
		}
	}
	return true
}

func main() {
	ch := make(chan map[int]bool)
	var g *UndirectedGraph = peterson_graph()
	go enumerate_subsets(len(g.Nodes()), ch)
	results := make(map[int]int)
	for subset := range ch {
		size := len(subset)
		switch {
		case size <= 1:
			results[size]++
		case size == 2:
			var nodes [2]Node
			i := 0
			for k := range subset {
				nodes[i] = Node(k)
				i++
			}

			if g.HasEdgeBetween(nodes[0], nodes[1]) {
				results[size]++
			}
		default:
			if are_connected(g, subset) {
				results[size]++
			}
		}
	}
	fmt.Println(results)
}
