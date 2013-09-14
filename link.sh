#!/bin/bash
base_port=$1
node_count=$2


node=$(( $node_count - 1))
while [ $node -gt -1 ]; do

  peer=$(( $node_count - 1))
  while [ $peer -gt -1 ]; do
    if [ $peer != $node ]; then
      curl -X POST --data "localhost:$(( $peer + $base_port ))" "localhost:$(( $node + $base_port ))/add_peer"
    fi
    peer=$(( $peer - 1 ))
  done

  node=$(( $node - 1 ))
done
