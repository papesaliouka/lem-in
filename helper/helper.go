package helper

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Graph map[string][]string
type Room struct {
	Name     string
	RoomType string
}

func GetRooms(filename string) []Room {
	rooms := []Room{}
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}

	lines := strings.Split(string(content), "\n")
	for i, line := range lines {
		if strings.Contains(line, "-") {
			continue
		}
		if line == "##start" {
			room := Room{
				Name:     string(lines[i+1][0]),
				RoomType: "start",
			}
			rooms = append(rooms, room)
		}
		if line == "##end" {
			room := Room{
				Name:     string(lines[i+1][0]),
				RoomType: "end",
			}
			rooms = append(rooms, room)
		}
	}
	return rooms
}

func ParseInputFile(filename string) (Graph, int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	graph := make(Graph)

	var numAnts int
	var numRooms int

	for scanner.Scan() {
		line := scanner.Text()
		if line == "##start" || line == "#" || line == "##end" {
			continue
		}

		if numAnts == 0 {
			numAnts, err = strconv.Atoi(line)
			if err != nil {
				fmt.Println("ERROR: invalid data format")
				return nil, 0, nil
			}
		} else if strings.Contains(line, "-") {
			parts := strings.Split(line, "-")
			from := parts[0]
			to := parts[1]
			graph.AddArc(from, to)
		} else {
			parts := strings.Split(line, " ")
			if len(parts) != 3 {
				// Line does not have three elements

				fmt.Println("ERROR: invalid data format")
				return nil, 0, nil
			}
			name := parts[0]
			first := parts[1]
			second := parts[2]

			_, err := strconv.Atoi(first)
			if err != nil {
				// not an integer
				fmt.Println("ERROR: invalid data format")
				return nil, 0, nil
			}
			_, err = strconv.Atoi(second)
			if err != nil {
				// not an integer
				fmt.Println("ERROR: invalid data format")
				return nil, 0, nil
			}
			graph.AddRoom(name)
			numRooms++
		}

	}

	err = scanner.Err()
	if err != nil {
		return nil, 0, err
	}

	return graph, numRooms, nil
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
