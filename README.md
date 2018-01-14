# ws - workspace hero
`ws` is a command line tool that helps you handling your workspaces. Its purpose is to
- **list** all of your workspaces
- get **git information** about your workspaces, like **git status** and **current branch**
- run **custom commands** like start an editor or run tests

## Installation
`go get github.com/windler/ws`

## Usage
First, you have to set your workspace directory in your config file. It is located at `~/.wshero`. Add a line similar to the following:
```yaml
wsdir: /home/windler/projects/
```
 Then, you can run `ws -h` to show the help page or `ws ls` to get workspace information.
```bash
ws ls
                    DIR                   |   GIT STATUS   | BRANCH
+-----------------------------------------+----------------+--------+
  /home/windler/projects/gittest          | UNMODIFED      | master
  /home/windler/projects/go               | Not a git repo | /

```

## Custom config path
The config file default to `~ /.wshero`. If you want to change the default file location you can set the `env WS_CFG`.

## Custom commands
You can create your own command which can be executed on your workspaces. With custom commands you can e.g.:
- start test environment
- run vsc commands
- run tests
- start editor
- ...

To define you own commands edit your config file (default `~/.wshero`). The following example shows commands to start/stop an test environment and just print the current workspace:

```yaml
wsdir: /home/windler/projects/
parallelprocessing: 3
tableformat: "{{cmd \"pws\" .}}|{{gitStatus .}}|{{gitBranch .}}"
customcommands:
- name: pws
  description: "print the current ws name"
  cmd: echo
  args:
  - "{{.WSRoot}}"
- name: code
  description: "edit ws in vscode"
  cmd: "code"
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

When you run a custom command it will be executed in the current workspace. If you want to run it in a specific workspace pass a pattern as the first argument. The first workspace that matches your pattern will be used. E.g. if you want to start your editor for the workspace `/home/windler/projects/gittest` using the `code` custom command type the following: 
```bash
ws code gittest
``` 

Custom command are also visible within the help-page

```bash
ws -h
(...)
COMMANDS:
     ls            List all workspaces with fancy information.
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

### cd to workspace
The `ws` command cannot change the `/proc/<PID>/cwd` of the terminal. Therefore, it is not possible to create a command that changes the terminals directory to a workspace root. As a workaround you can create a `.bashrc` / `.zhsrc` function that wraps the `ws` command and creates its own cd command. Assuming you have defined the above `pws` custom command, such a function could look like this:

```bash
 w() {
    case "$1" in
        cd)
            cd $(ws pws ${@:2})
            return 0
            ;;
    esac

    ws $@
 }
``` 

Calling `w cd my_ws` will change the terminals directory into the workspace root of `my_ws`. All other commands or flags will be delegated to the `ws` command.  

## Custom table layout
You can modify the table from the `ls` command by passing the flag `--table pattern` or permamently by setting `tableformat` in the config file. The columns are separated by the pipe (`|`). You have to use the `go-template` syntax. The template is feeded with the workspace dir. The following functions are available for the output:

| Function            | Description                                                                                 |
|---------------------|---------------------------------------------------------------------------------------------|
| wsRoot (dir)        | Prints the directory and adds an arrow if your current working direcotry is withing the dir |
| gitStatus (dir)     | Prints the git status of the dir                                                            |
| gitBranch (dir)     | Prints the current git branch of the dir                                                    |
| cmd (name, dir)     | Runs the custom command in the dir and prints the output                                    |

E.g. to print the current branch and the output of a custom command "pws" the pattern is the following: 
```bash
{{gitBranch .}}|{{cmd "pws" .}}
```