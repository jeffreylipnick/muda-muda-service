#!/usr/bin/env bash

# Unfortunately there doesn't seem to be an easy way to all the go tests from the root.
# So this script is used to find all the directories with go.mod files and run the PR checks
# from inside that directory.
# Reference: https://stackoverflow.com/questions/9612090/how-to-loop-through-file-names-returned-by-find

find . -name go.mod -execdir bash -c "go test" {} \;
