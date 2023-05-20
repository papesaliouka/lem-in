package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Graph map[string][]string

func ParseInputFile(filename string) (Graph, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	graph := make(Graph)

	var numAnts int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "#") {
			continue
		}

		if numAnts == 0 {
			numAnts, err = strconv.Atoi(line)
			if err != nil {
				fmt.Println("ERROR: invalid data format")
				return nil, nil
			}
		} else if line == "##start" || line == "##end" {
			continue
		} else if strings.Contains(line, "-") {
			parts := strings.Split(line, "-")
			from := parts[0]
			to := parts[1]
			graph.AddArc(from, to)
		} else {
			parts := strings.Split(line, " ")
			name := parts[0]
			graph.AddRoom(name)
		}
	}

	err = scanner.Err()
	if err != nil {
		return nil, err
	}

	return graph, nil
}

func (g Graph) AddRoom(name string) {
	if _, exists := g[name]; !exists {
		g[name] = []string{}
	}
}

func (g Graph) AddArc(from, to string) {
	g[from] = append(g[from], to)
	g[to] = append(g[to], from)
}

func main() {
	graph, err := ParseInputFile("input.txt")
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
}
