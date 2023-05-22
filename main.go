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

		rooms := helper.GetRooms(filename)	
		relations:= helper.GetRelations(filename,rooms)
		fmt.Println(relations)
	}

}