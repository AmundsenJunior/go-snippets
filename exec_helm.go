// $ go run exec_helm.go init --help
package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {
	args := os.Args
	args[0] = "helm"

	env := os.Environ()

	execHelm(args, env)
}

func execHelm(args []string, env []string) {
	helm, lookErr := exec.LookPath("helm")
	if lookErr != nil {
		panic(lookErr)
	}

	execErr := syscall.Exec(helm, args, env)
	if execErr != nil {
		panic(execErr)
	}
}

// References:
// https://gobyexample.com/execing-processes
