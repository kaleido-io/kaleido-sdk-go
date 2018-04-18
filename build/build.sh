#!/bin/bash

go get ./...
go build -o kld-$1 -v