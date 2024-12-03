#!/bin/bash

python $script_dir/plot_coverage_lines_executed.py
python $script_dir/plot_coverage_branches_executed.py
python $script_dir/plot_coverage_calls_executed.py
python $script_dir/plot_time.py