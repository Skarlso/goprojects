package main

import (
	"fmt"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

var configFile = "jira_config.toml"

//Issue a representation of a JIRA Issue
type Issue struct {
	name        string
	issueNumber string
	priority    string
}

//Credentials a representation of a JIRA config which helds API permissions
type Credentials struct {
	Username string
	Password string
}

func initConfig() Credentials {
	if _, err := os.Stat(configFile); err != nil {
		log.Fatalf("Error using config file: %v", err)
	}

	var cred Credentials

	if _, err := toml.DecodeFile(configFile, &cred); err != nil {
		log.Fatal("Error during decoding toml config: ", err)
	}

	return cred
}

func main() {
	cred := initConfig()
	fmt.Println("Usage information")
	fmt.Printf("Username: %s, Password: %s\n", cred.Username, cred.Password)
}
