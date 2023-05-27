package helper

import "math"

func BellmanFord(graph Relation, start string, end string) ([]string,int) {
	dist := make(map[string]int)
	prev := make(map[string]string)

	// Initialize distances with infinity
	for vertex := range graph {
		dist[vertex] = math.MaxInt32
	}
	dist[start] = 0

	// Relax edges repeatedly
	for i := 0; i < len(graph)-1; i++ {
		for vertex, edges := range graph {
			for _, edge := range edges {
				newDistance := dist[vertex] + edge.Distance
				if newDistance < dist[edge.Name] {
					dist[edge.Name] = newDistance
					prev[edge.Name] = vertex
				}
			}
		}
	}

	// Check for negative cycles
	for vertex, edges := range graph {
		for _, edge := range edges {
			newDistance := dist[vertex] + edge.Distance
			if newDistance < dist[edge.Name] {
				// Negative cycle detected
				return nil,math.MinInt32
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