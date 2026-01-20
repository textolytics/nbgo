#!/bin/bash

# Create Python virtual environment for testing
python3 -m venv venv

# Activate virtual environment
source venv/bin/activate

# Upgrade pip
pip install --upgrade pip setuptools wheel

# Install common Python packages
pip install pytest pytest-cov
pip install robotframework

echo "Python virtual environment created and activated at: $PWD/venv"
