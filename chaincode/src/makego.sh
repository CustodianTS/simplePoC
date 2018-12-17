#!/bin/bash
# Script to compile go modules of the custody asset use case
# Exit on first error, print all commands.
set -ev

# remove existing main compiled module
rm main

# now build the go modules
go build main.go data.go invokeCustodian.go invokeBank.go

