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

    graph := map[int][]int{
        1: {2, 3, 4},
        2: {3, 5, 7},
        3: {4, 5, 6},
        4: {6, 7},
        5: {6, 7},
        6: {7},
    }

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

    one_pitfall := map[int][]int{
        1: {2, 3, 4},
        2: {3, 5, 7},
        3: {4, 5, 6},
        4: {6, 7},
        5: {6, 7},
        6: {7},
    }
    
    cop_win := checkCopWin(one_pitfall)
    fmt.Println("Did the cop win? (one pitfall graph)")
    fmt.Println(cop_win)
    
    example_graph := map[int][]int{
        1: {2,3,4},
        3: {4},
    }

    cop_win = checkCopWin(example_graph)
    fmt.Println("Did the cop win? (cocalc example graph)")
    fmt.Println(cop_win)
}


/*
Before reflexion
Vertex 4 has neighbors [6 7]
Vertex 5 has neighbors [6 7]
Vertex 6 has neighbors [7]
Vertex 1 has neighbors [2 3 4]
Vertex 2 has neighbors [3 5 7]
Vertex 3 has neighbors [4 5 6]

After reflexion
Vertex 6 has neighbors [6]
Vertex 7 has neighbors [7]
Vertex 1 has neighbors [2 3 4 1]
Vertex 2 has neighbors [2]
Vertex 3 has neighbors [3]
Vertex 4 has neighbors [4]
Vertex 5 has neighbors [6 7 5]

0: [1 2]
1: [2 0]
2: [0 1]

[5 2 67]: [[6 2 82] [7 2 82] [5 2 82] [5 2 82]]
[2 5 82]: [[2 6 67] [2 7 67] [2 5 67] [2 5 67]]
[4 6 67]: [[4 6 82] [4 6 82]]
[4 7 67]: [[4 7 82] [4 7 82]]
[4 4 67]: [[4 4 82] [4 4 82]]
[5 4 67]: [[6 4 82] [7 4 82] [5 4 82] [5 4 82]]
[5 1 67]: [[6 1 82] [7 1 82] [5 1 82] [5 1 82]]
[2 4 67]: [[2 4 82] [2 4 82]]
[3 3 67]: [[3 3 82] [3 3 82]]
[4 6 82]: [[4 6 67] [4 6 67]]
[1 4 67]: [[2 4 82] [3 4 82] [4 4 82] [1 4 82] [1 4 82]]
[3 6 82]: [[3 6 67] [3 6 67]]
[7 5 67]: [[7 5 82] [7 5 82]]
[7 6 67]: [[7 6 82] [7 6 82]]
[7 2 82]: [[7 2 67] [7 2 67]]
[2 1 67]: [[2 1 82] [2 1 82]]
[4 1 67]: [[4 1 82] [4 1 82]]
[5 5 82]: [[5 6 67] [5 7 67] [5 5 67] [5 5 67]]
[5 7 67]: [[6 7 82] [7 7 82] [5 7 82] [5 7 82]]
[6 7 82]: [[6 7 67] [6 7 67]]
[4 1 82]: [[4 2 67] [4 3 67] [4 4 67] [4 1 67] [4 1 67]]
[2 3 67]: [[2 3 82] [2 3 82]]
[5 3 82]: [[5 3 67] [5 3 67]]
[6 3 82]: [[6 3 67] [6 3 67]]
[7 7 82]: [[7 7 67] [7 7 67]]
[2 1 82]: [[2 2 67] [2 3 67] [2 4 67] [2 1 67] [2 1 67]]
[7 5 82]: [[7 6 67] [7 7 67] [7 5 67] [7 5 67]]
[7 2 67]: [[7 2 82] [7 2 82]]
[1 3 67]: [[2 3 82] [3 3 82] [4 3 82] [1 3 82] [1 3 82]]
[2 7 67]: [[2 7 82] [2 7 82]]
[1 4 82]: [[1 4 67] [1 4 67]]
[2 4 82]: [[2 4 67] [2 4 67]]
[5 5 67]: [[6 5 82] [7 5 82] [5 5 82] [5 5 82]]
[5 6 67]: [[6 6 82] [7 6 82] [5 6 82] [5 6 82]]
[7 3 82]: [[7 3 67] [7 3 67]]
[4 3 82]: [[4 3 67] [4 3 67]]
[4 2 67]: [[4 2 82] [4 2 82]]
[2 7 82]: [[2 7 67] [2 7 67]]
[3 2 82]: [[3 2 67] [3 2 67]]
[1 2 82]: [[1 2 67] [1 2 67]]
[2 6 82]: [[2 6 67] [2 6 67]]
[7 7 67]: [[7 7 82] [7 7 82]]
[1 1 67]: [[2 1 82] [3 1 82] [4 1 82] [1 1 82] [1 1 82]]
[4 2 82]: [[4 2 67] [4 2 67]]
[6 7 67]: [[6 7 82] [6 7 82]]
[1 6 82]: [[1 6 67] [1 6 67]]
[5 7 82]: [[5 7 67] [5 7 67]]
[1 1 82]: [[1 2 67] [1 3 67] [1 4 67] [1 1 67] [1 1 67]]
[3 2 67]: [[3 2 82] [3 2 82]]
[6 6 67]: [[6 6 82] [6 6 82]]
[1 2 67]: [[2 2 82] [3 2 82] [4 2 82] [1 2 82] [1 2 82]]
[4 7 82]: [[4 7 67] [4 7 67]]
[7 1 82]: [[7 2 67] [7 3 67] [7 4 67] [7 1 67] [7 1 67]]
[3 5 67]: [[3 5 82] [3 5 82]]
[4 5 67]: [[4 5 82] [4 5 82]]
[7 4 82]: [[7 4 67] [7 4 67]]
[1 5 82]: [[1 6 67] [1 7 67] [1 5 67] [1 5 67]]
[6 6 82]: [[6 6 67] [6 6 67]]
[2 2 67]: [[2 2 82] [2 2 82]]
[1 6 67]: [[2 6 82] [3 6 82] [4 6 82] [1 6 82] [1 6 82]]
[6 1 82]: [[6 2 67] [6 3 67] [6 4 67] [6 1 67] [6 1 67]]
[7 3 67]: [[7 3 82] [7 3 82]]
[3 7 82]: [[3 7 67] [3 7 67]]
[1 5 67]: [[2 5 82] [3 5 82] [4 5 82] [1 5 82] [1 5 82]]
[2 6 67]: [[2 6 82] [2 6 82]]
[4 3 67]: [[4 3 82] [4 3 82]]
[4 5 82]: [[4 6 67] [4 7 67] [4 5 67] [4 5 67]]
[3 5 82]: [[3 6 67] [3 7 67] [3 5 67] [3 5 67]]
[6 1 67]: [[6 1 82] [6 1 82]]
[2 2 82]: [[2 2 67] [2 2 67]]
[7 1 67]: [[7 1 82] [7 1 82]]
[7 4 67]: [[7 4 82] [7 4 82]]
[1 7 67]: [[2 7 82] [3 7 82] [4 7 82] [1 7 82] [1 7 82]]
[6 3 67]: [[6 3 82] [6 3 82]]
[5 6 82]: [[5 6 67] [5 6 67]]
[6 2 67]: [[6 2 82] [6 2 82]]
[2 3 82]: [[2 3 67] [2 3 67]]
[5 4 82]: [[5 4 67] [5 4 67]]
[4 4 82]: [[4 4 67] [4 4 67]]
[6 5 82]: [[6 6 67] [6 7 67] [6 5 67] [6 5 67]]
[5 3 67]: [[6 3 82] [7 3 82] [5 3 82] [5 3 82]]
[1 7 82]: [[1 7 67] [1 7 67]]
[5 1 82]: [[5 2 67] [5 3 67] [5 4 67] [5 1 67] [5 1 67]]
[3 4 82]: [[3 4 67] [3 4 67]]
[3 1 82]: [[3 2 67] [3 3 67] [3 4 67] [3 1 67] [3 1 67]]
[2 5 67]: [[2 5 82] [2 5 82]]
[5 2 82]: [[5 2 67] [5 2 67]]
[3 7 67]: [[3 7 82] [3 7 82]]
[1 3 82]: [[1 3 67] [1 3 67]]
[6 4 82]: [[6 4 67] [6 4 67]]
[6 4 67]: [[6 4 82] [6 4 82]]
[6 5 67]: [[6 5 82] [6 5 82]]
[7 6 82]: [[7 6 67] [7 6 67]]
[6 2 82]: [[6 2 67] [6 2 67]]
[3 1 67]: [[3 1 82] [3 1 82]]
[3 3 82]: [[3 3 67] [3 3 67]]
[3 4 67]: [[3 4 82] [3 4 82]]
[3 6 67]: [[3 6 82] [3 6 82]]
Lengths:
[2 3 67]: +Inf
[3 7 82]: +Inf
[6 1 67]: +Inf
[3 4 67]: +Inf
[5 4 67]: +Inf
[5 1 82]: +Inf
[3 7 67]: +Inf
[6 4 82]: +Inf
[5 5 82]: 0
[7 4 82]: +Inf
[4 3 67]: +Inf
[2 2 82]: 0
[5 3 67]: +Inf
[5 5 67]: 0
[4 4 67]: 0
[5 1 67]: +Inf
[7 2 82]: +Inf
[1 2 82]: +Inf
[6 6 67]: 0
[1 3 82]: +Inf
[4 7 67]: +Inf
[6 5 67]: +Inf
[3 3 82]: 0
[3 2 82]: +Inf
[4 7 82]: +Inf
[6 1 82]: +Inf
[6 5 82]: +Inf
[1 4 82]: +Inf
[2 4 82]: +Inf
[6 6 82]: 0
[2 3 82]: +Inf
[5 2 82]: +Inf
[2 7 67]: +Inf
[3 1 82]: +Inf
[5 3 82]: +Inf
[2 1 82]: +Inf
[4 3 82]: +Inf
[2 7 82]: +Inf
[4 5 67]: +Inf
[4 5 82]: +Inf
[3 5 82]: +Inf
[2 5 67]: +Inf
[6 2 82]: +Inf
[3 1 67]: +Inf
[4 2 82]: +Inf
[3 5 67]: +Inf
[1 5 82]: +Inf
[1 5 67]: +Inf
[2 5 82]: +Inf
[2 1 67]: +Inf
[1 3 67]: +Inf
[3 3 67]: 0
[1 4 67]: +Inf
[7 6 67]: +Inf
[2 2 67]: 0
[1 6 67]: +Inf
[1 7 67]: +Inf
[2 4 67]: +Inf
[4 1 82]: +Inf
[7 1 82]: +Inf
[2 6 67]: +Inf
[7 3 67]: +Inf
[1 7 82]: +Inf
[5 7 67]: +Inf
[3 6 82]: +Inf
[6 3 82]: +Inf
[6 7 67]: +Inf
[3 2 67]: +Inf
[7 1 67]: +Inf
[4 1 67]: +Inf
[7 2 67]: +Inf
[5 6 67]: +Inf
[5 7 82]: +Inf
[1 2 67]: +Inf
[4 6 82]: +Inf
[6 7 82]: +Inf
[5 4 82]: +Inf
[3 4 82]: +Inf
[3 6 67]: +Inf
[6 4 67]: +Inf
[4 6 67]: +Inf
[5 2 67]: +Inf
[7 3 82]: +Inf
[1 1 67]: 0
[7 4 67]: +Inf
[5 6 82]: +Inf
[7 6 82]: +Inf
[7 5 67]: +Inf
[7 7 82]: 0
[7 5 82]: +Inf
[1 6 82]: +Inf
[1 1 82]: 0
[4 4 82]: 0
[6 3 67]: +Inf
[4 2 67]: +Inf
[2 6 82]: +Inf
[7 7 67]: 0
[6 2 67]: +Inf

Updated Lengths:
[4 2 67]: +Inf
[2 6 82]: +Inf
[7 7 67]: 0
[6 2 67]: +Inf
[2 3 67]: +Inf
[3 7 82]: +Inf
[6 1 67]: +Inf
[3 4 67]: +Inf
[5 4 67]: +Inf
[5 1 82]: +Inf
[3 7 67]: +Inf
[6 4 82]: +Inf
[5 5 82]: 0
[7 4 82]: +Inf
[4 3 67]: +Inf
[2 2 82]: 0
[5 3 67]: +Inf
[5 5 67]: 0
[4 4 67]: 0
[5 1 67]: +Inf
[7 2 82]: +Inf
[1 2 82]: 1
[6 6 67]: 0
[1 3 82]: 1
[4 7 67]: +Inf
[6 5 67]: +Inf
[3 3 82]: 0
[3 2 82]: +Inf
[4 7 82]: +Inf
[6 1 82]: +Inf
[6 5 82]: +Inf
[1 4 82]: 1
[2 4 82]: +Inf
[6 6 82]: 0
[2 3 82]: +Inf
[5 2 82]: +Inf
[2 7 67]: +Inf
[3 1 82]: +Inf
[5 3 82]: +Inf
[2 1 82]: +Inf
[4 3 82]: +Inf
[2 7 82]: +Inf
[4 5 67]: +Inf
[4 5 82]: +Inf
[3 5 82]: +Inf
[2 5 67]: +Inf
[6 2 82]: +Inf
[3 1 67]: +Inf
[4 2 82]: +Inf
[3 5 67]: +Inf
[1 5 82]: +Inf
[1 5 67]: +Inf
[2 5 82]: +Inf
[2 1 67]: +Inf
[1 3 67]: 1
[3 3 67]: 0
[1 4 67]: 1
[7 6 67]: +Inf
[2 2 67]: 0
[1 6 67]: +Inf
[1 7 67]: +Inf
[2 4 67]: +Inf
[4 1 82]: +Inf
[7 1 82]: +Inf
[2 6 67]: +Inf
[7 3 67]: +Inf
[1 7 82]: +Inf
[5 7 67]: 1
[3 6 82]: +Inf
[6 3 82]: +Inf
[6 7 67]: +Inf
[3 2 67]: +Inf
[7 1 67]: +Inf
[4 1 67]: +Inf
[7 2 67]: +Inf
[5 6 67]: 1
[5 7 82]: 1
[1 2 67]: 1
[4 6 82]: +Inf
[6 7 82]: +Inf
[5 4 82]: +Inf
[3 4 82]: +Inf
[3 6 67]: +Inf
[6 4 67]: +Inf
[4 6 67]: +Inf
[5 2 67]: +Inf
[7 3 82]: +Inf
[1 1 67]: 0
[7 4 67]: +Inf
[5 6 82]: 1
[7 6 82]: +Inf
[7 5 67]: +Inf
[7 7 82]: 0
[7 5 82]: +Inf
[1 6 82]: +Inf
[1 1 82]: 0
[4 4 82]: 0
[6 3 67]: +Inf

Did the cop win? (one pitfall graph)
true
Did the cop win? (cocalc example graph)
true
*/
