#!/bin/bash

fswatch -o "./" -e ".*" -i "\\.go$" | while read -r event; do
    # Run the Go program and capture its output
    go run main.go

done