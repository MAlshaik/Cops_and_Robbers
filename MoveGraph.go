package main

import "fmt"

func makeReflexive(graph map[int][]int) {
	for vertex := range graph {
		graph[vertex] = append(graph[vertex], vertex)
	}
}

func makeMoveGraph(G map[int][]int) map[[3]int][][3]int {
	moves := make(map[[3]int][][3]int)

	// Loop over each pair of vertices in G and their neighbors
	for x, _ := range G {
		for y, _ := range G {
			for _, z := range G[x] {
				// Add column move edge from (x,y,'C') to (z,y,'R')
				moves[[3]int{x, y, 'C'}] = append(moves[[3]int{x, y, 'C'}], [3]int{z, y, 'R'})
				// Add row move edge from (y,x,'R') to (y,z,'C')
				moves[[3]int{y, x, 'R'}] = append(moves[[3]int{y, x, 'R'}], [3]int{y, z, 'C'})
			}
		}
	}

	return moves
}

func main() {

	graph := make(map[int][]int)
	graph[1] = []int{2, 3, 4}
	graph[3] = []int{4}

	/* Builds a graph with 4 nodes numbered from 1 to 4 and declares the first node to
	have 2,3,4 as neigbors, 2 to have no neighbors, and 3 to have 4 as a neigbor.
	*/

	fmt.Println("Before reflexion")
	for vertex, neighbors := range graph {
		fmt.Printf("Vertex %d has neighbors %v\n", vertex, neighbors)
	}
	fmt.Println()
	// Roughly visualizes the graph before making it reflexive

	makeReflexive(graph)

	fmt.Println("After reflexion")
	for vertex, neighbors := range graph {
		fmt.Printf("Vertex %d has neighbors %v\n", vertex, neighbors)
	}
	//	After reflexion

	fmt.Println()

	// Creates a directed graph with three vertices and six directed edges
	diGraph := map[int][]int{
		1: {1, 2},
		2: {2, 0},
		3: {0, 1},
	}

	// Print the adjacency list of the di-graph
	for node, neighbors := range diGraph {
		fmt.Printf("%d: %v\n", node, neighbors)
	}

	fmt.Println()

	moves := makeMoveGraph(graph)

	for state, neighbors := range moves {
		fmt.Printf("%v: %v\n", state, neighbors)
	}

}
