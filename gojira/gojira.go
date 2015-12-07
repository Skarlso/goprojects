package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/BurntSushi/toml"
)

var configFile = "~/.jira_config.toml"
var parameter string

var flags struct {
	Comment     string
	Description string
	IssueKey    string
	Priority    string
	Resolution  string
	Title       string
	Project     string
}

//Issue is a representation of a Jira Issue
type Issue struct {
	Fields struct {
		Project struct {
			Key string `json:"key"`
		} `json:"project"`
		Summary     string `json:"summary"`
		Description string `json:"description"`
		Issuetype   struct {
			Name string `json:"name"`
		} `json:"issuetype"`
		Priority struct {
			ID string `json:"id"`
		} `json:"priority"`
	} `json:"fields"`
}

//Transition defines a transition json object. Used for starting, stoppinp
//generally for state stranfer
type Transition struct {
	Fields struct {
		Resolution struct {
			Name string `json:"name"`
		} `json:"resolution"`
	} `json:"fields"`
	Transition struct {
		ID string `json:"id"`
	} `json:"transition"`
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
	flag.StringVar(&flags.Priority, "p", "2", "The priority of an Issue which will be set.")
	flag.StringVar(&flags.IssueKey, "k", "", "Issue key of an issue.")
	flag.StringVar(&flags.Resolution, "r", "Done", "Resolution when an issue is closed. Ex.: Done, Fixed, Won't fix.")
	flag.StringVar(&flags.Title, "t", "Default Title", "Title of an Issue.")
	flag.StringVar(&flags.Project, "o", "IT", "Define a Project to create a ticket in.")
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
	case "create":
		createIssue()
	}
}

func closeIssue(issueKey string) {
	if issueKey == "" {
		log.Fatal("Please provide an issueID with -k")
	}
	fmt.Println("Closing issue number: ", issueKey)

	var trans Transition

	//TODO: Add the ability to define a comment for the close reason
	trans.Fields.Resolution.Name = flags.Resolution
	trans.Transition.ID = "2"
	marhsalledTrans, err := json.Marshal(trans)
	if err != nil {
		log.Fatal("Error occured when marshaling transition: ", err)
	}
	fmt.Println("Marshalled:", trans)
	sendRequest(marhsalledTrans, "POST", issueKey+"/transitions?expand=transitions.fields")
}

func startIssue(issueID string) {
	if issueID == "" {
		log.Fatal("Please provide an issueID with -i")
	}

	fmt.Println("Starting issue number:", issueID)
}

func createIssue() {
	fmt.Println("Creating new issue.")
	var issue Issue
	issue.Fields.Description = flags.Description
	issue.Fields.Priority.ID = flags.Priority
	issue.Fields.Summary = flags.Title
	issue.Fields.Project.Key = flags.Project
	issue.Fields.Issuetype.Name = "Task"
	marshalledIssue, err := json.Marshal(issue)
	if err != nil {
		log.Fatal("Error occured when Marshaling Issue:", err)
	}
	sendRequest(marshalledIssue, "POST", "")
}

func sendRequest(jsonStr []byte, method string, url string) {
	cred := &Credentials{}
	cred.initConfig()
	fmt.Println("Json:", string(jsonStr))
	req, err := http.NewRequest(method, cred.URL+url, bytes.NewBuffer(jsonStr))
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
