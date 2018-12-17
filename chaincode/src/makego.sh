#!/bin/bash
# Script to compile go modules of the custody asset use case
# Exit on first error, print all commands.
#set -ev

# remove existing main compiled module
mainfile="./main"
if [ -f "$mainfile" ]; then
    ls -l main
    rm main
fi

# now build the go modules
go build main.go data.go invokeCustodian.go invokeBank.go invokeExchange.go

if [ -f "$mainfile" ]; then
    ls -l main
fi
