#!/bin/bash

fswatch -o "./" -e ".*" -i "\\.go$" | while read -r event; do
    # Run the Go program and capture its output
    go run main.go > output.c

    clang-format -i output.c

    # if gcc returns 0, then print success message
    if gcc output.c -o output; then
        echo "Compilation successful!"
    else 
        echo "Compilation failed!"
    fi

done