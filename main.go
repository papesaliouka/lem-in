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
		table := helper.DFS(relations,helper.PeekStartRoom(rooms).Name)
		fmt.Println(table)
	}

}