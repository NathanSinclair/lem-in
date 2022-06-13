package main

// Searches path and re-orders the paths highest to lowest
func pathHighToLow(paths *[]path) []path {
	for i := range *paths {
		for u := range *paths {
			if (*paths)[i].moves >= (*paths)[u].moves {
				(*paths)[i], (*paths)[u] = (*paths)[u], (*paths)[i]
			}
		}
	}
	return *paths
}

// Ranges through the paths to rearrange the paths from lowest to highest.
func pathLowToHigh(paths *[]path) []path {
	for i := range *paths {
		for u := range *paths {
			if (*paths)[i].moves <= (*paths)[u].moves {
				(*paths)[i], (*paths)[u] = (*paths)[u], (*paths)[i]
			}
		}
	}
	return *paths
}

// The custom FastestPaths function makes a path that checks
// for the most efficient path from the top to bottom.
func fastestPaths(paths []path, ants int) []path {
	pathsCalc := make([][]path, len(paths))
	for i, path := range paths {
		pathsCalc[i] = append(pathsCalc[i], path)
		for u, nextPath := range paths {
			if comparePaths(path, nextPath) && i != u {
				for _, pathAfter := range pathsCalc[i] {
					if comparePaths(nextPath, pathAfter) {
						pathsCalc[i] = append(pathsCalc[i], nextPath)
					}
				}
			}
		}
	}
	for index := range pathsCalc {
		pathsCalc[index] = removeDuplicationPaths(pathsCalc[index])
	}
	for pIndex := range pathsCalc {
		bestPaths(&pathsCalc[pIndex])
	}
	for pIndex := range pathsCalc {
		pathHighToLow(&pathsCalc[pIndex])
	}
	for i, pathArr := range pathsCalc {
		high := pathArr[0].moves
		antToUse := ants
		for u, path := range pathArr {
			pathsCalc[i][u].ants = high - path.moves
			antToUse -= pathsCalc[i][u].ants
		}
		for antToUse >= 0 {
			for u := len(pathsCalc[i]) - 1; u >= 0; u-- {
				if antToUse > 0 {
					pathsCalc[i][u].ants++
					antToUse--
				} else {
					antToUse--
				}
			}
		}
	}
	for pIndex := range pathsCalc {
		pathLowToHigh(&pathsCalc[pIndex])
	}
	lowest := 100
	lowestPath := pathsCalc[0]
	for _, pathAll := range pathsCalc {
		if (pathAll[0].ants + pathAll[0].moves) < lowest {
			lowestPath = pathAll
			lowest = pathAll[0].ants + pathAll[0].moves
		}
	}
	return lowestPath
}
