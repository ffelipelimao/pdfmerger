package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	inputFolder := "./pdfs"
	outputFile := "./final_files.pdf"

	fmt.Printf("Looking for PDFs in folder: %s\n", inputFolder)

	pdfFiles, err := getPDFFilesInOrder(inputFolder)
	if err != nil {
		log.Fatalf("Error listing PDFs: %v", err)
	}

	if len(pdfFiles) == 0 {
		log.Fatal("No PDF files found in the folder")
	}

	fmt.Printf("Found %d PDF files in original order:\n", len(pdfFiles))
	for i, pdf := range pdfFiles {
		fmt.Printf("%d. %s\n", i+1, filepath.Base(pdf))
	}

	// Merge PDFs using Python
	fmt.Println("\nMerging PDFs using Python...")
	err = mergeWithPython(pdfFiles, outputFile)
	if err != nil {
		log.Fatalf("Error merging PDFs: %v", err)
	}

	fmt.Printf("Merge completed! File saved as: %s\n", outputFile)
}

func getPDFFilesInOrder(folder string) ([]string, error) {
	// REPLACE HERE THE ORDER OF THE FILES IN ./pdf folder
	manualOrder := []string{}

	var pdfFiles []string

	for _, filename := range manualOrder {
		fullPath := filepath.Join(folder, filename)
		if _, err := os.Stat(fullPath); err == nil {
			pdfFiles = append(pdfFiles, fullPath)
		}
	}

	return pdfFiles, nil
}

func mergeWithPython(files []string, output string) error {
	pythonCmd := "./venv/bin/python3"

	if _, err := os.Stat(pythonCmd); os.IsNotExist(err) {
		return fmt.Errorf("virtual environment not found. Please run ./setup.sh")
	}

	args := []string{"merge_pdfs.py", output}
	args = append(args, files...)

	fmt.Printf("Executing: %s merge_pdfs.py %s [%d files]\n", pythonCmd, output, len(files))

	cmd := exec.Command(pythonCmd, args...)

	outputBytes, err := cmd.CombinedOutput()
	outputStr := string(outputBytes)

	fmt.Printf("Python output:\n%s\n", outputStr)

	if err != nil {
		return fmt.Errorf("python script failed: %v", err)
	}

	return nil
}
