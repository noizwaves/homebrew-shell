package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	binPath, err := exec.LookPath("zsh")
	if err != nil {
		fmt.Printf("error looking up zsh path: %s\n", err)
		return
	}

	cmd := exec.Command(binPath)
	cmd.Env = []string{"ZDOTDIR=/Users/adam.neumann/workspace/homebrew-shell/zee-dot-dir"}
	cmd.Args = []string{}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("zsh returned an error: %s\n", err)
	}
}
