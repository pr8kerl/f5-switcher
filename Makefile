GOROOT := /usr/local/go
GOPATH := $(shell pwd)
GOBIN  := $(GOPATH)/bin
PATH   := $(GOROOT)/bin:$(PATH)
DEPS   := github.com/gin-gonic/gin github.com/jmcvetta/napping github.com/pr8kerl/f5-switcher/F5

all: server

update: $(DEPS)
	GOPATH=$(GOPATH) go get -u $^

server: main.go config.go group.go
    # always format code
		GOPATH=$(GOPATH) go fmt $^
    # binary
		GOPATH=$(GOPATH) go build -o $@ -v $^
		touch $@

windows:
	  gox -os="windows"

.PHONY: $(DEPS) clean

clean:
	rm -f server
