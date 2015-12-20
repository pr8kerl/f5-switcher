# f5-switcher

A web interface and REST api to more easily manage F5 application cutovers/deployments.

Only currently supports explicitly enabling or disabling pool members.

Uses Golang server-side and vue.js client side.

No auth built in to the app yet. Relying on web server in front of app to manage that (ie client certificate auth or similar).

## Features
* view status of logical application stacks (blue/green)
* switch between blue/green application stacks by updating F5 poolmembers
* switch between blue/green application stacks by modifying F5 virtual IP objects (changing the pool) - not yet implemented.

## Installation

* clone the repo

```
$ git clone git@github.com:pr8kerl/f5-switcher.git
Cloning into 'f5-switcher'...
remote: Counting objects: 181, done.
remote: Total 181 (delta 0), reused 0 (delta 0), pack-reused 181
Receiving objects: 100% (181/181), 1.36 MiB | 457.00 KiB/s, done.
Resolving deltas: 100% (89/89), done.
Checking connectivity... done.
$ cd f5-switcher/
```

* install all golang package dependencies 

```
$ make update
GOPATH=/home/ians/work/f5-switcher/tmp/f5-switcher go get -u github.com/gin-gonic/gin github.com/jmcvetta/napping github.com/pr8kerl/f5-switcher/F5
$ 
```

* build the server

```
ians@module:~/work/f5-switcher/tmp/f5-switcher$ make
GOPATH=/home/ians/work/f5-switcher/tmp/f5-switcher go fmt main.go config.go group.go
GOPATH=/home/ians/work/f5-switcher/tmp/f5-switcher go build -o server -v main.go config.go group.go
command-line-arguments
touch server
$
```

* create a working config

```
cp config.json.example config.json
vi config.json # edit to your needs
chmod 600 config.json
```

* run the server

```
$ ./server 
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET   /public/*filepath         --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (5 handlers)
[GIN-debug] HEAD  /public/*filepath         --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (5 handlers)
[GIN-debug] GET   /api/group                --> main.showGroup (5 handlers)
[GIN-debug] PUT   /api/group                --> main.putGroup (5 handlers)
[GIN-debug] Listening and serving HTTP on 127.0.0.1:5000
```

* view the app in your browser
![view app in browser](/screens/screen-index.png?raw=true "web app view")

* click on a stack and then click blue or green to make the switch
![switch a stack in browser](/screens/screen-modal.png?raw=true "web app view")


## REST api

You can use the API directly without the browser.


Endpoint  | Description
------------- | -------------
GET /api/group  | retrieve a list of all logical groups and their resources
PUT /api/group  | request a switch of a group to a new state

The PUT request expects the following json to be provided in the body.

```
{
  "name": "logical-group-name",
  "state": "blue"
}
```

