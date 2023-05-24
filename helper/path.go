package helper

import (
	"fmt"
	"strings"
)


func ValidatePaths(paths [][]string) [][]string {
	validPaths:= [][]string{}
	store := map[string][]string{}
	listStr:=[]string{}

	groups:= map[string][]string{}

	for _,path:= range paths{
		pathCopy:=path
		pathStr := strings.Join(path[1:len(path)-1], "")
		listStr = append(listStr,pathStr)
		store[pathStr]=pathCopy
	}

	for _,v:=range listStr{
		groups[string(v[0])]= append(groups[string(v[0])], v)
	}

	candidatesKeys:=[]string{}

	for k := range groups {targets,sources := GetTargetsAndSources(groups,k)
		for _,word:=range sources{
			key,isValid:=IsValidKeys(word,targets)
			if isValid{
				candidatesKeys = append(candidatesKeys, key)
			}
		}
	}

	finalGroup:=map[string][]string{}

	for i,cand := range candidatesKeys{
		finalGroup[string(candidatesKeys[i][0])]=append(finalGroup[string(candidatesKeys[i][0])], cand)
	}

	fmt.Println(finalGroup)

	return validPaths
}


func finalValidation(candidates []string)[]string{
	selected:=[]string{}

	for _,cand := range candidates{
		for _,cand2:= range candidates{
			if cand2 !=cand{
				if !contains(cand,cand2){
					selected = append(selected, cand2)
				}
			}
		}
	}
	fmt.Println(selected)
	return selected
}

func GetTargetsAndSources(group map[string][]string ,key string)([]string,[]string){
	targets := []string{}
	sources:=group[key]

	for k,arr:= range group{
		if k !=key{
			targets = append(targets, arr...)
		}
	}
	return targets,sources
}

func IsValidKeys(src string , targets []string ) (string,bool){
	valid :=[]string{}
	for _,destWord :=range targets{
		if contains(src,destWord){
			continue
		}else{
			valid = append(valid, src)
		}
	}

	if len(valid)>0{
		return valid[0],true
	}

	return "",false
	
}

func contains(src ,target string)bool{
	for _,v:=range src{
		if strings.Contains(target,string(v)){
			return true
		}
	}
	return false
}



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
