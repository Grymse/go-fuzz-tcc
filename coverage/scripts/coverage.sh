#!/bin/bash

script_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

rm -f $script_dir/../../../tinycc/*.gcda
rm -f $script_dir/../csv_files/time_log.csv

bash $script_dir/generate_and_compile.sh
bash $script_dir/coverage_to_csv.sh --reset
bash $script_dir/run_all_plots.sh