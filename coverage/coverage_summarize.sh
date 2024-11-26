#!/bin/bash

# Directory of the script
script_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# Path to the directory containing the CSV files
csv_dir="/home/faur/Programming/ITU/Masters/3rd_Semester/ASA/go-fuzz-tcc/coverage"

# Define the base name for the output file
output_base="$script_dir/summary_averages/gcov_summary_averages"

# Function to get the next available output file name
get_next_output_file() {
    local base_name="$1"
    local suffix=1
    local output_file="${base_name}_${suffix}.csv"
    
    # Loop to find the next available file name
    while [[ -f "$output_file" ]]; do
        suffix=$((suffix + 1))
        output_file="${base_name}_${suffix}.csv"
    done
    
    echo "$output_file"
}

# Get the next available output file
output_averages=$(get_next_output_file "$output_base")

# Initialize the output file with headers
echo "Filename,Average Lines Executed %,Average Branches Executed %,Average Calls Executed %" > "$output_averages"

# Initialize associative arrays for sum and count
declare -A sum_lines
declare -A count_lines
declare -A sum_branches
declare -A count_branches
declare -A sum_calls
declare -A count_calls

# Define the CSV files to process
gcov_files=(
    "$csv_dir/gcov_summary_lines_executed.csv"
    "$csv_dir/gcov_summary_branches_executed.csv"
    "$csv_dir/gcov_summary_calls_executed.csv"
)

# Function to process each CSV file
process_file() {
    local gcov_file="$1"
    local sum_array_name="$2"
    local count_array_name="$3"
    local label="$4"
    
    if [[ -f "$gcov_file" ]]; then
        echo "Processing $gcov_file..."

        while IFS=, read -r filename values; do
            # Skip the header row
            if [[ "$filename" != "Filename" && -n "$values" ]]; then
                filename=$(echo "$filename" | xargs)

                # Split the values into an array
                IFS=',' read -r -a values_array <<< "$values"
                
                # Sum all values and count them
                sum=0
                count=0
                for value in "${values_array[@]}"; do
                    if [[ "$value" =~ ^[0-9]+(\.[0-9]+)?$ ]]; then
                        sum=$(echo "$sum + $value" | bc)
                        count=$((count + 1))
                    fi
                done

                if [[ $count -gt 0 ]]; then
                    average=$(echo "scale=2; $sum / $count" | bc)
                    eval "${sum_array_name}[$filename]=$average"
                    eval "${count_array_name}[$filename]=$count"
                fi
            fi
        done < "$gcov_file"
    else
        echo "File $gcov_file not found."
    fi
}

# Process each of the three files
process_file "${gcov_files[0]}" "sum_lines" "count_lines" "Lines"
process_file "${gcov_files[1]}" "sum_branches" "count_branches" "Branches"
process_file "${gcov_files[2]}" "sum_calls" "count_calls" "Calls"

# Combine the results and write to the output file
for filename in "${!sum_lines[@]}"; do
    # Calculate averages for each category
    lines_avg=${sum_lines[$filename]:-0}
    branches_avg=${sum_branches[$filename]:-0}
    calls_avg=${sum_calls[$filename]:-0}
    
    echo "$filename,$lines_avg,$branches_avg,$calls_avg" >> "$output_averages"
done

echo "Averages saved to $output_averages"
