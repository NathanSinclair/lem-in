package main

// compares path with next path and finds path that hasn't been used
func bestPaths(allPaths *[]path) bool {
	for i, path := range *allPaths {
		for u, nextPath := range *allPaths {
			if !comparePaths(path, nextPath) && i != u {
				if path.moves < nextPath.moves {
					*allPaths = removePath(*allPaths, u)
					if bestPaths(allPaths) {
						return true
					}
				} else {
					*allPaths = removePath(*allPaths, i)
					if bestPaths(allPaths) {
						return true
					}
				}
			}
		}
	}
	return true
}
