package main

import (
	"log"
	"os"
	"os/exec"
)

var templateFile string

func main() {
	if len(os.Args[1:]) == 0 {
		log.Fatal("Please provide a templateFile.")
	}

	templateFile = os.Args[1]

	if _, err := os.Stat(templateFile); err != nil {
		log.Fatalf("Error using template file: %s, %v", templateFile, err)
	}

	validate := exec.Command("packer", "validate", templateFile)
	err := validate.Run()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Validate was a success. Starting packer build, and not waiting for it to finish.")
	packer := exec.Command("packer", "build", templateFile)
	packer.Stdout = os.Stdout
	packer.Stderr = os.Stderr
	err = packer.Run()
	if err != nil {
		log.Fatal(err)
	}
}
