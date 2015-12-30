#!/usr/bin/env bash

PREFIX=/usr/local

echo "Getting and building go-dwm."
go get github.com/akesling/go-dwm

echo "Installing ${GOPATH}/bin/go-dwm to ${PREFIX}/bin/dwm"
sudo cp -f ${GOPATH}/bin/go-dwm ${PREFIX}/bin/dwm
sudo chmod 755 ${PREFIX}/bin/dwm
