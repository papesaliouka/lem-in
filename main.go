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
		relations,rooms,_ :=helper.ParseInputFile(filename)
		table:= helper.DFS(relations,helper.PeekStartRoom(rooms).Name)
		// neeed to add another check to see if at least start and end are connected 
		// if all the nodes are not connected
		// Need to check as well if there are cycle in the graph 
		// but don't know how to do that yet
		if len(table)==len(rooms){
			fmt.Println(table)
		}else{
			fmt.Println("ERROR: invalid data format,a not all rooms are connected")
		}

	}

}