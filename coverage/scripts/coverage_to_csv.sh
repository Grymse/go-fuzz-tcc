#!/bin/bash

# Directory of the script
script_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
base_dir="$script_dir/../csv_files"  # Base directory for output files
mkdir -p "$base_dir"  # Ensure the base directory exists

results_file="$script_dir/../../results.txt"  # Results file in root directory

# Ensure results.txt exists
if [[ ! -f "$results_file" ]]; then
    touch "$results_file"
fi

echo "Current directory: $(pwd)"

# Function to extract variable values from fuzzer.go
extract_variables() {
    local fuzzer_file="$script_dir/../../fuzzer/fuzzer.go"
    if [[ -f "$fuzzer_file" ]]; then
        echo "Extracting variable values from fuzzer.go..."
        
        # Get a timestamp for this test
        local timestamp=$(date +"%Y-%m-%d %H:%M:%S")
        
        echo "Test: $test_dir_name ($timestamp)" >> "$results_file"
        
        # List of variables to extract
        local variables=("depth" "wavePeak" "waveValleyMin" "waveValleyMax" "waveValley" "maxWaves" "waveCount" "target")
        for var in "${variables[@]}"; do
            # Use grep and awk to extract the variable value
            local value=$(grep -E "var $var =" "$fuzzer_file" | awk -F '=' '{gsub(/ /, "", $2); print $2}')
            if [[ -z "$value" ]]; then
                value="N/A"  # Fallback if the variable is not found
            fi
            echo "$var = $value" >> "$results_file"
        done
        echo "" >> "$results_file"  # Add a blank line for separation
    else
        echo "fuzzer.go not found at $fuzzer_file" >> "$results_file"
    fi
}

# Create a new directory for this test run, e.g., Test 1, Test 2, etc.
test_dir_name="Test_1"
suffix=1
while [[ -d "$base_dir/$test_dir_name" ]]; do
    suffix=$((suffix + 1))
    test_dir_name="Test_$suffix"
done

# Create the new test directory
test_dir="$base_dir/$test_dir_name"
mkdir -p "$test_dir"  # Ensure the test directory exists

# Extract variables and append to results.txt
extract_variables

# Output files for each metric
output_lines_executed="${test_dir}/${test_dir_name}_lines_executed.csv"
output_branches_executed="${test_dir}/${test_dir_name}_branches_executed.csv"
output_calls_executed="${test_dir}/${test_dir_name}_calls_executed.csv"
output_time_log="${test_dir}/${test_dir_name}_time_log.csv"

# Check for --reset flag
if [[ "$1" == "--reset" ]]; then
    echo "==== Resetting CSV files... ===="
    # Delete the existing CSV files if --reset flag is provided
    rm -f "$output_lines_executed" "$output_branches_executed" "$output_calls_executed" "$output_time_log"
fi

# Create CSV headers if the output files don't exist
if [[ ! -f "$output_lines_executed" ]]; then
    echo "Filename,Lines Executed %" > "$output_lines_executed"
fi

if [[ ! -f "$output_branches_executed" ]]; then
    echo "Filename,Branches Executed %" > "$output_branches_executed"
fi

if [[ ! -f "$output_calls_executed" ]]; then
    echo "Filename,Calls Executed %" > "$output_calls_executed"
fi

# Process time log
time_log_file="$base_dir/time_log.csv"
if [[ -f "$time_log_file" ]]; then
    echo "Processing time log..."
    
    # Calculate the average time
    total_time=0
    count=0
    while IFS= read -r line; do
        # Skip the header line
        if [[ "$line" != "Time" ]]; then
            total_time=$(echo "$total_time + $line" | bc)
            count=$((count + 1))
        fi
    done < <(tail -n +2 "$time_log_file")  # Skip the header
    
    if [[ $count -gt 0 ]]; then
        average_time=$(echo "scale=2; $total_time / $count" | bc)
        echo "Average Time,$average_time" > "$output_time_log"
        echo "Time log processed. Average time: $average_time seconds"
    else
        echo "No valid time entries found in $time_log_file"
    fi
else
    echo "Time log file not found at $time_log_file"
fi

cd ../../../tinycc/ || { echo "Directory ../../..tinycc/ not found"; exit 1; }

# Process the .gcno and .gcda files for each .c file in the current directory
for file in *.c; do
    if [[ -f "$file" ]]; then
        # Check if both .gcno and .gcda files exist
        gcno_file="${file%.c}.gcno"
        gcda_file="${file%.c}.gcda"

        if [[ -f "$gcno_file" && -f "$gcda_file" ]]; then
            echo "Processing $file..."

            # Run gcov with -b and capture output, trimming newlines
            summary=$(gcov -b "$file" | tr -d '\n' | tr -s ' ')

            # Extract relevant coverage metrics
            lines_executed=$(echo "$summary" | grep -oP '(?<=Lines executed:)\d+\.\d+(?=%)' | head -1)
            branches_executed=$(echo "$summary" | grep -oP '(?<=Branches executed:)\d+\.\d+(?=%)' | head -1)
            functions_executed=$(echo "$summary" | grep -oP '(?<=Calls executed:)\d+\.\d+(?=%)' | head -1)

            # Default to 0.00 if metrics are missing
            lines_executed=${lines_executed:-0.00}
            branches_executed=${branches_executed:-0.00}
            functions_executed=${functions_executed:-0.00}

            # Append the data to the CSV files
            echo "$file,$lines_executed" >> "$output_lines_executed"
            echo "$file,$branches_executed" >> "$output_branches_executed"
            echo "$file,$functions_executed" >> "$output_calls_executed"
        else
            echo "Skipping $file: Missing .gcno or .gcda file"
        fi
    fi
done

echo ""
echo "==== Coverage summary updated in: ===="
echo "$output_lines_executed"
echo "$output_branches_executed"
echo "$output_calls_executed"
echo "$output_time_log"
