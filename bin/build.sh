#!/bin/bash

GOARCH=amd64 GOOS=linux go build -o dingus main.go config.go log.go
