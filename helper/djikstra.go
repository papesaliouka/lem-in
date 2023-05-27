package helper

import "math"

func Dijkstra(graph Relation, start string, end string) ([]string,int) {
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
	for curr != start {
		path = append([]string{curr}, path...)
		curr = prev[curr]
	}
	path = append([]string{start}, path...)

	return path,dist[end]
}


func minDistance(dist map[string]int, visited map[string]bool) string {
	min := math.MaxInt32
	var minVertex string

	for vertex, distance := range dist {
		if !visited[vertex] && distance < min {
			min = distance
			minVertex = vertex
		}
	}

	return minVertex
}
