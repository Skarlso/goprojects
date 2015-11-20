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
gojira -d 'Description' -p 'Priority' -t 'Issue Title' -o 'Project key like: IT' create
```

All available flags:

```go
Comment, "m", "Default Comment", "A Comment when changing the status of an Issue."
Description, "d", "Default Description", "Provide a description for a newly created Issue."
Priority, "p", "2", "The priority of an Issue which will be set."
IssueKey, "k", "", "Issue key of an issue."
Resolution, "r", "Done", "Resolution when an issue is closed. Ex.: Done, Fixed, Won't fix."
Title, "t", "Default Title", "Title of an Issue."
Project, "o", "IT", "Define a Project to create a ticket in."
```

Upcomming Features
------------------

- Ability to add in closeur reason.
- Ability to transition a ticket into any state.
- Ability to define transition state by name, rather than the current way which is by ID.
- Ability to add in a jira api token rather than a username password.
