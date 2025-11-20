#!/usr/bin/env python3
import sys
import os

try:
    from PyPDF2 import PdfMerger
    print("SUCCESS: PyPDF2 PdfMerger imported successfully")
except ImportError as e:
    print(f"ERROR: Failed to import PyPDF2: {e}")
    print("\nPlease run: ./setup.sh")
    sys.exit(1)

def merge_pdfs(input_files, output_file):
    merger = PdfMerger()
    
    for pdf in input_files:
        print(f"Adding: {pdf}")
        merger.append(pdf)
    
    print(f"Merging to: {output_file}")
    merger.write(output_file)
    merger.close()
    print(f"Successfully merged {len(input_files)} PDF files into {output_file}")

if __name__ == "__main__":
    if len(sys.argv) < 3:
        print("Usage: python merge_pdfs.py output.pdf input1.pdf input2.pdf ...")
        sys.exit(1)
    
    output_file = sys.argv[1]
    input_files = sys.argv[2:]
    
    print(f"Starting PDF merge of {len(input_files)} files...")
    merge_pdfs(input_files, output_file)