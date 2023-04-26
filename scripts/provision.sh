#!/bin/bash

# Directories
mkdir -p docs examples/basic_usage examples/advanced_usage sentries scripts

# Sentries
for sentry in ra lir hermes thoth tyche jupiter apollo demeter ares morrigan mercury osiris shiva sobek ptah lugh fenrir athena
do
  mkdir -p "sentries/$sentry"
  touch "sentries/$sentry/resource_$sentry.go"
done

# Documentation
touch docs/user_guide.md docs/developer_guide.md docs/api_reference.md

# Scripts
touch scripts/provision.sh scripts/update.sh

# Main files
touch main.go provider.go .gitignore 

# Write this script to provision.sh
cat > scripts/provision.sh <<EOL
$(cat provision.sh)
EOL
