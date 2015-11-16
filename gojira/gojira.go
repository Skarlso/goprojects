package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
)

var configFile = "jira_config.toml"
var parameter string

var flags struct {
	Comment     string
	Description string
	IssueKey    string
	Priority    string
	Resolution  string
}

//Issue a representation of a JIRA Issue
type Issue struct {
	Project     string
	Summary     string
	Description string
	IssueType   string
	Priority    string
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
	flag.StringVar(&flags.IssueKey, "k", "", "Issue key of an issue.")
	flag.StringVar(&flags.Resolution, "r", "Done", "Resolution when an issue is closed. Ex.: Done, Fixed, Won't fix.")
	flag.Parse()
}

func (cred *Credentials) initConfig() {
	if _, err := os.Stat(configFile); err != nil {
		log.Fatalf("Error using config file: %v", err)
	}

	if _, err := toml.DecodeFile(configFile, cred); err != nil {
		log.Fatal("Error during decoding toml config: ", err)
	}
}

func main() {
	if len(flag.Args()) < 1 {
		log.Fatal("Please provide an action to take. Usage information:")
	}
	parameter = flag.Arg(0)
	switch parameter {
	case "close":
		closeIssue(flags.IssueKey)
	case "start":
		startIssue(flags.IssueKey)
	case "open":
		createIssue()
	}
}

func closeIssue(issueKey string) {
	cred := &Credentials{}
	cred.initConfig()
	if issueKey == "" {
		printHelp()
		log.Fatal("Please provide an issueID with -i")
	}
	fmt.Println("Closing issue number: ", issueKey)

	var jsonStr = `
			{
				"fields": {
				"resolution": {
					"name": "Done"
				}
				},
				"transition": {
					"id": "5"
				}
			}
		`
	jsonStr = strings.Replace(jsonStr, "<resolution>", flags.Resolution, -1)
	req, err := http.NewRequest("POST", cred.URL+issueKey+"/transitions", bytes.NewBuffer([]byte(jsonStr)))
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(cred.Username, cred.Password)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
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
	json := `
	{
		"fields": {
		    "project":{ "key": "BLD" },
		    "summary": "Release '$version'",
		    "description": "Release workload tracker.",
		    "issuetype": { "name": "Task" },
		    "priority": { "id": "2" }
		    }
	}
	`
	fmt.Println("Json: ", json)
}

func getIssueID(name string) string {
	return "id"
}

func printHelp() {
	fmt.Println("Possible actions are: ")
	fmt.Println("-m 'Comment' -d 'Description' -p 'Priority' open")
	fmt.Println("-i 'Issue Number' close")
	fmt.Println("-i 'Issue Number' start")
}
