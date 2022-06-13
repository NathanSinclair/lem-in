package main

import (
	"os"
	"os/exec"
)

// method body - Deals with all errors when called and clear terminal for new arguments
func clearTerminal() string {
	clear := exec.Command("clear")
	clear.Stdout = os.Stdout
	clear.Run()
	return ""
}
