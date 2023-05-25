package main

import (
	"fmt"
	"lemin/helper"
	"os"
)

func main() {

	args := os.Args[1:]

	if len(args)>0{
		filename := args[0]
		relations,rooms,ants :=helper.ParseInputFile(filename)
		table:= helper.DFS(relations,helper.PeekStartRoom(rooms).Name)
		// neeed to add another check to see if at least start and end are connected 
		// if all the nodes are not connected
		if len(table)!=len(rooms){
			fmt.Println("ERROR: invalid data format,a not all rooms are connected")
			os.Exit(0)
		}

		start:=helper.PeekStartRoom(rooms).Name
		end:=helper.PeekEndRoom(rooms).Name

		allPaths := helper.FindAllPaths(relations,start,end)


		// fmt.Println("All possible paths")
		// for _,path:=range allPaths{
			
		// 	length,_:=helper.GetPathLength(relations,path)
		// 	fmt.Println(length,path)

		// }
		// fmt.Println("-------------------")

		nonCrossing:= helper.FindNonCrossingPaths(allPaths)
 
		// fmt.Println("valid paths")
		// for _,valid:= range nonCrossing{

		// 	length,_:=helper.GetPathLength(relations,valid)
		// 	fmt.Println(length,valid)

		// }
		// fmt.Println("---------------------")
		helper.BigTraversal(nonCrossing,ants)
	}
}