package main

// The custom comparePaths function ranges through
// a and b to return a boolean expression as to whether
// the pathNames are the same or not.
func comparePaths(a, b path) bool {
	for _, v := range a.pathNames {
		for _, r := range b.pathNames {
			if v == r && v != endRoom.name {
				return false
			}
		}
	}
	return true
}
