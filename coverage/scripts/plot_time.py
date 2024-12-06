import pandas as pd
import os
import matplotlib.pyplot as plt
import seaborn as sns

# Base directory where Test folders are located
summary_base_dir = "../csv_files/"
output_plot = "../graphs/average_time_comparison.png"

# Collect all "*_time_log.csv" files in "csv_files/Test x" folders
csv_files = []
for root, dirs, files in os.walk(summary_base_dir):
    for file in files:
        if file.endswith("_time_log.csv"):
            csv_files.append(os.path.join(root, file))

if not csv_files:
    raise FileNotFoundError(f"No matching CSV files found in {summary_base_dir}")

df_list = []

for file_path in csv_files:
    try:
        # Extract the Test folder name (e.g., "Test 1") from the file name
        test_folder = os.path.basename(file_path).split("_")[0] + " " + os.path.basename(file_path).split("_")[1]
        
        # Read CSV and assume it has only 'Average Time' column
        df = pd.read_csv(file_path, header=None, names=['Average Time'])
        
        # Add a column for the test folder name (from file name)
        df['Test Folder'] = test_folder
        df['Filename'] = os.path.basename(file_path)  # Add Filename column

        df_list.append(df)
    except Exception as e:
        print(f"Error reading file {file_path}: {e}")

if not df_list:
    raise ValueError("No valid data found in the CSV files.")

# Combine all data into a single DataFrame
all_data = pd.concat(df_list, ignore_index=True)

# Check if the required column exists
if 'Average Time' not in all_data.columns:
    raise ValueError("'Average Time' column not found in the data.")

# Sort data for better visualization
all_data.sort_values(by='Test Folder', inplace=True)

# Create the plot
plt.figure(figsize=(16, 8))
sns.set(style="whitegrid")

barplot = sns.barplot(
    data=all_data,
    x='Test Folder',
    y='Average Time',
    hue='Filename',
    dodge=True
)

# Add bar labels
for container in barplot.containers:
    barplot.bar_label(
        container,
        fmt='%.2f',
        label_type='edge',
        padding=3
    )

# Customize plot
plt.title("Average Time Comparison Across Test Folders")
plt.xlabel("Test Folder")
plt.ylabel("Average Time")
plt.xticks(rotation=45, ha='right')
plt.legend(title='Filename', bbox_to_anchor=(1.05, 1), loc='upper left')
plt.tight_layout()

# Save the plot
os.makedirs(os.path.dirname(output_plot), exist_ok=True)  # Ensure the output directory exists
plt.savefig(output_plot, dpi=300, bbox_inches='tight')

print(f"Plot saved to {output_plot}")
