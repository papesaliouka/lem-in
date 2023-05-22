package helper

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Graph map[string][]string



func ParseInputFile(filename string) (Graph, int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	graph := make(Graph)

	var numAnts int
	var numRooms int

	for scanner.Scan() {
		line := scanner.Text()
		if line == "##start" || strings.ContainsAny(line,"#") || line == "##end" {
			continue
		}

		if numAnts == 0 {
			numAnts, err = strconv.Atoi(line)
			if err != nil {
				fmt.Println("ERROR: invalid data format")
				return nil, 0, nil
			}
		} else if strings.Contains(line, "-") {
			parts := strings.Split(line, "-")
			from := parts[0]
			to := parts[1]
			graph.AddArc(from, to)
		} else {
			parts := strings.Split(line, " ")
			if len(parts) != 3 {
				// Line does not have three elements

				fmt.Println("ERROR: invalid data format")
				return nil, 0, nil
			}
			name := parts[0]
			first := parts[1]
			second := parts[2]

			_, err := strconv.Atoi(first)
			if err != nil {
				// not an integer
				fmt.Println("ERROR: invalid data format")
				return nil, 0, nil
			}
			_, err = strconv.Atoi(second)
			if err != nil {
				// not an integer
				fmt.Println("ERROR: invalid data format")
				return nil, 0, nil
			}
			graph.AddRoom(name)
			numRooms++
		}

	}

	err = scanner.Err()
	if err != nil {
		return nil, 0, err
	}

	return graph, numRooms, nil
}

func (g Graph) AddRoom(name string) {
	if _, exists := g[name]; !exists {
		g[name] = []string{}
	}
}

func (g Graph) AddArc(from, to string) {
	g[from] = append(g[from], to)
	g[to] = append(g[to], from)
}

func DFS(graph Graph, startRoom string) []string {
	visited := make(map[string]bool)
	dfsTraversal := make([]string, 0)

	dfsRecursive(graph, startRoom, visited, &dfsTraversal)

	return dfsTraversal
}

func dfsRecursive(graph Graph, room string, visited map[string]bool, traversal *[]string) {
	visited[room] = true
	*traversal = append(*traversal, room)

	neighbors := graph[room]
	for _, neighbor := range neighbors {
		if !visited[neighbor] {
			dfsRecursive(graph, neighbor, visited, traversal)
		}
	}
}
