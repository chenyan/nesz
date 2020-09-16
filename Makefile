GOCMD=go
GOBUILD=${GOCMD} build
GOCLEAN=${GOCMD} clean
GOENV=GOOS=linux GOARCH=amd64
GOVER=$(shell go version)
BUILDTIME=$(shell date +%F-%Z/%T)
GITBRANCH=$(shell git symbolic-ref HEAD)
GITCOMMIT=$(shell git rev-parse --short HEAD)
GITMESSAGE=$(shell git log -1 --oneline)

build: nesz

.PHONY: nesz

clean:
	${GOCLEAN}

nesz:
	${GOBUILD} nesz/cmd/nesz

