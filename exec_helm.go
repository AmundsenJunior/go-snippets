package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {
	binary, lookErr := exec.LookPath("helm")
	if lookErr != nil {
		panic(lookErr)
	}

	args := os.Args[1:]

	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}

// References:
// https://gobyexample.com/execing-processes
