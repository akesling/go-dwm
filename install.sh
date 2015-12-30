#!/usr/bin/env bash

PREFIX=/usr/local
BUILD_DIR=$(mktemp -d)

cd ${BUILD_DIR}

echo "Getting and building go-dwm in ${BUILD_DIR}."
go build github.com/akesling/go-dwm/cmd/go-dwm

echo "Installing ${pwd}/bin/go-dwm to ${PREFIX}/bin/dwm"
sudo cp -f go-dwm ${PREFIX}/bin/dwm
sudo chmod 755 ${PREFIX}/bin/dwm
