import pandas as pd
import os
import matplotlib.pyplot as plt
import seaborn as sns

# Base directory where Test folders are located
summary_base_dir = "../csv_files/"
output_plot = "../graphs/branches_executed.png"

# Collect all "*_branches_executed.csv" files in "csv_files/Test x" folders
csv_files = []
for root, dirs, files in os.walk(summary_base_dir):
    for file in files:
        if file.endswith("_branches_executed.csv"):
            csv_files.append(os.path.join(root, file))

if not csv_files:
    raise FileNotFoundError(f"No matching CSV files found in {summary_base_dir}")

df_list = []

for file_path in csv_files:
    try:
        # Extract the Test folder name (e.g., "Test 1") from the file name
        test_folder = os.path.basename(file_path).split("_")[0] + " " + os.path.basename(file_path).split("_")[1]
        df = pd.read_csv(file_path)
        df['Test Folder'] = test_folder  # Add a column for the test folder name
        df_list.append(df)
    except Exception as e:
        print(f"Error reading file {file_path}: {e}")

if not df_list:
    raise ValueError("No valid data found in the CSV files.")

# Combine all data into a single DataFrame
all_data = pd.concat(df_list, ignore_index=True)

# Check if the required column exists
if 'Branches Executed %' not in all_data.columns:
    raise ValueError("'Branches Executed %' column not found in the data.")

# Filter relevant data
lines_data = all_data[['Filename', 'Test Folder', 'Branches Executed %']].copy()

# Calculate min and max values to find ranges
value_ranges = lines_data.groupby('Filename')['Branches Executed %'].agg(['min', 'max'])
value_ranges['range'] = value_ranges['max'] - value_ranges['min']

# Filter files with changes in 'Branches Executed %'
changing_files = value_ranges[value_ranges['range'] > 0].index
lines_data_filtered = lines_data[lines_data['Filename'].isin(changing_files)].copy()

if lines_data_filtered.empty:
    raise ValueError("No changes detected in 'Branches Executed %' across files.")

# Sort data for better visualization
lines_data_filtered.sort_values(by='Test Folder', inplace=True)

# Create the plot
plt.figure(figsize=(16, 8))
sns.set(style="whitegrid")

barplot = sns.barplot(
    data=lines_data_filtered,
    x='Test Folder',
    y='Branches Executed %',
    hue='Filename',
    dodge=True
)

# Add bar labels
for container in barplot.containers:
    barplot.bar_label(
        container,
        fmt='%.2f%%',
        label_type='edge',
        padding=3
    )

# Customize plot
plt.title("Changes in Branches Executed % Across Test Folders (Filtered for Changes)")
plt.xlabel("Test Folder")
plt.ylabel("Branches Executed % (%)")
plt.xticks(rotation=45, ha='right')
plt.legend(title='Filename', bbox_to_anchor=(1.05, 1), loc='upper left')
plt.tight_layout()

# Save the plot
os.makedirs(os.path.dirname(output_plot), exist_ok=True)  # Ensure the output directory exists
plt.savefig(output_plot, dpi=300, bbox_inches='tight')

print(f"Plot saved to {output_plot}")
