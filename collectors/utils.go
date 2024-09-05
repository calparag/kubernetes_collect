package collectors

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

func SaveData(resourceType string, data interface{}) {
	// Define the output directory and file path
	outputDir := filepath.Join("data", resourceType)
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		log.Fatalf("Error creating directory %s: %v", outputDir, err)
	}

	filePath := filepath.Join(outputDir, resourceType+".json")
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("Error creating file %s: %v", filePath, err)
	}
	defer file.Close()

	//Marshal the data into JSON
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling JSON for %s: %v", resourceType, err)
	}

	// Write JSON data to file
	_, err = file.Write(jsonData)
	if err != nil {
		log.Fatalf("Error writing to file %s: %v", filePath, err)
	}

	log.Printf("%s information saved to %s", capitalize(resourceType), filePath)
}

func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return string(s[0]-32) + s[1:]
}
