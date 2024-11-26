#!/bin/bash

script_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

cd $script_dir/../..
go run main.go > output.c

clang-format -i output.c

# if gcc returns 0, then print success message
if gcc output.c -o output; then
    echo "Compilation successful!"
else 
    echo "Compilation failed!"
fi

../tinycc/tcc -c output.c