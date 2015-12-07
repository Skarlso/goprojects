package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

var templateFile string

//Provisioners Provisioners in the json file
type Provisioners struct {
	Type    string   `json:"type"`
	Scripts []string `json:"scripts"`
}

//Builders Builders in the Json file
type Builders struct {
	Type   string   `json:"type"`
	Floppy []string `json:"floppy_files"`
}

//Scripts a bundle of scripts I'd like to check for access
type Scripts struct {
	Provisioners []Provisioners
	Builders     []Builders
}

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
	var scripts Scripts
	content, err := ioutil.ReadFile(templateFile)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(content, &scripts)
	if err != nil {
		log.Fatal("Error during unmarshaling json:", err)
	}
	log.Println("Checking if all Builders scripts are available...")
	for _, builder := range scripts.Builders {
		for _, file := range builder.Floppy {
			log.Printf("Checking file: %s in builder: %s\n", file, builder.Type)
			if _, err := os.Stat(file); err != nil {
				log.Fatalf("Error using template file: %s, %v", templateFile, err)
			}
		}
	}
	log.Println("All builder files found. Moving on to Provisioners.")
	for _, provisioner := range scripts.Provisioners {
		for _, file := range provisioner.Scripts {
			log.Printf("Checking file: %s in provisioner: %s\n", file, provisioner.Type)
			if _, err := os.Stat(file); err != nil {
				log.Fatalf("Error using template file: %s, %v", templateFile, err)
			}
		}
	}

	log.Println("All done. Good to go.")
}
