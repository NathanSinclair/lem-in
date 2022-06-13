package main

// Appends the x and assigns y to the first index of the x value. This is returned as a slice of string to prepend a string line by line.
func prependString(x []string, y string) []string {
	x = append(x, "")
	copy(x[1:], x)
	x[0] = y
	return x
}
