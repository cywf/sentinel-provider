#!/bin/bash

# Execute each script in sequence
bash scripts/main_and_provider.sh
bash scripts/sentries.sh
bash scripts/scripts.sh
bash scripts/examples.sh
bash scripts/tests.sh

echo "Provisioning complete."