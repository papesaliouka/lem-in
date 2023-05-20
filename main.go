package main

import (
	"fmt"
	"lemin/helper"
)

func main() {
	graph, numRooms, err := helper.ParseInputFile("input.txt")
	if err != nil {
		fmt.Println("Error parsing input file:", err)
		return
	}
	rooms := helper.GetRooms("input.txt")

	table := helper.DFS(graph, rooms[0].Name)

	if numRooms != len(table) {
		fmt.Println("ERROR: Invalid data format")
		return
	}

	// for room, neighbors := range graph {
	// 	fmt.Printf("Room: %s, Neighbors: ", room)
	// 	for _, neighbor := range neighbors {
	// 		fmt.Printf("%s ", neighbor)
	// 	}
	// 	fmt.Println()
	// }

}
