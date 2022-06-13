package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
Returns a slice of the data read in a file. It will convert the xcoor and ycoor within the Room struct and the index of the file, from an int to a string. This will then be sliced.
*/

func ReadFile() []string {
	clearTerminal()

	if len(os.Args) != 2 {
		fmt.Println("ERROR: Invalid arguments")
		os.Exit(1)
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("ERROR: failed to open file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var start, end bool
	ants = 0
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		line := scanner.Text()
		var Rooms room
		if strings.Contains(line, " ") {
			split := strings.Split(line, " ")
			Rooms.name = split[0]
			Rooms.xcoor, _ = strconv.Atoi(split[1])
			Rooms.ycoor, _ = strconv.Atoi(split[2])
			if start {
				Rooms.start = true
				Rooms.visited = true
				start = false
			} else if end {
				Rooms.end = true
				end = false
			}
		}
		if line == "##start" {
			start = true
		} else if line == "##end" {
			end = true
		}
		if Rooms.name != "" {
			farmMap = append(farmMap, Rooms)
		}
		i, err := strconv.Atoi(line)
		if err == nil {
			ants = i
		}
		if strings.Contains(line, "-") {
			link := strings.Split(line, "-")
			for i, room := range farmMap {
				if room.name == link[0] {
					farmMap[i].connections = append(farmMap[i].connections, link[1])
				}
				if room.name == link[1] {
					farmMap[i].connections = append(farmMap[i].connections, link[0])
				}
			}
		}
	}
	return nil
}

// Get values and fill them to our structs & farmMap
func GetData() []string {
	var startBool, endBool bool
	for _, room := range farmMap {
		if room.start {
			startRoom = room
			startBool = true
		}
		if room.end {
			endRoom = room
			endBool = true
		}
	}
	if !startBool || !endBool {
		clearTerminal()
		fmt.Println("ERROR: invalid data format, no start room found")
		os.Exit(0)
	}
	if ants <= 0 {
		clearTerminal()
		fmt.Println("ERROR: invalid data format, invalid number of ants")
		os.Exit(0)
	}
	var emptyPath path
	for _, room := range farmMap {
		if len(room.connections) <= 0 {
			clearTerminal()
			fmt.Println("ERROR: invalid data format")
			os.Exit(0)
		}
	}
	for _, startLink := range startRoom.connections {
		for i, room := range farmMap {
			if room.name == startLink {
				paths = append(paths, solveFarmDown(emptyPath, farmMap, i))
			}
		}
	}
	emptyPath.pathNames = append(emptyPath.pathNames, endRoom.name)
	for _, endLink := range endRoom.connections {
		for i, room := range farmMap {
			if room.name == endLink {
				paths = append(paths, solveFarmUp(emptyPath, farmMap, i))
			}
		}
	}
	for i, path := range paths {
		for u, name := range path.pathNames {
			if name == startRoom.name {
				paths[i].pathNames = removeString(paths[i].pathNames, u)
			}
		}
	}
	removeDuplication()
	paths = fastestPaths(paths, ants)
	lNumber := 1
	fmt.Println()
	for ants > 0 {
		for i, currentPath := range paths {
			if currentPath.ants > 0 {
				var tempPathInfo pathInfo
				tempPathInfo.pathName = "L" + fmt.Sprint(lNumber) + "-"
				tempPathInfo.currentRoomIndex = 0
				tempPathInfo.pathUsed = currentPath
				tempPathInfo.finished = false
				allPathInfo = append(allPathInfo, tempPathInfo)
				paths[i].ants--
				lNumber++
			}
		}
		for _, info := range allPathInfo {
			if !info.finished && info.currentRoomIndex < len(info.pathUsed.pathNames) {
				fmt.Print(info.pathName + info.pathUsed.pathNames[info.currentRoomIndex] + " ")
			}
		}
		fmt.Println()
		for i := range allPathInfo {
			if allPathInfo[i].currentRoomIndex < len(allPathInfo[i].pathUsed.pathNames) {
				allPathInfo[i].currentRoomIndex++
			} else if !allPathInfo[i].finished {
				allPathInfo[i].finished = true
				ants--
			}
		}
	}
	return nil
}
