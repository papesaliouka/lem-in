package helper

import "fmt"

type AntPaths map[int][]string

func BigTraversal(connections int,paths [][]string,ants int){
	canMove:=[]int{}
	antPaths := giveEachAntHisPath(ants,paths)
	updater:=connections
	maxMove:=0
	minMove:= len(paths[0])

	for _,v:=range paths{
		maxMove+=len(v)
	}

	for i:=1; i<=connections;i++{
		canMove = append(canMove, i)
	}

	fmt.Println(canMove)

	step:=1

	for updater <=ants{
		makeAStep(&canMove,ants,antPaths)
		canMove = append(canMove, updater)
		fmt.Println()
		updateCanMove(canMove[len(canMove)-1],step,minMove,maxMove,&canMove)
		updater++
		step++
	}

}

func updateCanMove(lastValue, step,maxMove, minMove int, canMove *[]int){
	if len(*canMove) <= maxMove{
		goal := lastValue + (minMove*step)
			for i:=lastValue+1; i<=goal;i++{
				*canMove= append(*canMove, i)
			}
		} 
}


func makeAmove(move string, antNumb int ){
	fmt.Printf("L%d-%s ",antNumb,move)
}

func makeAStep(canMove *[]int, ants int, antPaths AntPaths){
	for i,v:=range *canMove{
		pat := antPaths[v]
		if len(pat)>0{
			move := pat[0]
			makeAmove(move,v)
			pat= pat[1:]
			antPaths[v]=pat
		}else{
			if i<len(*canMove){
				*canMove = append((*canMove)[:i],(*canMove)[i+1:]...)
			}
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

