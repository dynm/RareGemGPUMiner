#!/bin/bash

CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc CGO_LDFLAGS='-L.' go build -v -o raregem.exe