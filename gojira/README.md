Go Jira Client
==============

Simple Jira client written in Go.

Usage
-----

Some information about the cli commands.

```bash
# This will default outcome to 'Done'
gojira -k <IssueKey> close

# You can overwrite outcome with -o
gojira -o "Won't Fix" -k <IssueKey> close
```

```bash
# To a new issue. Priority has to be 1-5
gojira -d 'Description' -p 'Priority' -t 'Issue Title' -o 'Project like: IT' create
```

All available flags:

```go
flag.StringVar(&flags.Comment, "m", "Default Comment", "A Comment when changing the status of an Issue.")
flag.StringVar(&flags.Description, "d", "Default Description", "Provide a description for a newly created Issue.")
flag.StringVar(&flags.Priority, "p", "2", "The priority of an Issue which will be set.")
flag.StringVar(&flags.IssueKey, "k", "", "Issue key of an issue.")
flag.StringVar(&flags.Resolution, "r", "Done", "Resolution when an issue is closed. Ex.: Done, Fixed, Won't fix.")
flag.StringVar(&flags.Title, "t", "Default Title", "Title of an Issue.")
flag.StringVar(&flags.Project, "o", "IT", "Define a Project to create a ticket in.")
```

Issues: transitions throws an internal server error from JIRA. I have no idea why, because jira doesn't return an error.
I'm guessing it has something to do with how the JSON looks like as creating an issue works fine.
