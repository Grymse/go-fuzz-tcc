#!/bin/bash

# Directory of the script
script_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

echo "Current directory: $(pwd)"

# Output files for each metric
output_lines_executed="$script_dir/../csv_files/gcov_summary_lines_executed.csv"
output_branches_executed="$script_dir/../csv_files/gcov_summary_branches_executed.csv"
output_calls_executed="$script_dir/../csv_files/gcov_summary_calls_executed.csv"

# Check for --reset flag
if [[ "$1" == "--reset" ]]; then
    echo "Resetting CSV files..."
    # Delete the existing CSV files if --reset flag is provided
    rm -f "$output_lines_executed" "$output_branches_executed" "$output_calls_executed"
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

cd ../../../tinycc/ || { echo "Directory ../../..tinycc/ not found"; exit 1; }

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

            # Process lines executed CSV
            if ! grep -q "^$file," "$output_lines_executed"; then
                # File doesn't exist in CSV, so we append the new values
                echo "$file,$lines_executed" >> "$output_lines_executed"
            else
                # File exists, update its value (append the result at the end of the line)
                sed -i "/^$file,/s/\(.*\),\(.*\)$/\1,\2,$lines_executed/" "$output_lines_executed"
            fi

            # Process branches executed CSV
            if ! grep -q "^$file," "$output_branches_executed"; then
                # File doesn't exist in CSV, so we append the new values
                echo "$file,$branches_executed" >> "$output_branches_executed"
            else
                # File exists, update its value (append the result at the end of the line)
                sed -i "/^$file,/s/\(.*\),\(.*\)$/\1,\2,$branches_executed/" "$output_branches_executed"
            fi

            # Process calls executed CSV
            if ! grep -q "^$file," "$output_calls_executed"; then
                # File doesn't exist in CSV, so we append the new values
                echo "$file,$functions_executed" >> "$output_calls_executed"
            else
                # File exists, update its value (append the result at the end of the line)
                sed -i "/^$file,/s/\(.*\),\(.*\)$/\1,\2,$functions_executed/" "$output_calls_executed"
            fi
        else
            echo "Skipping $file: Missing .gcno or .gcda file"
        fi
    fi
done

echo "Coverage summary updated in:"
echo "$output_lines_executed"
echo "$output_branches_executed"
echo "$output_calls_executed"
