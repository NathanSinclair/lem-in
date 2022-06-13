package main

// Appends the currentpath, farm from a slice of the room struct) and index to the end of the path.
func solveFarmDown(currentpath path, farm []room, i int) path {
	currentpath.pathNames = append(currentpath.pathNames, farm[i].name)
	currentpath.moves++
	farm[i].visited = true
	if farm[i].end {
		return currentpath
	}
	xCord := 0
	dex := -1
	for _, link := range farm[i].connections {
		for i, room := range farm {
			if room.name == link {
				if room.end {
					return solveFarmDown(currentpath, farm, i)
				}
				if xCord < room.xcoor {
					xCord = room.xcoor
					dex = i
				}
			}
		}
	}
	if dex >= 0 {
		checkForEnd := solveFarmDown(currentpath, farm, dex)
		if checkForEnd.pathNames[len(checkForEnd.pathNames)-1] == endRoom.name {
			return checkForEnd
		}
	}
	return currentpath
}

// Prepends the currentPath, farm (from a slice of the room struct) and index of the path, from the end of the path.
func solveFarmUp(currentPath path, farm []room, i int) path {
	currentPath.pathNames = prependString(currentPath.pathNames, farm[i].name)
	currentPath.moves++
	farm[i].visited = true
	if farm[i].start {
		return currentPath
	}
	xCord := 100
	dex := -1
	for _, link := range farm[i].connections {
		for i, room := range farm {
			if room.name == link {
				if room.start {
					return solveFarmUp(currentPath, farm, i)
				}
				if xCord > room.xcoor {
					xCord = room.xcoor
					dex = i
				}
			}
		}
	}
	if dex >= 0 {
		checkForStart := solveFarmUp(currentPath, farm, dex)
		if checkForStart.pathNames[0] == startRoom.name {
			return checkForStart
		}
	}
	return currentPath
}
