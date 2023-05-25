package helper

import (
	"fmt"
)

type AntPaths map[int][]string

func ValidateStartingConnections(startConnections []RelationAndDistance, nonCrossing [][]string) int {

	connections :=0
	starts := []string{}

	nonCrossingStarter:=[]string{}

	for _,v:=range startConnections{
		starts = append(starts, v.Name)
	}

	for _,v:=range nonCrossing{
		nonCrossingStarter = append(nonCrossingStarter, v[0])
	}

	for _,v:=range starts{
		if contains(v,nonCrossingStarter){
			connections++
		}
	}

return connections
}

func BigTraversal(connections int,paths [][]string,ants int){
	canMove:=[]int{}
	antPaths := giveEachAntHisPath(ants,paths)
	maxMove:=0
	minMove:= connections

	for _,v:=range paths{
		maxMove+=len(v)
	}

	for i:=1; i<=connections;i++{
		canMove = append(canMove, i)
	}

	step:=1
	lastIndex := len(canMove)-1
	last:= canMove[lastIndex]

	for len(canMove)>0{		
		last = makeAStep(&canMove,ants,&antPaths,last, step,maxMove,minMove)
		step++
	}
}

func updateCanMove(lastValue, step,maxMove, minMove ,ants int ,canMove *[]int)int {
//	if len(*canMove) <= maxMove  {
		goal := lastValue + (minMove * step)
		for i:=lastValue+1; i<=goal;i++{
			if !containsAnt(i, canMove) && i<=ants {
				*canMove= append(*canMove, i)
				lastValue = i
			}
		}
	//}
	return lastValue 
}

func makeAStep(canMove *[]int, ants int, antPaths *AntPaths,lastValue,step,maxMove,minMove int) int{
	for _,v:=range *canMove{
		pat := (*antPaths)[v]
		if len(pat)>0{
			move := pat[0]
			makeAmove(move,v)
			pat= pat[1:]
			(*antPaths)[v]=pat
		}
	}

	for i,v:=range *canMove{
		pat := (*antPaths)[v]
		if len(pat)==0{
			if i<len(*canMove){
			 	*canMove = append((*canMove)[:i],(*canMove)[i+1:]...)
			}
		}
	}

	if len(*canMove)>0{
		fmt.Println()
		lastValue=updateCanMove(lastValue,step,maxMove,minMove,ants,canMove)
	}

	return lastValue
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

func containsAnt(ant int, ants *[]int)bool{
	for _,v:= range *ants{
		if v==ant{
			return true
		}
	}
	return false
}

func contains(ant string, ants []string)bool{
	for _,v:= range ants{
		if v ==ant{
			return true
		}
	}
	return false
}

func makeAmove(move string, antNumb int ){
	fmt.Printf("L%d-%s ",antNumb,move)
}

