#!/bin/bash

fswatch -o "./" -e ".*" -i "\\.go$" | while read -r event; do
    # Run the Go program and capture its output
    output=$(go run main.go)

    echo "$output"

    # Feed the output into ./resources/tinyc/tinyc
    echo "$output" | ./resources/tinyc/tinyc

done