#!/bin/bash

# Setup using UV package manager for Python
# Install uv if not present
if ! command -v uv &> /dev/null; then
    curl -LsSf https://astral.sh/uv/install.sh | sh
fi

# Create UV-managed Python environment
uv venv --python 3.11 venv_uv

# Activate the environment
source venv_uv/bin/activate

# Install dependencies
uv pip install pytest pytest-cov robotframework

echo "UV Python environment created and activated at: $PWD/venv_uv"
