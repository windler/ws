# ws - workspace hero
`ws` is a command line tool that helps you handling your workspaces. Its purpose is to
- **list** all of your workspaces
- get **git information** about your workspaces, like **git status** and **current branch**

## Installation
`go get github.com/windler/ws`

## Usage
First, run `ws setup ws` to set your workspace directory. Then, you can run `ws` or `ws ls` or just try `ws -h`:
```bash
ws ls
                    DIR                   |   GIT STATUS   | BRANCH
+-----------------------------------------+----------------+--------+
  /home/windler/projects/gittest          | UNMODIFED      | master
  /home/windler/projects/go               | Not a git repo | /

```

```bash
ws -h

NAME:
   ws - workspace hero

USAGE:
   ws [global options] command [command options] [arguments...]

VERSION:
   0.0.1

DESCRIPTION:
   Dev Workspace Swiss Knife.

AUTHOR:
   Nico Windler

COMMANDS:
     ls       List all workspaces with fancy information.
     setup    Configure everything to unleash the beauty. Alternatively, you can edit your personal config file.
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --config FILE, -c FILE  Load configuration from FILE
   --help, -h              show help
   --version, -v           print the version

COPYRIGHT:
   2017
```