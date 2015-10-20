#!/bin/bash

GOARCH=amd64 GOOS=linux go build -o ./dingus ./src/main.go ./src/config.go ./src/log.go
