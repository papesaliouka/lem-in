package helper

import (
	"math"
	"fmt"
)



func GetPathLength(adjList Relation, path []string) (int, error) {
	length := 0

	// Iterate over the nodes in the path
	for i := 0; i < len(path)-1; i++ {
		currNode := path[i]
		nextNode := path[i+1]

		// Check if an edge exists between the current and next node
		edges, ok := adjList[currNode]
		if !ok {
			return 0, fmt.Errorf("invalid path: node %s is not present in the graph", currNode)
		}

		edgeExists := false
		for _, edge := range edges {
			if edge.Name == nextNode {
				length += edge.Distance
				edgeExists = true
				break
			}
		}

		if !edgeExists {
			return 0, fmt.Errorf("invalid path: no edge found between nodes %s and %s", currNode, nextNode)
		}
	}

	return length, nil
}


func FindAllPaths(adjList Relation, start, end string) [][]string {
	visited := make(map[string]bool)
	path := []string{start}
	paths := [][]string{}

	dfs2(adjList, start, end, visited, path, &paths)

	return paths
}

func dfs2(adjList Relation, current, end string, visited map[string]bool, path []string, paths *[][]string) {
	visited[current] = true

	if current == end {
		// Append a copy of the current path to the paths slice
		*paths = append(*paths, append([]string(nil), path...))
	} else {
		edges := adjList[current]
		for _, edge := range edges {
			if !visited[edge.Name] {
				// Explore the unvisited neighbor
				dfs2(adjList, edge.Name, end, visited, append(path, edge.Name), paths)
			}
		}
	}

	visited[current] = false // Mark the current node as unvisited for other paths
}



func removeEdge(adjList Relation, from, to string) {
	// Remove the edge from the source node's adjacency list
	edges := adjList[from]
	for i, edge := range edges {
		if edge.Name == to {
			adjList[from] = append(edges[:i], edges[i+1:]...)
			break
		}
	}

	// Remove the edge from the target node's adjacency list
	edges = adjList[to]
	for i, edge := range edges {
		if edge.Name == from {
			adjList[to] = append(edges[:i], edges[i+1:]...)
			break
		}
	}
}




func GetShortestPaths(graph Relation, start, end string) [][]string{
	paths:=[][]string{}

	connections := len(graph[start])

	for i:=0; i<connections;i++{
		path := Dijkstra(graph,start,end)
		paths = append(paths, path[1:])

		for i,v:=range path{
			if v !=start && v!=end{
				removeEdge(graph, path[i], path[i+1])
			}
		}

	}

	return paths
}


func Dijkstra(graph Relation, start string, end string) []string {
	dist := make(map[string]int)
	visited := make(map[string]bool)
	prev := make(map[string]string)

	// Initialize distances with infinity
	for vertex := range graph {
		dist[vertex] = math.MaxInt32
	}
	dist[start] = 0

	for i := 0; i < len(graph); i++ {
		current := minDistance(dist, visited)
		visited[current] = true

		// Stop early if we reach the destination vertex
		if current == end {
			break
		}

		// Update distances of neighboring vertices
		for _, edge := range graph[current] {
			if !visited[edge.Name] {
				newDistance := dist[current] + edge.Distance
				if newDistance < dist[edge.Name] {
					dist[edge.Name] = newDistance
					prev[edge.Name] = current
				}
			}
		}
	}

	// Reconstruct the shortest path
	path := []string{}
	curr := end
	for curr != start && curr!="" {
		path = append([]string{curr}, path...)
		curr = prev[curr]
	}
	path = append([]string{start}, path...)

	return  path
}

func minDistance(dist map[string]int, visited map[string]bool) string {
	min := math.MaxInt32
	minVertex := ""

	for vertex, distance := range dist {
		if !visited[vertex] && distance < min {
			min = distance
			minVertex = vertex
		}
	}

	return minVertex
}

