#!/bin/bash

script_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

bash $script_dir/coverage_to_csv.sh --reset
rm -f $script_dir/../../../tinycc/*.gcda
rm -f $script_dir/../csv_files/time_log.csv

# run 10 times
for i in {1..10}
do
    start_time=$(date +%s.%N)
    bash $script_dir/compileToFile.sh
    end_time=$(date +%s.%N)
    generation_time_ms=$(echo "($end_time - $start_time) * 1000" | bc)
    
    bash $script_dir/writeTimeToFile.sh $generation_time_ms

    bash $script_dir/coverage_to_csv.sh
    rm -f $script_dir/../../../tinycc/*.gcda
done


bash $script_dir/coverage_summarize.sh

python $script_dir/plot_coverage_lines_executed.py
python $script_dir/plot_coverage_branches_executed.py
python $script_dir/plot_coverage_calls_executed.py
python $script_dir/plot_time.py