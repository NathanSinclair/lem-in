package main

type pathInfo struct {
	pathName         string
	currentRoomIndex int
	pathUsed         path
	finished         bool
}

type room struct {
	name        string
	xcoor       int
	ycoor       int
	start       bool
	end         bool
	visited     bool
	connections []string
}

type path struct {
	moves     int
	ants      int
	pathNames []string
}

var (
	farmMap     []room
	startRoom   room
	endRoom     room
	paths       []path
	ants        int
	allPathInfo []pathInfo
)
