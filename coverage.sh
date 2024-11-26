#!/bin/bash

bash coverage_to_csv.sh --reset

# run 10 times
for i in {1..10}
do
    bash compileToFile.sh
    bash coverage_to_csv.sh
    rm -f ../tinycc/*.gcda
done