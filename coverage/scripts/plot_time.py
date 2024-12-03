import os
import pandas as pd
import matplotlib.pyplot as plt
import seaborn as sns

# Directory containing the time summary files
summary_dir = "../summary_averages/"
output_plot = "../graphs/average_time_bar_plot.png"

# Find all "fuzzer_metrics_averages_time_summary_x" files
time_summary_files = sorted([
    f for f in os.listdir(summary_dir) 
    if f.startswith("fuzzer_metrics_averages_time_summary_") and f.endswith(".csv")
])

if not time_summary_files:
    raise FileNotFoundError(f"No matching files found in {summary_dir}")

# Read and combine all time summary files into a single DataFrame
df_list = []
for file in time_summary_files:
    file_path = os.path.join(summary_dir, file)
    try:
        df = pd.read_csv(file_path)
        df['Source'] = file  # Add a column to identify the source file
        df_list.append(df)
    except Exception as e:
        print(f"Error reading {file_path}: {e}")

if not df_list:
    raise ValueError("No valid data found in the time summary files.")

all_data = pd.concat(df_list, ignore_index=True)

# Ensure the necessary columns exist
if 'Average Time (ms)' not in all_data.columns or 'Source' not in all_data.columns:
    raise ValueError("'Average Time (ms)' or 'Source' column missing from data.")

# Sort the data by the 'Source' column based on the file names' natural order
all_data['Source'] = pd.Categorical(all_data['Source'], categories=time_summary_files, ordered=True)
all_data.sort_values(by='Source', inplace=True)

# Plot the data with adjusted figure height
plt.figure(figsize=(12, 8))  # Increased height from 6 to 8
sns.set(style="whitegrid")

barplot = sns.barplot(
    data=all_data,
    x='Source',
    y='Average Time (ms)',
    color="skyblue"
)

# Annotate the bars with their values
for container in barplot.containers:
    barplot.bar_label(
        container,
        fmt='%.2f',
        label_type='edge',
        padding=3
    )

plt.title("Average Time (ms) Across Fuzzer Metrics Summaries")
plt.xlabel("Summary File")
plt.ylabel("Average Time (ms)")
plt.xticks(rotation=45, ha='right')
plt.tight_layout()

# Save the plot
os.makedirs("../graphs", exist_ok=True)
plt.savefig(output_plot, dpi=300, bbox_inches='tight')
