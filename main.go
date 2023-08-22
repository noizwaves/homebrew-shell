package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	binPath, err := exec.LookPath("zsh")
	if err != nil {
		fmt.Printf("error looking up zsh path: %s\n", err)
		return
	}

	ex, err := os.Executable()
	if err != nil {
		fmt.Printf("error getting current executable: %s\n", err)
		return
	}

	zdotdirPath := filepath.Join(filepath.Dir(ex), "zee-dot-dir")

	cmd := exec.Command(binPath)
	cmd.Env = []string{
		"ZDOTDIR=" + zdotdirPath,
	}
	cmd.Args = []string{}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("zsh returned an error: %s\n", err)
	}
}
