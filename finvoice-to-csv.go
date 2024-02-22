package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joneskoo/finvoice-to-csv/finvoice"
)

func main() {
	// The directory containing the invoice files is provided as a command-line argument
	if len(os.Args) < 2 {
		log.Fatal("Please provide the directory path as an argument")
	}
	dirPath := os.Args[1]

	// Read the invoice files from the directory
	files, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatal("Failed to read directory:", err)
	}

	// Initialize finvoice CSV writer
	csvWriter := finvoice.NewCSVWriter(os.Stdout)
	defer csvWriter.Flush()

	// Process each invoice file
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".xml" {
			// Open the invoice file
			filePath := filepath.Join(dirPath, file.Name())
			inv, err := finvoice.FromFile(filePath)
			if err != nil {
				log.Fatal("Failed to parse file:", err)
			}

			// Write the InvoiceRow fields to the CSV file
			if err := csvWriter.Write(inv); err != nil {
				log.Fatal("Failed to write CSV:", err)
			}
		}
	}
}
