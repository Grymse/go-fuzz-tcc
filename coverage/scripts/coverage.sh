#!/bin/bash

script_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

bash $script_dir/coverage_to_csv.sh --reset
rm -f $script_dir/../../../tinycc/*.gcda

# run 10 times
for i in {1..10}
do
    bash $script_dir/compileToFile.sh
    bash $script_dir/coverage_to_csv.sh
    rm -f $script_dir/../../../tinycc/*.gcda
done

bash $script_dir/coverage_summarize.sh

python $script_dir/plot_coverage.py