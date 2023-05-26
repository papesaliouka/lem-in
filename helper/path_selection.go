package helper

import "fmt"



type Group map[string][][]string




func MakeGroups(paths[][]string)Group{
	groups:= Group{}

	for i,path:= range paths{
		if len(path)>0{
			groups[string(paths[i][0])]= append(groups[string(paths[i][0])], path)
		}
	}
	return groups
}

func getSourceAndTargets(key string, groups Group)([][]string,[][]string){
	sources:= groups[key]
	targets := [][]string{}

	for k,v:=range groups{
		if k!=key{
			targets = append(targets, v...)
		}
	}
	return sources,targets
}

func GetEligeablePaths(groups Group, relations Relation) Group {
	eligeables := Group{}

	//bigArr:=[][]string{}

	for _,elements:= range groups{
		nonCrossing:=FindNonCrossingPaths(elements)
		fmt.Println(nonCrossing)
	}





	return eligeables
}



// Function to check if any element of slice1 is present in slice2
func containsAny(slice1 []string, slice2 []string) bool {
	set := make(map[string]bool)
	for _, element := range slice1 {
		set[element] = true
	}
	for _, element := range slice2 {
		if set[element] {
			return true
		}
	}
	return false
}

