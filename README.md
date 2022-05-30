# load-balancer

[![CI](https://github.com/Ashilesh/load-balancer/actions/workflows/ci.yaml/badge.svg)](https://github.com/Ashilesh/load-balancer/actions/workflows/ci.yaml)
![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/Ashilesh/load-balancer?filename=src%2Fgo.mod)

A Layer-4 load balancer where you can distribute your workload to different servers.

## Run Executable

Download proper binaries according to your OS and architecture. You can download binaries from [here](https://github.com/Ashilesh/load-balancer/releases). While 
running executables you need to pass configuration file as [command line argument](#Command-Line-Arguments). Ex. For windows you can run executable by `load-balancer.exe -config config.json` where `config.json` is [configuration](#Configuration) file.


## Run Project

1. You need go v1.18 or higher to run this project. You can download Golang from [here](https://go.dev/dl/).
2. Clone or download this repository.
3. change directory to `load-balancer/src`
4. run `go mod tidy`
5. Create [configuration](#Configuration) file.
6. run `go run main.go -config <config-file-path>`

## Command Line Arguments

+ `-config <config-file-path>` For passing [configuration](#Configuration) file. This Argument is mandatory while running binaries/project.
+ `-log <log-dir-path>` will set log file directory. This argument is optional, by default it will store logs in current directory.

## Configuration

Configuration file must be in JSON format. You can modify load-balancer's configuration using these fields -

+ `host` It will set load-balancers host. ex `"host" : "localhost"`
+ `port` It will serve load balancer in this port. ex `"port" : "8080"`
+ `protocol` It will load balance this protocol. Allowed values are `websocket` and `http`. ex `"protocol" : "http"`
+ `networkType` Only tcp networkType is allowed. ex `"networkType" : "tcp"`
+ `nodes` It will contain URLs for server which we want to load balance. ex `"nodes" : ["localhost:7071"]`

### Example

<h5 a><strong><code>config.json</code></strong></h5>

```json 
{
    "host": "localhost",
    "port": "8080",
    "protocol": "http",
    "networkType": "tcp",
    "nodes": [
        "localhost:7071"
    ]
}
```
