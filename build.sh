#!/bin/bash

export GOPATH=$(pwd)
export GOROOT=$(find /usr/lib -type d -name "go-*"  | tail -n 1)

go install oaem
