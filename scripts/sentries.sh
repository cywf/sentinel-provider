#!/bin/bash

# This script can populate the individual resource files in the sentries folder
# Modify the provided template as needed

# Example for apollo
cat <<EOT > sentries/apollo/resource_apollo.go
// Your resource_apollo.go code here
EOT

# Add similar blocks for each of the other sentry resource files

echo "Sentries files populated."