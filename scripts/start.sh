#!/usr/bin/env sh

cd $(dirname "$0")/..
ROOT_DIR=$(pwd)

go run cmd/coolsim/main.go
