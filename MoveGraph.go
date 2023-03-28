// This code can also be views via github at https://github.com/MAlshaik/Cops_and_Robbers

package main

import "fmt"
import "math"

func contains(slice []int, val int) bool {
    for _, item := range slice {
        if item == val {
            return true
        }
    }
    return false
}

func makeReflexive(graph map[int][]int) map[int][]int{
	/*
    In this code, the makeReflexive function takes a list containing lists
    of integers and makes every vertex connect to itself by an edge.
    */
    for vertex := range graph {
		graph[vertex] = append(graph[vertex], vertex)
	}

    for vertex, adjacentVertices := range graph {
        for _, adjacentVertex := range adjacentVertices {
            if !contains(graph[adjacentVertex], vertex) {
                // the edge (vertex, adjacentVertex) already exists in the graph
                graph[adjacentVertex] = []int{adjacentVertex}
            }         
        }
    }
	return graph
}


func makeMoveGraph(G map[int][]int) map[[3]int][][3]int {
    /*
    In this code, the makeMoveGraph function takes in a directed graph G
    represented as a map of slices of integers, and returns a new map of 
    slices moves that represents the possible moves on a 2D game board.
    The move graph has three integers representing the state of the board. 
    The arrays are set as such (x,y,67('C')) or (x,y,82('R')). 67 and 82 
    represent the cops and robbers turn respectively. x and y are the 
    position of the cops and robbers respectively.
    */
    moves := make(map[[3]int][][3]int)

    // Add self-loops to all vertices
    for x := range G {
        G[x] = append(G[x], x)
    }
      

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


func initLengthDictionary(moveGraph map[[3]int][][3]int) map[[3]int]float64 {
    /*
    The function initLengthMap takes a move graph and iterates through every move
    to calculate how many moves (or the length) left for the cop to win
    */
	lengths := make(map[[3]int]float64)

	for key := range moveGraph {
        // If cop position is equal to robber then set the length equal to 0 
		if key[0] == key[1] {
			lengths[key] = 0
        // if the robber position is not equal to the robber then set the length equal to positive infinity
		} else {
			lengths[key] = math.Inf(1)
		}
	}
	return lengths
}


func updateLengthDictionary(M map[[3]int][][3]int, L map[[3]int]float64) map[[3]int]float64 {
    // This function updates the length dictionary
	changesMade := true
    for changesMade {
        changesMade = false

        // iterate over each vertex in the graph.
        for key := range M {
            // check if the distance for the current vertex is infinity, indicating that it has not been visited.
            if L[key] == math.Inf(1) {
                // if the current vertex is a cop vertex.
                if key[2] == 'C' {
                    smallest := math.Inf(1)

                    // iterate over all neighbors of the current vertex.
                    for _, nbr := range M[key] {
                        smallest = math.Min(smallest, L[nbr])
                    }

                    // if the distance to the current vertex can be updated.
                    if smallest != L[key] {
                        changesMade = true
                    }

                    // update the distance to the current vertex.
                    L[key] = smallest + 1
                } else { // the current vertex is a robber vertex.
                    largest := 0.0
                    for _, nbr := range M[key] {
                        largest = math.Max(largest, L[nbr])
                    }
                    // if the distance to the current vertex can be updated.
                    if largest != L[key] {
                        changesMade = true
                    }
                    // update the distance to the current vertex.
                    L[key] = largest
                }
            }
        }
    }
    // return the updated length dictionary.
    return L

}


func checkCopWin(G map[int][]int) bool {
    M := makeMoveGraph(G)
    L := initLengthDictionary(M)
    L = updateLengthDictionary(M, L)
    copWin := true
    for key := range M {
        if L[key] == math.Inf(1) {
            // G is robber win
            copWin = false
            break
        }
    }
    return copWin
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

	graph = makeReflexive(graph)

	fmt.Println("After reflexion")
	for vertex, neighbors := range graph {
		fmt.Printf("Vertex %d has neighbors %v\n", vertex, neighbors)
	}
	//	After reflexion
	fmt.Println()

	// Creates a directed graph with three vertices and six directed edges
	diGraph := map[int][]int{
	    0: {1, 2},
        1: {2, 0},
        2: {0, 1},	
	}

	// Print the adjacency list of the di-graph
	for node, neighbors := range diGraph {
		fmt.Printf("%d: %v\n", node, neighbors)
	}

	fmt.Println()

    testGraph := make(map[int][]int)
    testGraph[1] = []int{1,2,3, 4}
    testGraph[2] = []int{2}
    testGraph[3] = []int{3,4}
    testGraph[4] = []int{4}

	moves := makeMoveGraph(graph)

	for state, neighbors := range moves {
		fmt.Printf("%v: %v\n", state, neighbors)
	}
     
    lengths := initLengthDictionary(moves)
    fmt.Println("Lengths:")
    for state, length := range lengths {
        fmt.Printf("%v: %.0f\n", state, length)
    }
    fmt.Println()

    updated_lengths := updateLengthDictionary(moves, lengths)
      
    fmt.Println("Updated Lengths:")
    for state, length := range updated_lengths {
        fmt.Printf("%v: %.0f\n", state, length)
    }
    fmt.Println()
    
    cop_win := checkCopWin(graph)
    fmt.Println("Did the cop win?")
    fmt.Println(cop_win)
    
}


/* The following is what the output looks like in the console:

Before reflexion
Vertex 1 has neighbors [2 3 4]
Vertex 3 has neighbors [4]

After reflexion
Vertex 1 has neighbors [2 3 4 1]
Vertex 3 has neighbors [4 3]

The adjacency list of the di-graph:
3: [0 1]
1: [1 2]
2: [2 0]

Move graph:
[1 3 67]: [[1 3 82] [2 3 82]]
[3 2 82]: [[3 2 67] [3 0 67]]
[2 3 82]: [[2 0 67] [2 1 67]]
[1 1 82]: [[1 1 67] [1 2 67]]
[1 2 67]: [[1 2 82] [2 2 82]]
[3 1 82]: [[3 1 67] [3 2 67]]
[2 1 67]: [[2 1 82] [0 1 82]]
[2 2 67]: [[2 2 82] [0 2 82]]
[3 3 67]: [[0 3 82] [1 3 82]]
[3 1 67]: [[0 1 82] [1 1 82]]
[1 3 82]: [[1 0 67] [1 1 67]]
[2 1 82]: [[2 1 67] [2 2 67]]
[1 2 82]: [[1 2 67] [1 0 67]]
[2 2 82]: [[2 2 67] [2 0 67]]
[2 3 67]: [[2 3 82] [0 3 82]]
[3 2 67]: [[0 2 82] [1 2 82]]
[3 3 82]: [[3 0 67] [3 1 67]]
[1 1 67]: [[1 1 82] [2 1 82]]

Lengths:
[1 2 82]: +Inf
[1 1 82]: 0
[2 2 67]: 0
[2 1 67]: +Inf
[3 1 67]: +Inf
[1 3 82]: +Inf
[1 3 67]: +Inf
[3 2 82]: +Inf
[3 1 82]: +Inf
[2 2 82]: 0
[2 3 67]: +Inf
[3 3 82]: 0
[2 3 82]: +Inf
[3 3 67]: 0
[1 2 67]: +Inf
[2 1 82]: +Inf
[1 1 67]: 0
[3 2 67]: +Inf

Updated Lengths:
[3 1 67]: 1
[1 3 82]: 0
[1 3 67]: 1
[3 2 82]: 1
[1 2 67]: 1
[3 1 82]: 1
[2 2 82]: 0
[2 3 67]: 1
[3 3 82]: 0
[2 3 82]: 1
[3 3 67]: 0
[2 1 82]: 1
[1 1 67]: 0
[3 2 67]: 1
[1 2 82]: 1
[1 1 82]: 0
[2 2 67]: 0
[2 1 67]: 1

Did the cop win?
true

*/
