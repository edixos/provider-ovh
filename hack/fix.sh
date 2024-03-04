#!/bin/bash

# Directory containing the files
DIRECTORY="./.work/ovh/ovh/website/docs/r"

# Iterate over each file in the directory
for file in "$DIRECTORY"/*
do
    # Extract the name of the file without the path
    filename=$(basename "$file")

    # Replace the first occurrence of '---' with the desired header
    # The header includes the file name for page_title and description
    sed -i '' '1 s/---/---\npage_title: '"$filename"'\ndescription: '"$filename"'/' "$file"
done