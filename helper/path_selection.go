package helper

import (
	"math"
)

type Group map[string][][]string



func getSourcesAndTargets(groups Group, key string)([][]string,[][]string){
	sources := groups[key]
	targets :=[][]string{}

	for k,v:=range groups{
		if k!=key{
			targets = append(targets, v...)
		}
	}
	return sources,targets
}

func GenerateAllPossibilities(thePromised []string, eligeables [][]string)[][][]string{
	groups:= MakeGroups(eligeables)
	allPossibilities :=[][][]string{}

	for key,_:=range groups{
		sources,targets:= getSourcesAndTargets(groups,key)
		allPossibilities= GenerateCombinations(thePromised,targets,sources)
		break
	}

	return allPossibilities
}


func GiveTheOneWithMostCandidates(paths[][]string,groups Group)[]string{
	thePromised:=[]string{}
	max:= math.MinInt32

	for _,path := range paths{
		eligeables:=GetEligeables(path,groups)
		_,flat:=Flat2DArray(eligeables)
		if len(flat)>max{
			max = len(flat)
			thePromised = path
		}
	}
	return thePromised
}

func GetEligeables(shortestOfAll []string,groups Group) [][]string{

	eligeables := [][]string{}

	for _,paths:= range groups{
			for _,path:=range paths{
				if !HasCommonElements(shortestOfAll[:len(shortestOfAll)-1],path[:len(path)-1]){
					eligeables = append(eligeables, path)
				}
			}
		
	}
	return eligeables
}

func GetSmallestPathOfEachGroup(groups Group) ([][]string, []string) {
	smallestPaths := [][]string{}
	var theSmallestOfAll []string
	lenSmallestOfAll := math.MaxInt32

	for _, paths := range groups {
		smallestPath := getSmallestPath(paths)
		_, flat := Flat2DArray(smallestPaths)
		if len(flat) < lenSmallestOfAll && len(flat) > 0 {
			lenSmallestOfAll = len(flat)
			theSmallestOfAll = smallestPath
		}
		smallestPaths = append(smallestPaths, smallestPath)
	}

	return smallestPaths, theSmallestOfAll
}

func getSmallestPath(paths [][]string) []string {
	var smallest []string
	lenSmall := math.MaxInt32

	for _, path := range paths {
		if len(path) < lenSmall && len(path) > 0 {
			lenSmall = len(path)
			smallest = path
		}
	}

	return smallest
}


func MakeGroups(paths[][]string)Group{
	groups:= Group{}

	for i,path:= range paths{
		if len(path)>0{
			groups[string(paths[i][0])]= append(groups[string(paths[i][0])], path)
		}
	}
	return groups
}