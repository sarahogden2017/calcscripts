package main

import (
	"os/exec"
	"fmt"
)

func main() {
	err := exec.Command("C:\\Windows\\System32\\calc.exe").Run()
	if err != nil {
		fmt.Println("Failed to start process:", err)
	}
}
