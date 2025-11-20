#!/bin/bash
echo "Setting up PDF merger..."

# Remove existing venv if any
rm -rf venv

# Create fresh virtual environment
python3 -m venv venv

# Install compatible PyPDF2 version
./venv/bin/pip install PyPDF2==3.0.1

# Test the installation
echo "Testing installation..."
./venv/bin/python3 -c "from PyPDF2 import PdfMerger; print('SUCCESS: PyPDF2 is working!')"

if [ $? -eq 0 ]; then
    echo "Setup complete! Now run: go run main.go"
else
    echo "Setup failed!"
    exit 1
fi