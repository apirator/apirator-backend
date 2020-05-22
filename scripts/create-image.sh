#!/usr/bin/env bash

set -e

function run {
    echo "$(tput setaf 6)$1$(tput sgr0)"
    eval $1
}

VERSION=$(grep -o '".*"' ./version/version.go | sed 's/"//g')
run "docker build -t apirator/apirator-backend:$VERSION ."
run "docker push apirator/apirator-backend:$VERSION"

