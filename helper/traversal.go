package helper

import "fmt"

type AntPaths map[int][]string

func BigTraversal(connections int,paths [][]string,ants int){
	pathIndex:=0
	canMove:=[]int{}
	maxMovePerTurn:=0

	for i:=1; i<=connections;i++{
		canMove = append(canMove, i)
	}

	for j:=0;j<len(paths);j++{
		maxMovePerTurn += len(paths[j])
	}

	fmt.Println(maxMovePerTurn)

	for i:=1;i<=ants;i++{
		pathIndex = (pathIndex +1) % len(paths) 
	}

	antPaths := giveEachAntHisPath(ants,paths)


	updater:=connections

	for updater <=ants{
		updater++
		canMove = append(canMove, updater)
	}

 	i:=0
	for {
		makeAStep(canMove,ants,antPaths)
		fmt.Println()
		if i==8 {
			break
		}
		i++
	}






}


func makeAmove(move string, antNumb int ){
	fmt.Printf("L%d-%s ",antNumb,move)
}

func makeAStep(canMove []int, ants int, antPaths AntPaths){
	for _,v:=range canMove{
		pat := antPaths[v]
		if len(pat)>0{
			move := pat[0]
			makeAmove(move,v)
			pat= pat[1:]
			antPaths[v]=pat
		}
	}
}

func giveEachAntHisPath(ants int, paths [][]string) AntPaths {
	antPaths := AntPaths{}
	pathIndex:=0
	for i:= 1; i<=ants ;i++{
		pathIndex = (pathIndex +1) % len(paths)
		antPaths[i]=paths[pathIndex]
	}
	return antPaths
}

