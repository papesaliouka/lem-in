package main

import (
	"fmt"
	"lemin/helper"
)

func main() {
	graph, _, err := helper.ParseInputFile("input.txt")
	if err != nil {
		fmt.Println("Error parsing input file:", err)
		return
	}

	for room, neighbors := range graph {
		fmt.Printf("Room: %s, Neighbors: ", room)
		for _, neighbor := range neighbors {
			fmt.Printf("%s ", neighbor)
		}
		fmt.Println()
	}
	helper.GetRooms("input.txt")
}
