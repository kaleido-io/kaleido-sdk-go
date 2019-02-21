#!/bin/bash

FILE="$HOME/.kld.yaml"

/bin/cat <<EOM >$FILE
api:
  url: $1
  key: $2
  debug: true
services:
  idregistry:
    id: $3
networks:
  url: $4
EOM

