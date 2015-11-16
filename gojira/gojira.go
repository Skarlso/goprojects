package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

var configFile = "jira_config.toml"
var parameter string
var flags struct {
	Comment     string
	Description string
	IssueID     string
	Priority    string
}

//Issue a representation of a JIRA Issue
type Issue struct {
	name     string
	id       string
	priority string
}

//Credentials a representation of a JIRA config which helds API permissions
type Credentials struct {
	Username string
	Password string
	URL      string
}

func init() {
	flag.StringVar(&flags.Comment, "m", "Default Comment", "A Comment when changing the status of an Issue.")
	flag.StringVar(&flags.Description, "d", "Default Description", "Provide a description for a newly created Issue.")
	flag.StringVar(&flags.Priority, "p", "Critical", "The priority of an Issue which will be set.")
	flag.StringVar(&flags.IssueID, "i", "", "Issue number of an issue.")
	flag.Parse()
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
	if len(flag.Args()) < 1 {
		log.Fatal("Please provide an action to take. Usage information:")
	}
	parameter = flag.Arg(0)
	switch parameter {
	case "close":
		closeIssue(flags.IssueID)
	case "start":
		startIssue(flags.IssueID)
	case "open":
		createIssue()
	}

	fmt.Println("Cred:", cred)
}

func closeIssue(issueID string) {
	if issueID == "" {
		printHelp()
		log.Fatal("Please provide an issueID with -i")
	}
	fmt.Println("Closing issue number: ", issueID)
}

func startIssue(issueID string) {
	if issueID == "" {
		printHelp()
		log.Fatal("Please provide an issueID with -i")
	}

	fmt.Println("Starting issue number:", issueID)
}

func createIssue() {
	fmt.Println("Creating new issue.")
}

func printHelp() {
	fmt.Println("Possible actions are: ")
	fmt.Println("-m 'Comment' -d 'Description' -p 'Priority' open")
	fmt.Println("-i 'Issue Number' close")
	fmt.Println("-i 'Issue Number' start")
}
