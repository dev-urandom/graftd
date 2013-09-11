graftd
======

A trivial example server using [graft](https://github.com/dev-urandom/graft)

![get_on_this_level](https://i.minus.com/ibd7dgjN0UsxhZ.gif)

## Setup

0. Clone
1. Install dependences `go get`
2. Build with `go build`

## Usage

### Manual

`./graftd` will start graftd on the default port 7777. You can pass in a port with the `-port` flag. Here is an example that starts graftd on 8888: `./graftd -port=8888`

### Foreman

Normally you want a cluster of graftds. You can use foreman to do this. `foreman start` will start a cluster of 3 servers starting on port 5000 and incrementing by one. You can set how big of a cluster you want with the `-c` flag. `foreman start -c graft=10` will start a cluster with 10 servers.
