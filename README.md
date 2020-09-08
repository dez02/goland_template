# golang-template

## Requirements

- [Git](https://git-scm.com/)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- [GNU/Make](https://www.gnu.org/software/make/)

## Installation

```console
$ git clone https://github.com/aminnairi/golang-template
$ cd golang-template
```

Notes:
- Replace `aminnairi/golang-template` with `USERNAME/PROJECTNAME` where `USERNAME` is your GitHub's username and `PROJECTNAME` is the name of your project in [`go.mod`](./go.mod)
- Replace `aminnairi/golang-template` with `USERNAME/PROJECTNAME` where `USERNAME` is your GitHub's username and `PROJECTNAME` is the name of your project in [`main.go`](./main.go)

## Commands

Command | Description
---|---
`make start` | Start the Docker Compose services.
`make stop` | Stop the Docker Compose services.
`make restart` | Restart the Docker Compose services.
`make shell` | Start a shell into the Go container.
