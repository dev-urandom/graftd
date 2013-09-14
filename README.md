![get_on_this_level](https://i.minus.com/ibd7dgjN0UsxhZ.gif)
# graftd

A trivial example server using [graft](https://github.com/dev-urandom/graft)

## Setup

0. `git clone https://github.com/dev-urandom/graftd.git`
1. `go get`
2. `go build`

## Usage

### Manual

`./graftd` will start graftd on the default port 7777. You can pass in a port with the `-port` flag. Here is an example that starts graftd on 8888: `./graftd -port=8888`

### Foreman

Normally you want a cluster of graftds. You can use foreman to do this. `foreman start` will start a cluster of 3 servers starting on port 5000 and incrementing by one. You can set how big of a cluster you want with the `-c` flag. `foreman start -c graftd=10` will start a cluster with 10 servers.

You can also use `./start.sh` which takes an optional cluster size (the default is 3). `start.sh` will run `foreman start -c graftd=<size>` and `./join.sh <base port> <size>` (`base_port` is `5000` by default) which will add all the servers to each other. This will result in a cluster that is fully linked.

### API

Server status

```bash
$ curl -X GET localhost:5000/status
# => 
{
    "id": "localhost:5000",
    "lastCommitIndex": 0,
    "lastLogIndex": 0,
    "peers": [],
    "state": "follower",
    "term": 1,
    "votedFor": ""
}
```

Start election

```bash
$ curl -X POST localhost:5000/start_election
# => {"message":"started election"}
```

Append entry

```bash
$ curl -X POST localhost:5000/append_entry --data "foo"
# => {"message":"commited to log"}
```

Get log

```bash
$ curl -X GET localhost:5000/log
# => {"log":[{"Term":2,"Data":"foo"}]}
```

Add peer (you shouldn't need to do this is you use `start.sh`

```bash
$ curl -X POST localhost:5000/add_peer --data "localhost:5001"
```
