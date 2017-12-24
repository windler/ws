# asd - workspace hero
`asd` is a command line tool that helps you handling your workspaces. Its purpose is to
- **list** all of your workspaces
- get **git information** about your workspaces, like **git status** and **current branch**

This tool is still under work. Its main purpose was to get used to the language `go`.

## Why the name?
This is a command line tool. You will use it often. You will type it often. So why bother typing some tool name? Typing `asd` is much simpler.

## Installation
`go get github.com/windler/asd`

## Usage
First, run `asd setup ws` to set your workspace directory. Then, you can run `asd` or `asd ls` or just try `asd -h`:
```bash
NAME:
   asd - workspace hero

USAGE:
   asd [global options] command [command options] [arguments...]

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

## Future work
- **start** and **stop** dev environment e.g. `docker-compose` or `vagrant` 
- **run tests** 