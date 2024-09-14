# Toodles

A simple, attractive to-do application built with Go and plenty of [charm](https://charm.sh/)

**Why another to-do app?** 

Well, the reason is simple: I am new to `go`, I wanted to build something with it, and being a fan of Taskwarrior, it's not quite appealing. So, I decided to bash my fingers on the keyboard and build something that is feature rich & also pleasing to the eye.

| ![1](https://ik.imagekit.io/rayshold/projects/toodles/1.png) | ![2](https://ik.imagekit.io/rayshold/projects/toodles/2.png) |
|---|---|

### Table of Contents

- [Installation](#installation) 
- [Commands](#commands) 
- [Local Development](#local-development)
- [Uninstalling](#uninstalling)


## Installation

**1.** Make sure you have `go` installed in your system

**2.** Run following commands which will install `toodles` on your system
```bash
git clone https://github.com/ImRayy/toodles.git
cd toodles/
go build && go install
```

**3.** Add following line to your shell config

For bash or zsh users add following line to `.bashrc` or `.zshrc`
```
export PATH="$HOME/go/bin:$PATH"
```

For fish users add following line to `config.fish`
```
set -gx PATH "$HOME/go/bin" $PATH
```

**4.** You're good, now you should be able to run `toodles` 

## Commands

```bash
# List all pending tasks 
toodles 

# Add a task
toodles add "Task Title"

# Edit a task
toodles edit <ID> "New Task Title"

# Set priority of a task ["low", "normal", "high"]
toodles priority <ID> high 

# Remove a task
toodles remove <ID>

# Mark task as complete
toodles done <ID>

# Undo task that marked as completed
toodles undo <ID>

# List tasks that are completed
toodles listdone

# List all tasks
toodles listall
```

## Local Development

#### Requirements

1. `go` >= `1.16`
2. `goose` >= `3.1.0` *(only if you want to create new schema, so optional)*  

Once `go` installed on your system you can install `goose` by running 

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```
[source](https://github.com/pressly/goose)

**To create new schema**
```bash
goose -dir=data/migrations create <table_name> sql
```

**Execute schema**
```bash
goose -dir=data/migrations sqlite3 <db_name.db> up # or down to drop table
```

## Uninstalling

Just remove `$HOME/go/bin/toodles`

```bash
rm ~/go/bin/toodles
```

# Credits

- Thanks to [taskwarrior](https://github.com/GothenburgBitFactory/taskwarrior) for functionality inspiration
- Thanks to [please](https://github.com/NayamAmarshe/please) for UI and README inspiration
