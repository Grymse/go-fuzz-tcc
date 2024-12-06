#!/bin/bash

script_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

cd $script_dir/../..

for i in {1..10}
do
    start_time=$(date +%s.%N)

    go run main.go > $script_dir/../code_files/code$i.c

    end_time=$(date +%s.%N)
    generation_time_ms=$(echo "($end_time - $start_time) * 1000" | bc)
    
    bash $script_dir/writeTimeToFile.sh $generation_time_ms

    clang-format -i $script_dir/../code_files/code$i.c
    if gcc $script_dir/../code_files/code$i.c -o $script_dir/../code_files/output/code$i; then
        echo "Compilation successful!"
    else 
        echo "Compilation failed!"
    fi

    # Specify the output directory for the object file when using tcc
    ../tinycc/tcc -c $script_dir/../code_files/code$i.c -o $script_dir/../code_files/output/code$i.o
done
