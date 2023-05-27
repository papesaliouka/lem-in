package helper

import (
	"fmt"
	"os"
)



func Flat2DArray(v [][]string) ([][]string, []string) {
	flatened := [][]string{}
	flatened = append(flatened, v...)
	flat := []string{}
	for _, val := range flatened {
		flat = append(flat, val...)
	}
	return flatened, flat
}

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

func RemoveStart(paths[][]string)[][]string{

	trimmed:=[][]string{}

	for _,v:= range paths{
		if len(v)>0 {
			lastIndex:= len(v)-1
			if lastIndex>0{
				path:= v[1:]
				path = path[:lastIndex-1]
				trimmed = append(trimmed, path)
			}
		}
	}

	return trimmed

}

// Function to check if two string arrays have common elements
func HasCommonElements(arr1, arr2 []string) bool {
	set := make(map[string]bool)

	// Add elements of arr1 to the set
	for _, element := range arr1 {
		set[element] = true
	}

	// Check if elements of arr2 are already present in the set
	for _, element := range arr2 {
		if set[element] {
			return true
		}
	}

	return false
}


func HasCommonElements2(arr [][]string) bool {
	set := make(map[string]bool)

	// Iterate over each inner array
	for _, subArr := range arr {
		// Check if any element in the current inner array is already present in the set
		for _, element := range subArr {
			if set[element] {
				return true
			}
			set[element] = true
		}
	}

	return false
}



func DFS(graph Relation, startRoom string) []string {
	visited := make(map[string]bool)
	dfsTraversal := make([]string, 0)

	dfsRecursive(graph, startRoom, visited, &dfsTraversal)

	return dfsTraversal
}

func dfsRecursive(graph Relation, room string, visited map[string]bool, traversal *[]string,) {
	visited[room] = true
	*traversal = append(*traversal, room)
	neighbors := graph[room]
	for _, neighbor := range neighbors {
		if !visited[neighbor.Name] {
			dfsRecursive(graph, neighbor.Name, visited, traversal)
		}
	}
}
