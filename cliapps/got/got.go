package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args[1:]) == 0 {
		log.Fatal("Please provide some arguments for git")
	}

	git := exec.Command("git", os.Args[1:]...)
	git.Stdout = os.Stdout
	git.Stderr = os.Stderr
	if err := git.Run(); err != nil {
		log.Fatal("Problem running git:", err)
	}
}
