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

	argumentList := []string{"build"}
	if len(os.Args[1:]) > 1 {
		argumentList = append(argumentList, os.Args[2:]...)
	}
	argumentList = append(argumentList, templateFile)

	if _, err := os.Stat(templateFile); err != nil {
		log.Fatalf("Error using template file: %s, %v", templateFile, err)
	}

	checkIfScriptsExistInTemplate()

	validate := exec.Command("packer", "validate", templateFile)
	err := validate.Run()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Validate was a success. Starting packer with arguments:", argumentList)
	packer := exec.Command("packer", argumentList...)
	packer.Stdout = os.Stdout
	packer.Stderr = os.Stderr
	err = packer.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func checkIfScriptsExistInTemplate() {

}
