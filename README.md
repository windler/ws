# ws - workspace hero
`ws` is a command line tool that helps you handling your workspaces. Its purpose is to
- **list** all of your workspaces
- get **git information** about your workspaces, like **git status** and **current branch**

## Installation
`go get github.com/windler/ws`

## Usage
First, run `ws setup ws` to set your workspace directory. Then, you can run `ws` or `ws ls` to get workspace information.
```bash
ws ls
                    DIR                   |   GIT STATUS   | BRANCH
+-----------------------------------------+----------------+--------+
  /home/windler/projects/gittest          | UNMODIFED      | master
  /home/windler/projects/go               | Not a git repo | /

```

Type `ws -h` to get the helppage.

## Custom config path
The config file default to `~ /.wshero`. If you want to change the default file location you can set the `env WS_CFG`.

## Custom commands
You can create your own command which can be executed on your workspaces. With custom commands you can e.g.:
- start test environment
- run vsc commands
- run tests
- ...

To define you own commands edit your config file (default `~/.wshero`). The following example shows commands to start/stop an test environment and just print the current workspace:

```(yaml)
wsdir: /home/windler/projects/
parallelprocessing: 3
customcommands:
- name: pws
  description: "print the current ws name"
  cmd: echo
  args:
  - "{{.WSRoot}}"
- name: testenv_up
  description: "starts a dev environment in background"
  cmd: "docker-compose"
  args:
  - "-f"
  - "{{.WSRoot}}/project/docker-compose.yml"
  - "-p"
  - "{{.WSRoot}}"
  - "up"
  - "-d"
- name: testenv_down
  description: "stops the dev environment"
  cmd: docker-compose
  args:
  - "-f"
  - "{{.WSRoot}}/project/docker-compose.yml"
  - "-p"
  - "{{.WSRoot}}"
  - "down"
```

Custom command are also visible within the helppage

```(bash)
ws -h
(...)
COMMANDS:
     ls            List all workspaces with fancy information.
     setup         Configure everything to unleash the beauty. Alternatively, you can edit your personal config file.
     pws           print the current ws name
     testenv_up    starts a dev environment in background
     testenv_down  stops the dev environment
     help, h       Shows a list of commands or help for one command
(...)
```

### variables
You can use variables in your custom cammands using `go-template` syntax. The following variables are available:

| Variable       | Description                                        |
|----------------|----------------------------------------------------|
| WSRoot         | The absolute path of the current workspace         |