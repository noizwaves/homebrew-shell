package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	helpFlag := flag.Bool("help", false, "display usage")
	flag.Parse()

	if *helpFlag {
		fmt.Printf("Usage: %s [OPTIONS]\n", os.Args[0])
		fmt.Println("Options:")
		flag.PrintDefaults()
		return
	}

	brewPath, err := exec.LookPath("brew")
	if err != nil {
		fmt.Printf("error finding 'brew' executable: %s\n", err)
		return
	}
	zshPath, err := exec.LookPath("zsh")
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

	cmd := exec.Command(brewPath)
	cmd.Args = []string{
		"IGNORED", // TODO: figure out why this placeholder first arg is required
		"bundle",
		"exec",
		"--no-upgrade", // pls no network calls
		"--",
		"env",
		// force zsh to use our special zshrc file
		"ZDOTDIR=" + zdotdirPath,
		zshPath,
	}
	cmd.Env = []string{
		// required by homebrew
		"HOME=" + os.Getenv("HOME"),
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("zsh returned an error: %s\n", err)
	}
}
