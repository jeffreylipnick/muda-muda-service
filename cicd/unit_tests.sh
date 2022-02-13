#!/usr/bin/env bash

# Run this script from the root of the repo.
#
# Unfortunately I haven't found an easy way to all the go tests from the root.
# With running the tests like this, there's a danger that someone could add a 
# nested module and forget to update this list.  However, this approach is very
# simple, and at a certain point you have to trust the maintainers of a service
# to catch this kind of stuff during PR review.

root_dir=$(pwd)
go test
cd ${root_dir}/api && go test
