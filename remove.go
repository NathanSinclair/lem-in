package main

// returns true for the path names found.
func removeDuplication() bool {
	for i := 0; i < len(paths); i++ {
		for u := 0; u < len(paths); u++ {
			if compareSlices(paths[i].pathNames, paths[u].pathNames) && i != u {
				paths = removePath(paths, i)
				if removeDuplication() {
					return true
				}
			}
		}
	}
	return true
}

// returns a slice of the path struct and removes the paths with identical names.
func removeDuplicationPaths(paths []path) []path {
	for i := 0; i < len(paths); i++ {
		for u := 0; u < len(paths); u++ {
			if compareSlices(paths[i].pathNames, paths[u].pathNames) && i != u {
				paths = removePath(paths, i)
				i--
				u--
			}
		}
	}
	return paths
}

// appends the editedPath and index to the end of the slice of string.
func removeString(editedPath []string, i int) []string {
	return append(editedPath[:i], editedPath[i+1:]...)
}

// appends the editedPath and index to the end of the slice of path.
func removePath(modifiedPaths []path, i int) []path {
	return append(modifiedPaths[:i], modifiedPaths[i+1:]...)
}
