#!/bin/bash

# Check if the time argument is provided
if [ -z "$1" ]; then
    echo "Usage: $0 <time_in_seconds>"
    exit 1
fi

script_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"


# Get the time measurement from the argument
time_measurement=$1

# Define the CSV file name
csv_file="$script_dir/../csv_files/time_log.csv"

# Add headers if the file does not exist
if [ ! -f "$csv_file" ]; then
    echo "Timestamp,Time (seconds)" > "$csv_file"
fi

# Append the timestamp and time measurement to the CSV file
echo "$time_measurement" >> "$csv_file"

echo "Logged: $time_measurement seconds to $csv_file"