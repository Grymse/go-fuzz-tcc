import pandas as pd
import os
import matplotlib.pyplot as plt
import seaborn as sns

summary_dir = "../summary_averages/"
output_plot = "../graphs/summary_lines_executed_plot.png"

csv_files = [f for f in os.listdir(summary_dir) if f.startswith("fuzzer_metrics_averages") and f.endswith(".csv")]

if not csv_files:
    raise FileNotFoundError(f"No matching CSV files found in {summary_dir}")

df_list = []

for file in csv_files:
    file_path = os.path.join(summary_dir, file)
    try:
        df = pd.read_csv(file_path)
        df['File'] = file  # Add a column for the file name (to identify which file it came from)
        df_list.append(df)
    except Exception as e:
        print(f"Error reading file {file_path}: {e}")

if not df_list:
    raise ValueError("No valid data found in the CSV files.")
all_data = pd.concat(df_list, ignore_index=True)

if 'Average Lines Executed %' not in all_data.columns:
    raise ValueError("'Average Lines Executed %' column not found in the data.")

lines_data = all_data[['Filename', 'File', 'Average Lines Executed %']].copy()


value_ranges = lines_data.groupby('Filename')['Average Lines Executed %'].agg(['min', 'max'])
value_ranges['range'] = value_ranges['max'] - value_ranges['min']

changing_files = value_ranges[value_ranges['range'] > 0].index
lines_data_filtered = lines_data[lines_data['Filename'].isin(changing_files)].copy()

if lines_data_filtered.empty:
    raise ValueError("No changes detected in 'Average Lines Executed %' across files.")

lines_data_filtered.sort_values(by='File', inplace=True)

plt.figure(figsize=(16, 8))
sns.set(style="whitegrid")

barplot = sns.barplot(
    data=lines_data_filtered,
    x='File',
    y='Average Lines Executed %',
    hue='Filename',
    dodge=True
)

for container in barplot.containers:
    barplot.bar_label(
        container,
        fmt='%.2f%%',
        label_type='edge',
        padding=3
    )

plt.title("Changes in Lines Executed Across Summary Files (Filtered for Changes)")
plt.xlabel("Summary File")
plt.ylabel("Average Lines Executed (%)")
plt.xticks(rotation=45, ha='right')
plt.legend(title='Filename', bbox_to_anchor=(1.05, 1), loc='upper left')
plt.tight_layout()

plt.savefig(output_plot, dpi=300, bbox_inches='tight')