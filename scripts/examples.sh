#!/bin/bash

# This script can populate the example files in the examples folder
# Modify the provided template as needed

cat <<EOT > examples/basic_usage/example_basic_usage.go
// Please enter relevant code here
EOT

cat <<EOT > examples/advanced_usage/example_advanced_usage.go
// Please enter relevant code here
EOT

# Add similar blocks for any other example files you want to create

echo "Examples files populated."