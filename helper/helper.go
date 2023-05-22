package helper

import (
	"fmt"
	"os"
)

func ParseInputFile(filename string) (Relation,[]Room, int) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)		
	}
	defer file.Close()

	antNumbers :=GetAntsNumber(filename) 
	rooms := GetRooms(filename)
	relations :=GetRelations(filename,rooms)

	return relations,rooms,antNumbers
}


func DFS(graph Relation, startRoom string) []string {
	visited := make(map[string]bool)
	dfsTraversal := make([]string, 0)

	dfsRecursive(graph, startRoom, visited, &dfsTraversal)

	return dfsTraversal
}

func dfsRecursive(graph Relation, room string, visited map[string]bool, traversal *[]string) {
	visited[room] = true
	*traversal = append(*traversal, room)

	neighbors := graph[room]
	for _, neighbor := range neighbors {
		if !visited[neighbor.Neighbor] {
			dfsRecursive(graph, neighbor.Neighbor, visited, traversal)
		}
	}
}
