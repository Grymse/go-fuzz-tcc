#!/bin/bash

script_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

bash $script_dir/coverage_to_csv.sh --reset

# run 10 times
for i in {1..10}
do
    bash $script_dir/compileToFile.sh
    bash $script_dir/coverage_to_csv.sh
    rm -f ../../tinycc/*.gcda
done

bash $script_dir/coverage_summarize.sh